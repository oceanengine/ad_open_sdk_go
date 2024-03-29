/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeTemplateListGetV2ResponseDataListsInner struct for CreativeTemplateListGetV2ResponseDataListsInner
type CreativeTemplateListGetV2ResponseDataListsInner struct {
	// 模板的封面图
	CoverUrl  *string                                      `json:"cover_url,omitempty"`
	ImageMode *CreativeTemplateListGetV2DataListsImageMode `json:"image_mode,omitempty"`
	// 模板预览链接
	PreviewUrl *string `json:"preview_url,omitempty"`
	// 模板描述，用以说明例如\"这个模板需要计划下自提素材中含有竖版视频才可用\"等必要说明信息
	TemplateDesc *string `json:"template_desc,omitempty"`
	// 模板ID
	TemplateId *int64 `json:"template_id,omitempty"`
	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`
	// 模板标签列表
	TemplateTags []*CreativeTemplateListGetV2ResponseDataListsInnerTemplateTagsInner `json:"template_tags,omitempty"`
	TemplateType *CreativeTemplateListGetV2DataListsTemplateType                     `json:"template_type,omitempty"`
}
