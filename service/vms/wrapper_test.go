package vms

import (
	"testing"
)

const (
	testAk = "your ak"
	testSk = "your sk"
)

func init() {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
}

func TestBindAXB(t *testing.T) {
	req := &BindAXBRequest{
		NumberPoolNo: "NP161156328504091435",
		PhoneNoA:     "188xxxx1647",
		PhoneNoB:     "137xxxx8258",
		PhoneNoX:     "170xxxx3159",
		ExpireTime:   1632313906,
	}
	result, statusCode, err := DefaultInstance.BindAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSelectNumberAndBindAXBindAXB(t *testing.T) {
	req := &SelectNumberAndBindAXBRequest{
		NumberPoolNo:    "NP161156328504091435",
		PhoneNoA:        "188xxxx1647",
		PhoneNoB:        "137xxxx8258",
		ExpireTime:      1632313906,
		CityCode:        "010",
		DegradeCityList: "010,020",
	}
	result, statusCode, err := DefaultInstance.SelectNumberAndBindAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUnbindAXB(t *testing.T) {
	req := &SpecificSubIdRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S1632298175315938db419",
	}
	result, statusCode, err := DefaultInstance.UnbindAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQuerySubscription(t *testing.T) {
	req := &SpecificSubIdRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S163229841631591737db3",
	}
	result, statusCode, err := DefaultInstance.QuerySubscription(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQuerySubscriptionForList(t *testing.T) {
	req := &QuerySubscriptionForListRequest{
		NumberPoolNo: "NP161156328504091435",
		PhoneNoX:     "170xxxx3159",
		Status:       1,
		Offset:       0,
		Limit:        20,
	}
	result, statusCode, err := DefaultInstance.QuerySubscriptionForList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpgradeAXToAXB(t *testing.T) {
	req := &UpgradeAXToAXBRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S16323073363159890f81b",
		PhoneNoB:     "131xxxx7582",
	}
	result, statusCode, err := DefaultInstance.UpgradeAXToAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateAXB(t *testing.T) {
	req := &UpdateAXBRequest{
		NumberPoolNo: "NP161156328504091435",
		SubId:        "S16323075803159b97ba7a",
		UpdateType:   "updateExpireTime",
		ExpireTime:   1632313906,
	}
	result, statusCode, err := DefaultInstance.UpdateAXB(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestBindAXN(t *testing.T) {
	req := &BindAXNRequest{
		NumberPoolNo: "NP162981168404095092",
		PhoneNoA:     "188xxxx1647",
		PhoneNoX:     "170xxxx8991",
		PhoneNoB:     "137xxxx8258",
		ExpireTime:   1632313906,
	}
	result, statusCode, err := DefaultInstance.BindAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestSelectNumberAndBindAXN(t *testing.T) {
	req := &SelectNumberAndBindAXNRequest{
		NumberPoolNo: "NP162981168404095092",
		PhoneNoA:     "188xxxx5753",
		PhoneNoB:     "137xxxx8257",
		ExpireTime:   1642672859,
	}
	result, statusCode, err := DefaultInstance.SelectNumberAndBindAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateAXN(t *testing.T) {
	req := &UpdateAXNRequest{
		NumberPoolNo: "NP162981168404095092",
		SubId:        "S16323109008991825a8b7",
		UpdateType:   "updatePhoneNoA",
		PhoneNoA:     "188xxxx5753",
	}
	result, statusCode, err := DefaultInstance.UpdateAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUnbindAXN(t *testing.T) {
	req := &SpecificSubIdRequest{
		NumberPoolNo: "NP162981168404095092",
		SubId:        "S16323109008991825a8b7",
	}
	result, statusCode, err := DefaultInstance.UnbindAXN(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestBindAXG(t *testing.T) {
	req := &BindAXGRequest{
		NumberPoolNo:      "NP1760xxxxx0904269",
		PhoneNoA:          "19200486749",
		PhoneNoB:          "19200486751",
		ExpireTime:        1761908994,
		UserData:          "userdata",
		OutId:             "outId",
		GroupName:         "ces123",
		GroupNumbers:      []string{"19200486752", "19200486753"},
		CityCode:          "010",
		CityCodeByPhoneNo: "A",
		DegradeCityList:   []string{"010", "020"},
		RandomFlag:        1,
		OrderId:           "orderId",
		IndustrialId:      "industrial",
	}
	result, statusCode, err := DefaultInstance.BindAXG(req)
	t.Logf("result = %+v\n %+v\n %+v\n", result, statusCode, err)
	t.Logf("resultError = %+v\n", result.ResponseMetadata.Error)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUnbindAXG(t *testing.T) {
	req := &UnbindAXGRequest{
		NumberPoolNo: "NP176xxxx010904269",
		SubId:        "SS1760945125168448b6w66orfqsi0",
		DelGroup:     true,
	}
	result, statusCode, err := DefaultInstance.UnbindAXG(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateAXG(t *testing.T) {
	req := &UpdateAXGRequest{
		NumberPoolNo: "NP1760xxxx2010904269",
		SubId:        "SS17618226141a981806w66orfqze0",
		UpdateType:   "updateIndustrialId",
		GroupId:      "G1760945070085e79fe95c-f4a9-47e4-a337-d72f86b9f494",
		PhoneNoA:     "13596013572",
		ExpireTime:   1761290038,
		UserData:     "updagte",
		OutId:        "oute123",
		OrderId:      "order123",
		IndustrialId: "industrial123",
	}
	result, statusCode, err := DefaultInstance.UpdateAXG(req)
	t.Logf("result = %+v\n", result)
	t.Logf("resultError = %+v\n", result.ResponseMetadata.Error)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestCreateAXGGroup(t *testing.T) {
	req := &CreateAXGGroupRequest{
		NumberPoolNo: "NP176xxxx52010904269",
		GroupName:    "测试G组名称",
		Numbers:      []string{"19200486752", "19200486753"},
		ExpireTime:   1761030837,
	}
	result, statusCode, err := DefaultInstance.CreateAXGGroup(req)
	t.Logf("result = %+v\n", result)
	t.Logf("resultError = %+v\n", result.ResponseMetadata.Error)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateAXGGroup(t *testing.T) {
	req := &UpdateAXGGroupRequest{
		NumberPoolNo:     "NP17601xxxx010904269",
		SubId:            "",
		GroupId:          "G17609444514630dd22ce5-f51c-4689-9225-1952678aaabe",
		UpdateType:       "addNumbers",
		GroupName:        "测试G组名称123",
		Numbers:          []string{"19200486755", "19200486753", "19200486754", "19200486755", "19200486756", "19200486757", "19200486758", "19200486759", "19200486760", "19200486761", "19200486762", "19200486763"},
		OverflowStrategy: "fifo",
		ExpireTime:       1761030838,
	}
	result, statusCode, err := DefaultInstance.UpdateAXGGroup(req)
	t.Logf("result = %+v\n", result)
	t.Logf("resultError = %+v\n", result.ResponseMetadata.Error)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestDeleteAXGGroup(t *testing.T) {
	req := &DeleteAXGGroupRequest{
		NumberPoolNo: "NP1760xxxx2010904269",
		GroupId:      "G17609444514630dd22ce5-f51c-4689-9225-1952678aaabe",
	}
	result, statusCode, err := DefaultInstance.DeleteAXGGroup(req)
	t.Logf("result = %+v\n", result)
	t.Logf("resultError = %+v\n", result.ResponseMetadata.Error)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestClick2Call(t *testing.T) {
	req := &Click2CallRequest{
		Caller:             "137XXXX8257",
		Callee:             "158XXXX9130",
		CallerNumberPoolNo: "NP163517154204092175",
		CalleeNumberPoolNo: "NP163517154204092175",
	}
	result, statusCode, err := DefaultInstance.Click2Call(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestCancelClick2Call(t *testing.T) {
	req := &CancelClick2CallRequest{
		CallId: "Dccfebdedfe",
	}
	result, statusCode, err := DefaultInstance.CancelClick2Call(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestClick2CallLite(t *testing.T) {
	req := &Click2CallLiteRequest{
		Caller:       "137XXXX8257",
		Callee:       "158XXXX9130",
		NumberPoolNo: "NPXXXXX810901043",
	}
	result, statusCode, err := DefaultInstance.Click2CallLite(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQueryAudioRecordFileUrl(t *testing.T) {
	req := &QueryAudioRecordFileUrlRequest{
		CallId: "*",
	}
	result, statusCode, err := DefaultInstance.QueryAudioRecordFileUrl(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQueryAudioRecordToTextFileUrl(t *testing.T) {
	req := &QueryAudioRecordToTextFileRequest{
		CallIdList: "Vcc01b1fe30f4868c4b7c9742bdd036f12559,Vcc017784d0be628542aa9a676af3d8fa2e06",
	}
	result, statusCode, err := DefaultInstance.QueryAudioRecordToTextFileUrl(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQueryCallRecordMsg(t *testing.T) {
	req := &QueryCallRecordMsgRequest{
		CallIdList: "S164275060051193662574_1dc1f1b3a8891a8b,S1015_c13f7b27b41e7773",
	}
	result, statusCode, err := DefaultInstance.QueryCallRecordMsg(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestCreateNumberPool(t *testing.T) {
	req := &CreateNumberPoolRequest{
		Name:           "测试创建号码池",
		ServiceType:    "100",
		SubServiceType: "102",
	}
	result, statusCode, err := DefaultInstance.CreateNumberPool(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateNumberPool(t *testing.T) {
	req := &UpdateNumberPoolRequest{
		NumberPoolNo:   "NP164612245301914680",
		Name:           "测试创建号码池2",
		ServiceType:    "200",
		SubServiceType: "201",
	}
	result, statusCode, err := DefaultInstance.UpdateNumberPool(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberPoolList(t *testing.T) {
	req := &NumberPoolListRequest{
		SubServiceType: "101",
		Limit:          "10",
		Offset:         "0",
	}
	result, statusCode, err := DefaultInstance.NumberPoolList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestNumberList(t *testing.T) {
	req := &NumberListRequest{
		Number: "17192359311",
		Limit:  "3",
		Offset: "0",
	}
	result, statusCode, err := DefaultInstance.NumberList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestEnableOrDisableNumber(t *testing.T) {
	req := &EnableOrDisableNumberRequest{
		NumberList: "18792770474,18792770475",
		EnableCode: "1",
	}

	result, statusCode, err := DefaultInstance.EnableOrDisableNumber(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQueryNumberApplyRecordList(t *testing.T) {
	req := &QueryNumberApplyRecordListRequest{
		ApplyBillId: "NA5345832249",
		Limit:       "10",
		Offset:      "0",
	}

	result, statusCode, err := DefaultInstance.QueryNumberApplyRecordList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestCreateNumberApplication(t *testing.T) {

	detailList := []NumberApplicationCityItem{
		{
			CityCode:       "010",
			CountryIsoCode: "CN",
			NumberCount:    "1",
		},
	}

	req := &CreateNumberApplicationRequest{
		QualificationNo:               "QUA164679537228220075",
		NumberPoolNo:                  "NP164015715228226293",
		NumberPurpose:                 "1",
		NumberType:                    "1",
		SubServiceType:                "101",
		Remark:                        "remark3",
		NumberApplicationCityItemList: detailList,
	}
	result, statusCode, err := DefaultInstance.CreateNumberApplication(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestAddQualification(t *testing.T) {
	qualificationMainInfoFormDO := QualificationMainInfoFormDO{
		QualificationEntity:                            "3",
		CertificateThreeInOne:                          1,
		EnterpriseAddress:                              "222",
		LegalRepresentativeName:                        "2",
		LegalRepresentativeId:                          "2",
		UnitSocialCreditCode:                           "2",
		LegalRepresentativeFrontIdPhotoFileCode:        "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		DocOfNumberApplyPhotoFileCode:                  "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CommitmentLetterOfNetAccessPhotoFileCode:       "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		ThreeInOneBusinessLicensePhotoFileCode:         "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CodeOfOrganizationCertificate:                  "2",
		BusinessLicensePhotoFileCode:                   "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfOrganizationCodesPhotoFileCode:    "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfTaxationRegistrationPhotoFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationAdminInfoFormDO := QualificationAdminInfoFormDO{
		Name:                          "2",
		ContactNumber:                 "2",
		IdCardNumber:                  "2",
		IdCardFrontPhotoFileCode:      "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		IdCardPhotoWithPeopleFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			SceneType:         1001,
			Description:       "2",
			ScenarioOfCalling: "2",
		},
	}

	req := &AddQualificationRequest{
		QualificationMainInfoFormDO:         qualificationMainInfoFormDO,
		QualificationAdminInfoFormDO:        qualificationAdminInfoFormDO,
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultInstance.AddQualification(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateQualification(t *testing.T) {
	qualificationMainInfoFormDO := QualificationMainInfoFormDO{
		QualificationEntity:                            "3",
		QualificationNo:                                "QUA164664762128220429",
		CertificateThreeInOne:                          1,
		EnterpriseAddress:                              "22",
		LegalRepresentativeName:                        "2",
		LegalRepresentativeId:                          "2",
		UnitSocialCreditCode:                           "2",
		LegalRepresentativeFrontIdPhotoFileCode:        "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		DocOfNumberApplyPhotoFileCode:                  "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CommitmentLetterOfNetAccessPhotoFileCode:       "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		ThreeInOneBusinessLicensePhotoFileCode:         "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CodeOfOrganizationCertificate:                  "2",
		BusinessLicensePhotoFileCode:                   "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfOrganizationCodesPhotoFileCode:    "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		CertificateOfTaxationRegistrationPhotoFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationAdminInfoFormDO := QualificationAdminInfoFormDO{
		Name:                          "2",
		ContactNumber:                 "2",
		IdCardNumber:                  "2",
		IdCardFrontPhotoFileCode:      "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
		IdCardPhotoWithPeopleFileCode: "21000428220368035c-864c-4309-94d7-fd0f2c5d16efLegalRepresentativeFrontIDPhoto.jpeg",
	}

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			SceneType:         1001,
			Description:       "2",
			ScenarioOfCalling: "2",
		},
	}

	req := &UpdateQualificationRequest{
		QualificationMainInfoFormDO:         qualificationMainInfoFormDO,
		QualificationAdminInfoFormDO:        qualificationAdminInfoFormDO,
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultInstance.UpdateQualification(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestAddQualificationScene(t *testing.T) {

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			QualificationNo:   "QUA164664762128220429",
			SceneType:         1001,
			Description:       "22212",
			ScenarioOfCalling: "2222",
		},
	}

	req := &AddQualificationSceneRequest{
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultInstance.AddQualificationScene(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestUpdateQualificationScene(t *testing.T) {

	qualificationScenarioInfoFormDOList := []QualificationScenarioInfoFormDO{
		{
			QualificationNo:   "QUA164664762128220429",
			SceneType:         1001,
			Description:       "2221",
			ScenarioOfCalling: "2222",
		},
	}

	req := &UpdateQualificationSceneRequest{
		QualificationScenarioInfoFormDOList: qualificationScenarioInfoFormDOList,
	}

	result, statusCode, err := DefaultInstance.UpdateQualificationScene(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQueryQualification(t *testing.T) {

	qualificationNoList := []string{"QUA164679537228220075"}

	req := &QueryQualificationRequest{
		QualificationNoList: qualificationNoList,
		ApprovalStatus:      "2",
		Limit:               20,
	}

	result, statusCode, err := DefaultInstance.QueryQualification(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestQuerySingleInfo(t *testing.T) {
	response, statusCode, err := DefaultInstance.SingleInfo("832a5c905c98480ba764928f3cd98c28")
	t.Logf("response = %+v\n", response)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)

	item := &SingleParam{SingleOpenId: "adfadsfds", Phone: "19900000000", NumberPoolNo: "NP163271249910906158", NumberType: 0, Type: 3, TtsContent: "abcd"}
	request := &SingleAppendRequest{List: []*SingleParam{item}}
	response1, statusCode, err := DefaultInstance.SingleBatchAppend(request)
	t.Logf("response = %+v\n", response1)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestRisk_CanCall(t *testing.T) {
	req := &RiskControlReq{
		CustomerNumberList: "188xxxx1647",
		BusinessLineId:     "abc",
		CallType:           1,
	}
	result, statusCode, err := DefaultInstance.QueryCanCall(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestRisk_AddBlackList(t *testing.T) {
	req := &AddBlackListReq{
		CustomerNumber: "188xxxxxxxx",
		ServiceType:    100,
		SubServiceType: 101,
		Type:           1,
		FreezeTime:     12,
		Remark:         "测试api加黑",
	}
	result, statusCode, err := DefaultInstance.AddBlackList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestRisk_DeleteBlackList(t *testing.T) {
	req := &DeleteBlackListReq{
		CustomerNumber: "188xxxxxxxx",
		ServiceType:    100,
		SubServiceType: 101,
		Type:           1,
	}
	result, statusCode, err := DefaultInstance.DeleteBlackList(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}

func TestRouteAAuth(t *testing.T) {
	req := &RouteAAuthRequest{
		RequestId:   "sadasdasewr",
		IdType:      1,
		Name:        "张三",
		PhoneNumber: "138000000xx",
		IdNumber:    "220121xxx",
		Photo:       "123211231232432",
	}
	result, statusCode, err := DefaultInstance.RouteAAuth(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
