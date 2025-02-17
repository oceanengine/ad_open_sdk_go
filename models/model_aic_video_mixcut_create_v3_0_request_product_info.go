/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicVideoMixcutCreateV30RequestProductInfo 商品信息，用于生成脚本文案
type AicVideoMixcutCreateV30RequestProductInfo struct {
	// 利益点，总字数上限：100，示例：[七夕全场包邮，满300立减50] 字数14
	BenefitPoints []string `json:"benefit_points,omitempty"`
	// 商品名，尽量简练，且需要包括商品主体，比如插座、空调等，字数上限：20 推荐输入：立白洗衣凝珠 不推荐输入：88rising龙年限定双龙刺绣麂皮绒摇粒绒
	ProductName string `json:"product_name"`
	// 商品卖点，总字数上限：100，示例：[加大加厚，吸油锁水] 字数8
	ProductSellingPoints []string `json:"product_selling_points"`
}
