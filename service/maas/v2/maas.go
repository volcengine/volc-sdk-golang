package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/martian/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/maas"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api/v2"
	"github.com/volcengine/volc-sdk-golang/service/maas/sse"
)

// MaaS ... use base client
type MaaS struct {
	*base.Client
	apikey    string
	apikeyTtl int64
}

func (cli *MaaS) SetApikey(apikey string) {
	cli.apikey = apikey
}

// NewInstance ...
func NewInstance(host, region string) *MaaS {
	instance := &MaaS{}
	instance.Client = base.NewClient(&base.ServiceInfo{
		Timeout: maas.ServiceTimeout,
		Scheme:  "https",
		Host:    host,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{
			Region:  region,
			Service: maas.ServiceName,
		},
	}, map[string]*base.ApiInfo{
		maas.APIChat: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/chat",
		},
		maas.APIStreamChat: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/chat",
		},
		maas.APICert: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/cert",
		},
		maas.APIClassification: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/classification",
		},
		maas.APITokenization: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/tokenization",
		},
		maas.APIEmbeddings: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/embeddings",
		},
		maas.APIImagesQuickGen: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/images/quick-gen",
		},
		maas.APIAudioSpeech: {
			Method: http.MethodPost,
			Path:   "/api/v2/endpoint/%s/audio/speech",
		},
		maas.TOP: {
			Method: http.MethodPost,
			Path:   "/",
		},
	})

	return instance
}

func (cli *MaaS) Images() *images {
	return &images{cli}
}

func (cli *MaaS) Audio() *audio {
	return &audio{
		Speech{m: cli},
	}
}

// POST method
// Chat ...
func (cli *MaaS) Chat(endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	return cli.ChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// ChatWithCtx ...
func (cli *MaaS) ChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	req.Stream = false

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}

	return cli.ChatImpl(ctx, endpointId, bts)
}

// POST method
// SecretChat is like `Chat`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretChat(endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	return cli.SecretChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// SecretChatWithCtx is like `ChatWithCtx`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (*api.ChatResp, int, error) {
	key, nonce, req, err := cli.encryptChatRequest(ctx, endpointId, req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to encrypt chat request: %v", err), "")
	}

	output, status, err := cli.ChatWithCtx(ctx, endpointId, req)
	if err != nil {
		return nil, status, err
	}

	output, err = cli.decryptChatResponse(key, nonce, output)
	if err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to decrypt chat response: %v", err), "")
	}
	return output, status, nil
}

// POST method
// SecretStreamChat is like `StreamChat`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretStreamChat(endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	return cli.SecretStreamChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// SecretStreamChatWithCtx is like `StreamChatWithCtx`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretStreamChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	key, nonce, req, err := cli.encryptChatRequest(ctx, endpointId, req)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to encrypt chat request: %v", err), "")
	}

	resps, err := cli.StreamChatWithCtx(ctx, endpointId, req)
	if err != nil {
		return nil, err
	}

	outputs := make(chan *api.ChatResp, maas.RespBufferSize)
	go func() {
		defer func() {
			_ = recover()
			close(outputs)
		}()

		for resp := range resps {
			output, err := cli.decryptChatResponse(key, nonce, resp)
			if err != nil {
				resp.Error = api.NewClientSDKRequestError(fmt.Sprintf("failed to decrypt chat response: %v", err), "")
				outputs <- resp
				continue
			}
			outputs <- output
		}
	}()
	return outputs, nil
}

// POST method
// StreamChat make stream chat request
//  1. if any error returned, a channel=`nil` is returned;
//  2. if no error returned, the channel are closed after all responses processed.
func (cli *MaaS) StreamChat(endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	return cli.StreamChatWithCtx(context.Background(), endpointId, req)
}

// POST method
// StreamChat make stream chat request
//  1. if any error returned, a channel=`nil` is returned;
//  2. if no error returned, the channel are closed after all responses processed.
func (cli *MaaS) StreamChatWithCtx(ctx context.Context, endpointId string, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	req.Stream = true

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}

	return cli.StreamChatImpl(ctx, endpointId, bts)
}

func (cli *MaaS) ChatImpl(ctx context.Context, endpointId string, body []byte) (*api.ChatResp, int, error) {
	ctx = getContext(ctx)

	if endpointId != "" && cli.apikey == "" && cli.apikeyTtl == 0 {
		if err := cli.getKey(endpointId); err != nil {
			return nil, 0, err
		}
	}

	respBody, status, err := cli.request(ctx, maas.APIChat, nil, endpointId, body, cli.apikey)
	if err != nil {
		return nil, status, err
	}

	output := new(api.ChatResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), reqIdFromCtx(ctx))
	}
	output.ReqId = reqIdFromCtx(ctx)
	return output, status, nil
}

