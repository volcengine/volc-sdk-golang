package acep

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	ACEP "github.com/volcengine/volc-sdk-golang/service/acep"
)

func Test_AddPropertyRule(t *testing.T) {
	// 强烈建议不要把 VOLC_ACCESSKEY 和 VOLC_SECRETKEY 保存到工程代码里，否则可能导致 AccessKey 泄露，威胁您账号下所有资源的安全。
	// 本示例通过从环境变量中读取 VOLC_ACCESSKEY 和 VOLC_SECRETKEY，来实现 API 访问的身份验证。运行代码示例前，请配置环境变量 VOLC_ACCESSKEY 和 VOLC_SECRETKEY
	service := ACEP.NewInstance()
	service.SetCredential(base.Credentials{
		AccessKeyID:     os.Getenv("VOLC_ACCESSKEY"),
		SecretAccessKey: os.Getenv("VOLC_SECRETKEY"),
	})

	body := &ACEP.AddPropertyRuleBody{
		// 机型参数规则名称，账户下规则名称唯一，长度不能超过 200 个字符。
		RuleName: `Oppo1`,
	}

	// 实例 Settings 系统属性，非持久化，立即生效，重启实例后失效；详细信息，参考 [SystemProperty](#systemproperty) 定义。
	OverlayProperty1 := ACEP.AddPropertyRuleBodyOverlayPropertyItem{
		// 属性名称，已可修改的系统属性及属性值，参考 [System Properties 属性列表](https://www.volcengine.com/docs/6394/671880#system-properties-%E5%B1%9E%E6%80%A7%E5%88%97%E8%A1%A8)。
		SystemPropertyName: `ro.product.name`,
		// 属性值。取值如下：
		// <li> bool: "true"、"false" </li>
		// <li> intlong: "21312" </li>
		// <li> float: "2131.09" </li>
		// <li> string: "safehg" </li>
		// note：
		// 当 SystemPropertyValueType 参数取值不同，该 SystemPropertyValue 参数对应的取值不同。
		SystemPropertyValue: `samsung`,
		// 属性值数据类型。取值如下：
		// <li> int </li>
		// <li> long </li>
		// <li> string </li>
		// <li> float </li>
		// <li> bool </li>
		SystemPropertyValueType: `string`,
	}

	body.OverlayProperty = append(body.OverlayProperty, &OverlayProperty1)

	// 实例初始化系统属性，持久化，重启实例后生效；详细信息，参考 [SystemProperty](#systemproperty ) 定义：适用于只读系统属性，或 AOSP 原生非持久化的系统属性的修改，如 ro.product.model。
	OverlayPersistProperty1 := ACEP.AddPropertyRuleBodyOverlayPersistPropertyItem{
		// 属性名称，已可修改的系统属性及属性值，参考 [System Properties 属性列表](https://www.volcengine.com/docs/6394/671880#system-properties-%E5%B1%9E%E6%80%A7%E5%88%97%E8%A1%A8)。
		SystemPropertyName: `ro.product.name`,
		// 属性值。取值如下：
		// <li> bool: "true"、"false" </li>
		// <li> intlong: "21312" </li>
		// <li> float: "2131.09" </li>
		// <li> string: "safehg" </li>
		// note：
		// 当 SystemPropertyValueType 参数取值不同，该 SystemPropertyValue 参数对应的取值不同。
		SystemPropertyValue: `samsung`,
		// 属性值数据类型。取值如下：
		// <li> int </li>
		// <li> long </li>
		// <li> string </li>
		// <li> float </li>
		// <li> bool </li>
		SystemPropertyValueType: `string`,
	}

	body.OverlayPersistProperty = append(body.OverlayPersistProperty, &OverlayPersistProperty1)

	// 实例 Settings 属性列表；详细信息，参考 [Settings](#settings) 定义。
	OverlaySettings1 := ACEP.AddPropertyRuleBodyOverlaySettingsItem{
		// 属性名称，已可修改的 Settings 属性及属性值，参考 [Settings 属性列表](https://www.volcengine.com/docs/6394/671880#settings-%E5%B1%9E%E6%80%A7%E5%88%97%E8%A1%A8)
		SettingsName: `locale_language`,
		// 设置类型，取值如下：
		// <li> global </li>
		// <li> secure </li>
		// <li> system </li>
		SettingsType: `global`,
		// 属性值。各类型属性取值示例如下：
		// <li> bool: "true"、"false" </li>
		// <li> intlong: "21312" </li>
		// <li> float: "2131.09" </li>
		// <li> string: "safehg" </li>
		SettingsValue: `on`,
		// 属性值数据类型，取值如下：
		// <li> int </li>
		// <li> long </li>
		// <li> string </li>
		// <li> float </li>
		// <li> bool </li>
		SettingsValueType: `string`,
	}

	body.OverlaySettings = append(body.OverlaySettings, &OverlaySettings1)

	resp, err := service.AddPropertyRule(context.Background(), body)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		fmt.Printf("success %+v", resp)
	}
}
