package maas

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api"
	"github.com/volcengine/volc-sdk-golang/service/maas/sse"
)

// MaaS ... use base client
type MaaS struct {
	*base.Client
}

// NewInstance ...
func NewInstance(host, region string) *MaaS {
	instance := &MaaS{}
	instance.Client = base.NewClient(&base.ServiceInfo{
		Timeout: ServiceTimeout,
		Scheme:  "https",
		Host:    host,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{
			Region:  region,
			Service: ServiceName,
		},
	}, map[string]*base.ApiInfo{
		APIChat: {
			Method: http.MethodPost,
			Path:   "/api/v1/chat",
		},
		APIStreamChat: {
			Method: http.MethodPost,
			Path:   "/api/v1/chat",
		},
		APICert: {
			Method: http.MethodPost,
			Path:   "/api/v1/cert",
		},
		APIClassification: {
			Method: http.MethodPost,
			Path:   "/api/v1/classification",
		},
		APITokenization: {
			Method: http.MethodPost,
			Path:   "/api/v1/tokenization",
		},
		APIEmbeddings: {
			Method: http.MethodPost,
			Path:   "/api/v1/embeddings",
		},
	})

	return instance
}

// POST method
// Chat ...
func (cli *MaaS) Chat(req *api.ChatReq) (*api.ChatResp, int, error) {
	return cli.ChatWithCtx(context.Background(), req)
}

// POST method
// ChatWithCtx ...
func (cli *MaaS) ChatWithCtx(ctx context.Context, req *api.ChatReq) (*api.ChatResp, int, error) {
	req.Stream = false

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()))
	}

	return cli.ChatImpl(ctx, bts)
}

// POST method
// SecretChat is like `Chat`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretChat(req *api.ChatReq) (*api.ChatResp, int, error) {
	return cli.SecretChatWithCtx(context.Background(), req)
}

// POST method
// SecretChatWithCtx is like `ChatWithCtx`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretChatWithCtx(ctx context.Context, req *api.ChatReq) (*api.ChatResp, int, error) {
	key, nonce, req, err := cli.encryptChatRequest(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to encrypt chat request: %v", err))
	}

	output, status, err := cli.ChatWithCtx(ctx, req)
	if err != nil {
		return nil, status, err
	}

	output, err = cli.decryptChatResponse(key, nonce, output)
	if err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to decrypt chat response: %v", err))
	}
	return output, status, nil
}

// POST method
// SecretStreamChat is like `StreamChat`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretStreamChat(req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	return cli.SecretStreamChatWithCtx(context.Background(), req)
}

// POST method
// SecretStreamChatWithCtx is like `StreamChatWithCtx`, except its messages are encrypted
// to ensure that messages are not intercepted by receivers other than the model.
func (cli *MaaS) SecretStreamChatWithCtx(ctx context.Context, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	key, nonce, req, err := cli.encryptChatRequest(req)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to encrypt chat request: %v", err))
	}

	resps, err := cli.StreamChatWithCtx(ctx, req)
	if err != nil {
		return nil, err
	}

	outputs := make(chan *api.ChatResp, RespBufferSize)
	go func() {
		defer func() {
			_ = recover()
			close(outputs)
		}()

		for resp := range resps {
			output, err := cli.decryptChatResponse(key, nonce, resp)
			if err != nil {
				resp.Error = api.NewClientSDKRequestError(fmt.Sprintf("failed to decrypt chat response: %v", err))
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
func (cli *MaaS) StreamChat(req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	return cli.StreamChatWithCtx(context.Background(), req)
}

// POST method
// StreamChat make stream chat request
//  1. if any error returned, a channel=`nil` is returned;
//  2. if no error returned, the channel are closed after all responses processed.
func (cli *MaaS) StreamChatWithCtx(ctx context.Context, req *api.ChatReq) (ch <-chan *api.ChatResp, err error) {
	req.Stream = true

	bts, err := json.Marshal(req)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()))
	}

	return cli.StreamChatImpl(ctx, bts)
}

func (cli *MaaS) ChatImpl(ctx context.Context, body []byte) (*api.ChatResp, int, error) {
	respBody, status, err := cli.Client.CtxJson(ctx, APIChat, nil, string(body))
	if err != nil {
		errVal := &api.ChatResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error())
		}
		return nil, status, errVal.Error
	}

	output := new(api.ChatResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()))
	}
	return output, status, nil
}

func (cli *MaaS) StreamChatImpl(ctx context.Context, body []byte) (<-chan *api.ChatResp, error) {
	apiInfo := cli.ApiInfoList[APIStreamChat]
	if apiInfo == nil {
		return nil, api.NewClientSDKRequestError("the related api does not exist")
	}

	// build request
	req, err := MakeRequest(apiInfo, "", cli.ServiceInfo, nil, "application/json")
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to make request: %v", err))
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	timeout := GetTimeout(cli.ServiceInfo.Timeout, apiInfo.Timeout)

	req = cli.ServiceInfo.Credentials.Sign(req)

	ctx, cancel := context.WithTimeout(ctx, timeout)
	req = req.WithContext(ctx)

	// do request
	resp, err := cli.Client.Client.Do(req)
	if err != nil {
		cancel()
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("request error: %v", err))
	}

	if resp.StatusCode != 200 { // fast fail
		res := &api.ChatResp{}
		if er := json.NewDecoder(resp.Body).Decode(res); er != nil {
			res.Error = api.NewClientSDKRequestError(fmt.Sprintf("failed to call service: http status_code=%d", resp.StatusCode))
		}
		cancel()
		_ = resp.Body.Close()
		return nil, res.Error
	}

	// parse response
	ch := make(chan *api.ChatResp, RespBufferSize)
	go func() {
		defer func() {
			_ = recover()
			_ = resp.Body.Close()
			cancel()
			close(ch)
		}()

		stream := sse.NewEventStreamFromReader(resp.Body, MaxBufferSize)
		for {
			event, err := stream.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				if errors.Is(err, context.DeadlineExceeded) {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(fmt.Sprintf("call service timeout: timeout=%s", timeout.String())),
					}
				} else {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(err.Error()),
					}
				}
				return
			}
			if event != nil {
				if bytes.Equal(event.Data, []byte(Terminator)) {
					return
				}

				item := &api.ChatResp{}
				if err = json.Unmarshal(event.Data, item); err != nil {
					ch <- &api.ChatResp{
						Error: api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response(data=%s): %v", string(event.Data), err)),
					}
					return
				}
				ch <- item
			}
		}
	}()

	return ch, nil
}