func (cli *MaaS) StreamChatImpl(ctx context.Context, endpointId string, body []byte) (<-chan *api.ChatResp, error) {
	ctx = getContext(ctx)

	apiInfo := cli.ApiInfoList[maas.APIStreamChat]
	if apiInfo == nil {
		return nil, api.NewClientSDKRequestError("the related api does not exist", reqIdFromCtx(ctx))
	}

	// build request
	req, err := maas.MakeRequest(apiInfo, endpointId, cli.ServiceInfo, nil, "application/json")
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to make request: %v", err), reqIdFromCtx(ctx))
	}
	req.Header.Add(reqIdHeaderKey, reqIdFromCtx(ctx))
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	timeout := maas.GetTimeout(cli.ServiceInfo.Timeout, apiInfo.Timeout)

	req = cli.ServiceInfo.Credentials.Sign(req)
	req.Header.Set(reqAuthorizationHeaderKey, "Bearer "+cli.apikey)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	req = req.WithContext(ctx)

	// do request
	resp, err := cli.Client.Client.Do(req)
	if err != nil {
		cancel()
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("request error: %v", err), reqIdFromCtx(ctx))
	}

	if resp.StatusCode != 200 { // fast fail
		res := &api.ChatResp{}
		if er := json.NewDecoder(resp.Body).Decode(res); er != nil {
			res.Error = api.NewClientSDKRequestError(fmt.Sprintf("failed to call service: http status_code=%d", resp.StatusCode), reqIdFromCtx(ctx))
		}
		cancel()
		_ = resp.Body.Close()
		return nil, res.Error
	}

	// parse response
	ch := make(chan *api.ChatResp, maas.RespBufferSize)
	go func() {
		defer func() {
			_ = recover()
			_ = resp.Body.Close()
			cancel()
			close(ch)
		}()

		stream := sse.NewEventStreamFromReader(resp.Body, maas.MaxBufferSize)
		for {
			event, err := stream.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				if errors.Is(err, context.DeadlineExceeded) {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(fmt.Sprintf("call service timeout: timeout=%s", timeout.String()), reqIdFromCtx(ctx)),
					}
				} else {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(err.Error(), reqIdFromCtx(ctx)),
					}
				}
				return
			}
			if event != nil {
				if bytes.Equal(event.Data, []byte(maas.Terminator)) {
					return
				}

				item := &api.ChatResp{}
				if err = json.Unmarshal(event.Data, item); err != nil {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response(data=%s): %v", string(event.Data), err), reqIdFromCtx(ctx)),
					}
					return
				}
				item.ReqId = reqIdFromCtx(ctx)
				if item.Error != nil {
					item.Error.ReqId = reqIdFromCtx(ctx)
				}
				ch <- item
			}
		}
	}()

	return ch, nil
}

func (cli *MaaS) initCertByReq(ctx context.Context, endpointId string, req *api.ChatReq) (*maas.KeyAgreementClient, error) {
	certReq := &api.CertReq{}
	body, err := json.Marshal(certReq)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), reqIdFromCtx(ctx))
	}
	respBody, _, err := cli.request(ctx, maas.APICert, nil, endpointId, body, cli.apikey)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to get CA from proxy: %s", err.Error()), reqIdFromCtx(ctx))
	}
	output := new(api.CertResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), reqIdFromCtx(ctx))
	}

	// todo: check chain
	return maas.NewP256KeyAgreementClient(output.Cert)
}

func (cli *MaaS) encryptChatRequest(ctx context.Context, endpointId string, req *api.ChatReq) ([]byte, []byte, *api.ChatReq, error) {
	ka, err := cli.initCertByReq(ctx, endpointId, req)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to init cert: %w", err)
	}
	key, nonce, token, err := ka.GenerateECIESKeyPair()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	req.CryptoToken = token

	for i := range req.Messages {
		content, ok := req.Messages[i].Content.(string)
		if ok && content != "" {
			secret, err := maas.AesGcmEncryptBase64String(key, nonce, content)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to encrypt message: %w", err)
			}
			req.Messages[i].Content = secret
		}
	}

	return key, nonce, req, nil
}

func (cli *MaaS) decryptChatResponse(key, nonce []byte, resp *api.ChatResp) (*api.ChatResp, error) {
	if len(resp.Choices) == 0 {
		return resp, nil
	}
	if secret, ok := resp.Choices[0].Message.Content.(string); ok && secret != "" {
		plain, err := maas.AesGcmDecryptBase64String(key, nonce, secret)
		if err != nil {
			return nil, err
		}
		resp.Choices[0].Message.Content = plain
	}
	return resp, nil
}

// POST method
// Tokenization
func (cli *MaaS) Tokenization(endpointId string, req *api.TokenizeReq) (*api.TokenizeResp, int, error) {
	return cli.TokenizationWithCtx(context.Background(), endpointId, req)
}

