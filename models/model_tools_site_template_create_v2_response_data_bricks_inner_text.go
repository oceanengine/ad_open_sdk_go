/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerText 文本组件描述
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerText struct {
	// 文本内容，长度至少为1
	Content string                                                       `json:"content"`
	LinkDto *ToolsSiteTemplateCreateV2ResponseDataBricksInnerTextLinkDto `json:"link_dto,omitempty"`
}
