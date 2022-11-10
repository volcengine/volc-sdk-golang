package imagex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/avast/retry-go"
)

type uploadPartCommonResponse struct {
	Version string          `json:"Version"`
	Success int             `json:"success,omitempty"`
	Error   uploadPartError `json:"error"`
}

type uploadPartResponse struct {
	uploadPartCommonResponse
	PayLoad initPartPayLoad `json:"payload,omitempty"`
}

type uploadMergeResponse struct {
	uploadPartCommonResponse
	PayLoad mergePartPayLoad `json:"payload,omitempty"`
}

type mergePartPayLoad struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type uploadPartError struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type initPartPayLoad struct {
	UploadID string `json:"uploadID"`
}

type directUploadResponse struct {
	Success int                 `json:"success"`
	Payload directUploadPayload `json:"payload,omitempty"`
}

type directUploadPayload struct {
	Hash string `json:"hash"`
}

func (c *ImageX) directUpload(ctx context.Context, host string, storeInfo StoreInfo, imageBytes []byte) error {
	if len(imageBytes) == 0 {
		return fmt.Errorf("file size is zero")
	}

	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(imageBytes))
	url := fmt.Sprintf("https://%s/%s", host, storeInfo.StoreUri)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(imageBytes))
	if err != nil {
		return fmt.Errorf("fail to new put request %s, %v", url, err)
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", storeInfo.Auth)
	req = req.WithContext(ctx)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("fail to do request to %s, %v", url, err)
	}

	//goland:noinspection GoDeprecation
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("fail to read response body %s, %v", url, err)
	}

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status=%v, body=%s, url=%s", rsp.StatusCode, string(body), url)
	}
	defer rsp.Body.Close()

	var putResp = directUploadResponse{
		Success: -1,
	}
	if err = json.Unmarshal(body, &putResp); err != nil {
		return fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(body), err)
	}
	if putResp.Success != 0 {
		return fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(body), putResp)
	}
	if checkSum != putResp.Payload.Hash {
		return fmt.Errorf("url %s crc32 not match, got: %s, want: %s", url, putResp.Payload.Hash, checkSum)
	}
	return nil
}

type segmentedUploadParam struct {
	host string
	StoreInfo
	content     io.Reader
	size        int64
	isLargeFile bool
}

func (c *segmentedUploadParam) chunkUpload() error {
	uploadID, err := c.initUploadPart()
	if err != nil {
		return err
	}
	cur := make([]byte, MinChunkSize)
	parts := make([]string, 0)

	num := c.size / MinChunkSize
	cnt := 0
	lastNum := int(num) - 1

	// 读 n-1 片并上传上去
	var part string
	for i := 0; i < lastNum; i++ {
		n, err := io.ReadFull(c.content, cur)
		if err != nil {
			return err
		}
		cnt += n
		partNumber := i
		if c.isLargeFile {
			partNumber++
		}
		err = retry.Do(func() error {
			part, err = c.uploadPart(uploadID, partNumber, cur)
			return err
		}, retry.Attempts(3))
		if err != nil {
			return err
		}
		parts = append(parts, part)
	}

	// 读 n 和 n+1片（如有）上传上去
	//goland:noinspection GoDeprecation
	bts, err := ioutil.ReadAll(c.content)
	if err != nil {
		return err
	}
	total := len(bts) + cnt
	if total != int(c.size) {
		return fmt.Errorf(fmt.Sprintf("last part download size mismatch ,download %d , size %d", total, c.size))
	}
	if c.isLargeFile {
		lastNum++
	}
	err = retry.Do(func() error {
		part, err = c.uploadPart(uploadID, lastNum, bts)
		return err
	}, retry.Attempts(3))
	if err != nil {
		return err
	}
	parts = append(parts, part)
	return c.uploadMergePart(uploadID, parts)
}

func (c *segmentedUploadParam) initUploadPart() (string, error) {
	url := fmt.Sprintf("https://%s/%s?uploads", c.host, c.StoreUri)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", c.Auth)
	if c.isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}

	client := http.DefaultClient
	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	//goland:noinspection GoDeprecation
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}

	res := &uploadPartResponse{}
	res.Success = -1
	if err := json.Unmarshal(b, res); err != nil {
		return "", fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(b), err)
	}
	if res.Success != 0 {
		return "", fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(b), res)
	}
	return res.PayLoad.UploadID, nil
}

func (c *segmentedUploadParam) uploadPart(uploadID string, partNumber int, data []byte) (string, error) {
	url := fmt.Sprintf("https://%s/%s?partNumber=%d&uploadID=%s", c.host, c.StoreUri, partNumber, uploadID)
	checkSum := fmt.Sprintf("%08x", crc32.ChecksumIEEE(data))
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-CRC32", checkSum)
	req.Header.Set("Authorization", c.Auth)
	if c.isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}

	client := http.DefaultClient
	rsp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	//goland:noinspection GoDeprecation
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return "", err
	}

	res := &uploadPartResponse{}
	res.Success = -1

	if err := json.Unmarshal(b, res); err != nil {
		return "", fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(b), err)
	}
	if res.Success != 0 {
		return "", fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(b), res)
	}
	return checkSum, nil
}

func (c *segmentedUploadParam) uploadMergePart(uploadID string, checkSum []string) error {
	url := fmt.Sprintf("https://%s/%s?uploadID=%s", c.host, c.StoreUri, uploadID)
	body, err := c.genMergeBody(checkSum)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", c.Auth)
	if c.isLargeFile {
		req.Header.Set("X-Storage-Mode", "gateway")
	}

	client := http.DefaultClient
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	//goland:noinspection GoDeprecation
	b, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return err
	}

	res := &uploadMergeResponse{}
	res.Success = -1
	if err := json.Unmarshal(b, res); err != nil {
		return fmt.Errorf("fail to unmarshal response url %s, body: %s, err:  %v", url, string(b), err)
	}
	if res.Success != 0 {
		return fmt.Errorf("fail to put url %s, body %s, err: %+v", url, string(b), res)
	}
	return nil
}

func (c *segmentedUploadParam) genMergeBody(checkSum []string) (string, error) {
	if len(checkSum) == 0 {
		return "", fmt.Errorf("body crc32 empty")
	}
	s := make([]string, len(checkSum))
	for partNumber, crc := range checkSum {
		s[partNumber] = fmt.Sprintf("%d:%s", partNumber, crc)
	}
	return strings.Join(s, ","), nil
}
