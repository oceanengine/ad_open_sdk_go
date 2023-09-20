/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseData
type ToolsSiteTemplateCreateV2ResponseData struct {
	// 组件列表
	Bricks []*ToolsSiteTemplateCreateV2ResponseDataBricksInner `json:"bricks,omitempty"`
	// 站点ID，可通过[【橙子建站】（https://www.chengzijianzhan.com/）]平台或[【获取橙子建站站点列表】](https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710620579852#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80)获取
	SiteId *int64 `json:"site_id,omitempty"`
	// 模板ID
	TemplateId *int64 `json:"template_id,omitempty"`
	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`
}
