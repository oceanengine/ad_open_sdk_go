/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerText 文本组件描述
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerText struct {
	// 文本内容，长度至少为1
	Content string                                                      `json:"content"`
	LinkDto *ToolsSiteTemplateSiteCreateV2RequestBricksInnerTextLinkDto `json:"link_dto,omitempty"`
}
