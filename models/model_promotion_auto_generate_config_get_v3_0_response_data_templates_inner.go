/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAutoGenerateConfigGetV30ResponseDataTemplatesInner struct for PromotionAutoGenerateConfigGetV30ResponseDataTemplatesInner
type PromotionAutoGenerateConfigGetV30ResponseDataTemplatesInner struct {
	// 模板ID
	TemplateId int64 `json:"template_id"`
	// 模板填充的图片内容
	TemplateImgSchema []*PromotionAutoGenerateConfigGetV30ResponseDataTemplatesInnerTemplateImgSchemaInner `json:"template_img_schema,omitempty"`
	// 模板填充的文本内容
	TemplateTextSchema []*PromotionAutoGenerateConfigGetV30ResponseDataTemplatesInnerTemplateTextSchemaInner `json:"template_text_schema,omitempty"`
	TemplateType       PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType                            `json:"template_type"`
}
