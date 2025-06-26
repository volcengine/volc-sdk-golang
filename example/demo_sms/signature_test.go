package demo_sms

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func TestGetSignatureAndOrderList(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.GetSignatureAndOrderListRequest{
		SubAccount: "smsAccount",
		Signature:  "sign",
		PageIndex:  1,
		PageSize:   10,
	}
	result, statusCode, err := sms.DefaultInstance.GetSignatureAndOrderList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsSignature(t *testing.T) {
	fileBase64String, _ := getBase64StringFromFile("/Users//Pictures/IMG_9033.JPG")

	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.ApplySmsSignatureRequest{
		Desc:       "测试 SDK",
		SubAccount: "smsAccount",
		Content:    "sign",
		Domain:     "http链接",
		Source:     sms.SignSourceTypeBrand,
		UploadFileList: []sms.SignAuthFile{
			{
				FileType:    sms.DocTypeThreeInOne,
				FileContent: fileBase64String,
				FileSuffix:  "jpg",
			},
		},
		Purpose: sms.SignPurposeForOwn,
	}
	result, statusCode, err := sms.DefaultInstance.ApplySmsSignature(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteSignature(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.DeleteSignatureRequest{
		SubAccount: "smsAccount",
		Id:         "idOfSignatureToDelete",
		IsOrder:    true,
	}
	result, statusCode, err := sms.DefaultInstance.DeleteSignature(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySmsSignatureV2(t *testing.T) {
	fileBase64String, err := getBase64StringFromFile("/Users//Pictures/IMG_9033.JPG")
	if err != nil {
		t.Logf("err = %+v\n", err)
		return
	}
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.ApplySmsSignatureRequestV2{
		Desc:                      "测试 SDK",
		SubAccount:                "smsAccount",
		Content:                   "sign",
		Source:                    sms.SignSourceCompany,
		SignatureIdentificationID: 123,
		Purpose:                   sms.SignPurposeForOwn,
		AppIcp: sms.AppIcp{
			AppIcpFilling: "appName",
			AppIcpFileList: []sms.SignAuthFile{
				{
					FileType:    sms.DocTypeAppIcpCertificate,
					FileContent: fileBase64String,
					FileSuffix:  "jpg",
				},
			},
		},
		Trademark: sms.Trademark{
			TrademarkCn:     "商标中文名",
			TrademarkEn:     "商标英文名",
			TrademarkNumber: "商标编号",
			TrademarkFileList: []sms.SignAuthFile{
				{
					FileType:    sms.DocTypeTrademarkCertificate,
					FileContent: fileBase64String,
					FileSuffix:  "jpg",
				},
			},
		},
		Scene: "scene",
	}
	result, statusCode, err := sms.DefaultInstance.ApplySmsSignatureV2(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateSmsSignature(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testAk)
	req := &sms.ApplySmsSignatureRequestV2{
		Desc:                      "测试 SDK",
		SubAccount:                "smsAccount",
		Content:                   "sign",
		Source:                    sms.SignSourceApp,
		SignatureIdentificationID: 123,
		Purpose:                   sms.SignPurposeForOwn,
		AppIcp: sms.AppIcp{
			AppIcpFilling: "appName",
			AppIcpFileList: []sms.SignAuthFile{
				{
					FileType:   sms.DocTypeAppIcpCertificate,
					FileUrl:    "公网http链接地址",
					FileSuffix: "jpg",
				},
			},
		},
		Trademark: sms.Trademark{
			TrademarkCn:     "商标中文名",
			TrademarkEn:     "商标英文名",
			TrademarkNumber: "商标编号",
			TrademarkFileList: []sms.SignAuthFile{
				{
					FileType:   sms.DocTypeTrademarkCertificate,
					FileUrl:    "公网http链接地址",
					FileSuffix: "jpg",
				},
			},
		},
		Scene: "test",
	}
	result, statusCode, err := sms.DefaultInstance.UpdateSmsSignature(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestGetSignatureIdentList(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.GetSignatureIdentListRequest{
		Ids:       nil,
		PageIndex: 5,
		PageSize:  1,
	}
	result, statusCode, err := sms.DefaultInstance.GetSignatureIdentList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestApplySignatureIdent(t *testing.T) {
	fileBase64String, err := getBase64StringFromFile("/Users//Pictures/IMG_9033.JPG")
	if err != nil {
		t.Logf("err = %+v\n", err)
		return
	}
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.ApplySignatureIdentRequest{
		Purpose:      1,
		MaterialName: "名字",
		BusinessInfo: sms.BusinessInfo{
			BusinessCertificateType: 1,
			BusinessCertificate: sms.SignAuthFile{
				FileType:    1,
				FileContent: fileBase64String,
				FileSuffix:  "jpeg",
				FileUrl:     "",
			},
			BusinessCertificateName:                "xxx有限公司",
			UnifiedSocialCreditIdentifier:          "1234556",
			BusinessCertificateValidityPeriodStart: "2024-04-18",
			BusinessCertificateValidityPeriodEnd:   "2029-05-18",
			LegalPersonName:                        "黄xx",
		},
		OperatorPersonInfo: sms.PersonInfo{
			CertificateType:   0,
			PersonCertificate: []sms.SignAuthFile{{FileType: 8, FileContent: fileBase64String, FileSuffix: "jpeg"}, {FileType: 9, FileContent: fileBase64String, FileSuffix: "jpeg"}},
			PersonName:        "周xx",
			PersonIDCard:      "1111111",
			PersonMobile:      "15000000000",
		},
		ResponsiblePersonInfo: sms.PersonInfo{
			CertificateType:   0,
			PersonCertificate: []sms.SignAuthFile{{FileType: 10, FileContent: fileBase64String, FileSuffix: "jpeg"}, {FileType: 11, FileContent: fileBase64String, FileSuffix: "jpeg"}},
			PersonName:        "周xx",
			PersonIDCard:      "1111111",
			PersonMobile:      "15000000000",
		},
	}
	result, statusCode, err := sms.DefaultInstance.ApplySignatureIdent(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestBatchBindSignatureIdent(t *testing.T) {
	sms.DefaultInstance.Client.SetAccessKey(testAk)
	sms.DefaultInstance.Client.SetSecretKey(testSk)
	req := &sms.BatchBindSignatureIdentRequest{
		SubAccount: "123456",
		Signatures: []string{"签名"},
		Id:         5,
	}
	result, statusCode, err := sms.DefaultInstance.BatchBindSignatureIdent(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