func (cli *MaaS) TokenizationWithCtx(ctx context.Context, endpointId string, req *api.TokenizeReq) (*api.TokenizeResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return cli.tokenizationImpl(ctx, endpointId, bts)
}

func (cli *MaaS) tokenizationImpl(ctx context.Context, endpointId string, body []byte) (*api.TokenizeResp, int, error) {
	ctx = getContext(ctx)

	if endpointId != "" && cli.apikey == "" && cli.apikeyTtl == 0 {
		if err := cli.getKey(endpointId); err != nil {
			return nil, 0, err
		}
	}

	respBody, status, err := cli.request(ctx, maas.APITokenization, nil, endpointId, body, cli.apikey)
	if err != nil {
		return nil, status, err
	}

	output := new(api.TokenizeResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), reqIdFromCtx(ctx))
	}
	output.ReqId = reqIdFromCtx(ctx)
	return output, status, nil
}

// POST method
// Classification
func (cli *MaaS) Classification(endpointId string, req *api.ClassificationReq) (*api.ClassificationResp, int, error) {
	return cli.ClassificationWithCtx(context.Background(), endpointId, req)
}

func (cli *MaaS) ClassificationWithCtx(ctx context.Context, endpointId string, req *api.ClassificationReq) (*api.ClassificationResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), reqIdFromCtx(ctx))
	}
	return cli.classificationImpl(ctx, endpointId, bts)
}

func (cli *MaaS) classificationImpl(ctx context.Context, endpointId string, body []byte) (*api.ClassificationResp, int, error) {
	ctx = getContext(ctx)

	if endpointId != "" && cli.apikey == "" && cli.apikeyTtl == 0 {
		if err := cli.getKey(endpointId); err != nil {
			return nil, 0, err
		}
	}

	respBody, status, err := cli.request(ctx, maas.APIClassification, nil, endpointId, body, cli.apikey)
	if err != nil {
		return nil, status, err
	}

	output := new(api.ClassificationResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), reqIdFromCtx(ctx))
	}
	output.ReqId = reqIdFromCtx(ctx)
	return output, status, nil
}

// POST method
// Embeddings
func (cli *MaaS) Embeddings(endpointId string, req *api.EmbeddingsReq) (*api.EmbeddingsResp, int, error) {
	return cli.EmbeddingsWithCtx(context.Background(), endpointId, req)
}

func (cli *MaaS) EmbeddingsWithCtx(ctx context.Context, endpointId string, req *api.EmbeddingsReq) (*api.EmbeddingsResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return cli.embeddingsImpl(ctx, endpointId, bts)
}

func (cli *MaaS) embeddingsImpl(ctx context.Context, endpointId string, body []byte) (*api.EmbeddingsResp, int, error) {
	ctx = getContext(ctx)

	if endpointId != "" && cli.apikey == "" && cli.apikeyTtl == 0 {
		if err := cli.getKey(endpointId); err != nil {
			return nil, 0, err
		}
	}

	respBody, status, err := cli.request(ctx, maas.APIEmbeddings, nil, endpointId, body, cli.apikey)
	if err != nil {
		return nil, status, err
	}

	output := new(api.EmbeddingsResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), reqIdFromCtx(ctx))
	}
	output.ReqId = reqIdFromCtx(ctx)
	return output, status, nil
}

// POST method
// Tokenization
func (cli *MaaS) CreateOrRefreshApiKey(req *api.CreateOrRefreshApiKeyRequest) (*api.CreateOrRefreshApiKeyResponse, int, error) {
	return cli.CreateOrRefreshApiKeyWithCtx(context.Background(), req)
}

func (cli *MaaS) CreateOrRefreshApiKeyWithCtx(ctx context.Context, req *api.CreateOrRefreshApiKeyRequest) (*api.CreateOrRefreshApiKeyResponse, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()), "")
	}
	return cli.createOrRefreshApiKeyImpl(ctx, bts)
}

func (cli *MaaS) createOrRefreshApiKeyImpl(ctx context.Context, body []byte) (*api.CreateOrRefreshApiKeyResponse, int, error) {
	ctx = getContext(ctx)

	query := map[string][]string{"Version": {"2024-01-01"}, "Action": {"CreateOrRefreshAPIKey"}}
	respBody, status, err := cli.request(ctx, maas.TOP, query, "", body, "")
	if err != nil {
		return nil, status, err
	}

	var respMap map[string]interface{}
	err = json.Unmarshal(respBody, &respMap)
	if err != nil {
		return nil, status, err
	}

	respResult, ok := respMap["Result"]
	if !ok {
		return nil, status, fmt.Errorf("api_key error, please contact oncall")
	}

	respResultByteData, err := json.Marshal(respResult)
	if err != nil {
		return nil, status, err
	}

	output := new(api.CreateOrRefreshApiKeyResponse)
	if err = json.Unmarshal(respResultByteData, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()), "")
	}
	return output, status, nil
}

