/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2Request struct for ToolsSiteTemplateSiteCreateV2Request
type ToolsSiteTemplateSiteCreateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 组件列表
	Bricks []*ToolsSiteTemplateSiteCreateV2RequestBricksInner `json:"bricks"`
	// 站点名称，范围：长度 >= 1
	Name string `json:"name"`
	// 站点ID，未填充时会创建新的站点。填充时，会对当前站点的落地页进行修改，可通过[【橙子建站】（https://www.chengzijianzhan.com/）]平台或[【获取橙子建站站点列表】](https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710620579852#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80)获取
	SiteId *int64 `json:"site_id,omitempty"`
	// 模板ID，【基于站点创建模板】接口，返回的模板ID
	TemplateId int64 `json:"template_id"`
}
