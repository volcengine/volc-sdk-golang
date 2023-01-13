package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual"
)

func main() {
	testAk := "ak"
	testSk := "sk"

	visual.DefaultInstance.Client.SetAccessKey(testAk)
	visual.DefaultInstance.Client.SetSecretKey(testSk)
	//visual.DefaultInstance.SetRegion("region")
	//visual.DefaultInstance.SetHost("host")

	form := url.Values{}
	form.Add("image_base64", "")
	//form.Add("version", "v2") // version=v1 or v2

	var action string
	action = "OCRNormal"          // 通用OCR
	action = "MultiLanguageOCR"   // 多语种OCR
	action = "BankCard"           // 银行卡识别，有v1,v2版本
	action = "IDCard"             // 身份证识别，有v1,v2版本
	action = "DrivingLicense"     // 驾驶证识别
	action = "VehicleLicense"     // 行驶证识别
	action = "OcrTaibao"          // 台湾居民来往大陆通行证识别
	action = "OcrVatInvoice"      // 增值税发票识别
	action = "OcrTaxiInvoice"     // 出租车发票识别
	action = "OcrQuotaInvoice"    // 定额发票识别
	action = "OcrTrainTicket"     // 火车票识别
	action = "OcrFlightInvoice"   // 飞机行程单识别
	action = "OcrFinance"         // 混贴报销场景
	action = "OcrRollInvoice"     // 增值税卷票识别
	action = "OcrPassInvoice"     // 高速公路过路费票识别
	action = "OcrFoodProduction"  // 食品生产许可证识别
	action = "OcrFoodBusiness"    // 食品经营许可证识别
	action = "OcrClueLicense"     // 营业执照识别
	action = "OCRTrade"           // 商标证识别
	action = "OCRRuanzhu"         // 软件著作权识别
	action = "OCRCosmeticProduct" // 化妆品生产许可证识别
	action = "OcrSeal"            // 印章识别
	action = "OcrTextAlignment"   // 合同校验
	action = "OCRPdf"             // PDF识别
	action = "OCRTable"           // 表格识别

	resp, status, err := visual.DefaultInstance.OCRApi(form, action)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
