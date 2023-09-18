/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonDownloadEventIosLink ios链接信息
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerButtonDownloadEventIosLink struct {
	// 应用描述，为了展示效果，推荐12个中文字符长度
	Description *string                                                               `json:"description,omitempty"`
	LinkType    ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventIosLinkLinkType `json:"link_type"`
	// 链接地址，当`link_type`为URL时，必填
	Url *string `json:"url,omitempty"`
}