func (r *MaaS) initCertByReq(req *api.ChatReq) (*KeyAgreementClient, error) {
	certReq := &api.CertReq{
		Model: req.Model,
	}
	body, err := json.Marshal(certReq)
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()))
	}
	respBody, _, err := r.Client.Json(APICert, nil, string(body))
	if err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to get CA from proxy: %s", err.Error()))
	}
	output := new(api.CertResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()))
	}

	// use returned model
	req.Model = output.Model

	// todo: check chain
	return NewP256KeyAgreementClient(output.Cert)
}

func (cli *MaaS) encryptChatRequest(req *api.ChatReq) ([]byte, []byte, *api.ChatReq, error) {
	ka, err := cli.initCertByReq(req)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to init cert: %w", err)
	}
	key, nonce, token, err := ka.GenerateECIESKeyPair()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate session token: %w", err)
	}

	req.CryptoToken = token

	for i := range req.Messages {
		content := req.Messages[i].GetContent()
		if content != "" {
			secret, err := AesGcmEncryptBase64String(key, nonce, req.Messages[i].Content)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("failed to encrypt message: %w", err)
			}
			req.Messages[i].Content = secret
		}
	}

	return key, nonce, req, nil
}

func (cli *MaaS) decryptChatResponse(key, nonce []byte, resp *api.ChatResp) (*api.ChatResp, error) {
	if secret := resp.GetChoice().GetMessage().GetContent(); secret != "" {
		plain, err := AesGcmDecryptBase64String(key, nonce, secret)
		if err != nil {
			return nil, err
		}
		resp.GetChoice().GetMessage().Content = plain
	}
	return resp, nil
}

// POST method
// Tokenization
func (cli *MaaS) Tokenization(req *api.TokenizeReq) (*api.TokenizeResp, int, error) {
	return cli.TokenizationWithCtx(context.Background(), req)
}

func (cli *MaaS) TokenizationWithCtx(ctx context.Context, req *api.TokenizeReq) (*api.TokenizeResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()))
	}
	return cli.tokenizationImpl(ctx, bts)
}

func (cli *MaaS) tokenizationImpl(ctx context.Context, body []byte) (*api.TokenizeResp, int, error) {
	respBody, status, err := cli.Client.CtxJson(ctx, APITokenization, nil, string(body))
	if err != nil {
		errVal := &api.TokenizeResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error())
		}
		return nil, status, errVal.Error
	}

	output := new(api.TokenizeResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()))
	}
	return output, status, nil
}

// POST method
// Classification
func (cli *MaaS) Classification(req *api.ClassificationReq) (*api.ClassificationResp, int, error) {
	return cli.ClassificationWithCtx(context.Background(), req)
}

func (cli *MaaS) ClassificationWithCtx(ctx context.Context, req *api.ClassificationReq) (*api.ClassificationResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()))
	}
	return cli.classificationImpl(ctx, bts)
}

func (cli *MaaS) classificationImpl(ctx context.Context, body []byte) (*api.ClassificationResp, int, error) {
	respBody, status, err := cli.Client.CtxJson(ctx, APIClassification, nil, string(body))
	if err != nil {
		errVal := &api.ClassificationResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error())
		}
		return nil, status, errVal.Error
	}

	output := new(api.ClassificationResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()))
	}
	return output, status, nil
}

// POST method
// Embeddings
func (cli *MaaS) Embeddings(req *api.EmbeddingsReq) (*api.EmbeddingsResp, int, error) {
	return cli.EmbeddingsWithCtx(context.Background(), req)
}

func (cli *MaaS) EmbeddingsWithCtx(ctx context.Context, req *api.EmbeddingsReq) (*api.EmbeddingsResp, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, api.NewClientSDKRequestError(fmt.Sprintf("failed to marshal request: %s", err.Error()))
	}
	return cli.embeddingsImpl(ctx, bts)
}

func (cli *MaaS) embeddingsImpl(ctx context.Context, body []byte) (*api.EmbeddingsResp, int, error) {
	respBody, status, err := cli.Client.CtxJson(ctx, APIEmbeddings, nil, string(body))
	if err != nil {
		errVal := &api.EmbeddingsResp{}
		if er := json.Unmarshal(respBody, errVal); er != nil {
			errVal.Error = api.NewClientSDKRequestError(err.Error())
		}
		return nil, status, errVal.Error
	}

	output := new(api.EmbeddingsResp)
	if err = json.Unmarshal(respBody, output); err != nil {
		return nil, status, api.NewClientSDKRequestError(fmt.Sprintf("failed to unmarshal response: %s", err.Error()))
	}
	return output, status, nil
}