func (cli *MaaS) doRequest(inputContext context.Context, api string, req *http.Request, timeout time.Duration, authApikey string) ([]byte, int, bool, error) {
	req = cli.ServiceInfo.Credentials.Sign(req)

	if authApikey != "" {
		req.Header.Set(reqAuthorizationHeaderKey, "Bearer "+authApikey)
	}

	log.Debugf("req = %v", req)

	ctx := inputContext
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := cli.Client.Client.Do(req)
	if err != nil {
		// should retry when client sends request error.
		return []byte(""), 500, true, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), resp.StatusCode, false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		needRetry := false
		// should retry when server returns 5xx error.
		if resp.StatusCode >= http.StatusInternalServerError {
			needRetry = true
		}
		return body, resp.StatusCode, needRetry, fmt.Errorf("api %s http code %d body %s", api, resp.StatusCode, string(body))
	}

	return body, resp.StatusCode, false, nil
}

func (cli *MaaS) request(ctx context.Context, apiKey string, query url.Values, endpointId string, body []byte, authApikey string) ([]byte, int, error) {
	apiInfo := cli.ApiInfoList[apiKey]
	if apiInfo == nil {
		return nil, 500, api.NewClientSDKRequestError("the related api does not exist", reqIdFromCtx(ctx))
	}

	// build request
	req, err := maas.MakeRequest(apiInfo, endpointId, cli.ServiceInfo, query, "application/json")
	if err != nil {
		return nil, 500, api.NewClientSDKRequestError(fmt.Sprintf("failed to make request: %v", err), reqIdFromCtx(ctx))
	}
	req.Header.Add(reqIdHeaderKey, reqIdFromCtx(ctx))
	requestBody := bytes.NewReader(body)
	timeout := maas.GetTimeout(cli.ServiceInfo.Timeout, apiInfo.Timeout)
	retrySettings := maas.GetRetrySetting(&cli.ServiceInfo.Retry, &apiInfo.Retry)

	var resp []byte
	var code int

	err = backoff.Retry(func() error {
		_, err = requestBody.Seek(0, io.SeekStart)
		if err != nil {
			// if seek failed, stop retry.
			return backoff.Permanent(err)
		}
		req.Body = ioutil.NopCloser(requestBody)
		var needRetry bool
		resp, code, needRetry, err = cli.doRequest(ctx, apiKey, req, timeout, authApikey)
		if needRetry {
			return err
		} else {
			return backoff.Permanent(err)
		}
	}, backoff.WithMaxRetries(backoff.NewConstantBackOff(*retrySettings.RetryInterval), *retrySettings.RetryTimes))

	if err != nil {
		errVal := &api.ErrorResp{}
		if er := json.Unmarshal(resp, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error(), reqIdFromCtx(ctx))
		}
		errVal.Error.ReqId = reqIdFromCtx(ctx)
		err = errVal.Error
	}

	return resp, code, err
}

func (cli *MaaS) getKey(endpointId string) error {
	var err error
	apikey := cli.apikey
	if apikey == "" {
		cli.apikey, err = cli.applyApikey(endpointId, 0, nil)
		if err != nil {
			return err
		}
	}

	if time.Now().Unix()+300 > cli.apikeyTtl {
		cli.apikey, err = cli.applyApikey(endpointId, 0, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cli *MaaS) applyApikey(endpointId string, Ttl int64, endpointIdList []string) (string, error) {
	topCli := NewInstance("open.volcengineapi.com", "cn-beijing")
	topCli.ServiceInfo.Credentials.AccessKeyID = cli.ServiceInfo.Credentials.AccessKeyID
	topCli.ServiceInfo.Credentials.SecretAccessKey = cli.ServiceInfo.Credentials.SecretAccessKey

	if Ttl == 0 {
		Ttl = 604800
	}
	req := &api.CreateOrRefreshApiKeyRequest{
		Ttl:            Ttl, // expected apiKey expired time
		EndpointIdList: []string{endpointId},
	}
	rsp, _, err := topCli.CreateOrRefreshApiKey(req)
	if err != nil {
		errVal := &api.Error{}
		if errors.As(err, &errVal) { // the returned error always type of *api.Error
			fmt.Printf("meet maas error=%v", errVal)
		}
		return "", err
	}
	fmt.Printf("rsp.ApiKey=%v", rsp.ApiKey)
	apikey := rsp.ApiKey
	cli.apikey = apikey
	cli.apikeyTtl = time.Now().Unix() + Ttl
	return apikey, nil
}
