/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2Request struct for ToolsSiteTemplateCreateV2Request
type ToolsSiteTemplateCreateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 站点ID，可通过[【橙子建站】]（https://www.chengzijianzhan.com/）平台或[【获取橙子建站站点列表】](https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710620579852#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80)获取
	SiteId int64 `json:"site_id"`
	// 模板名称，可选字段，用户不填写的情况下，会使用原站点名会作为模板名，范围：长度 >= 1
	TemplateName *string `json:"template_name,omitempty"`
}
