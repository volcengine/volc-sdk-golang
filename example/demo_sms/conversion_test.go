package demo_sms

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

/*
return message_ids(the message_ids are the ones in Send/BatchSend response body) which have converted
*/
func TestSmsConversion(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)

	req := &sms.ConversionRequest{
		MessageIDs: []string{"MessageID"},
	}
	result, statusCode, err := sms.DefaultInstance.Conversion(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
