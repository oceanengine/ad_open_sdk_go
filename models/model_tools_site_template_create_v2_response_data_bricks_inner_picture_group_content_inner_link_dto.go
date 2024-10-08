/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureGroupContentInnerLinkDto 跳转链接信息
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureGroupContentInnerLinkDto struct {
	LinkType ToolsSiteTemplateCreateV2DataBricksPictureGroupContentLinkDtoLinkType `json:"link_type"`
	// 快应用地址，当`link_type`等于`QUICK_APP`时，有值
	QuickApp *string `json:"quick_app,omitempty"`
	// scheme地址，当`link_type`等于`SCHEME`时，有值
	Scheme *string `json:"scheme,omitempty"`
	// 链接地址，当`link_type`等于`URL`时，有值
	Url *string `json:"url,omitempty"`
}
