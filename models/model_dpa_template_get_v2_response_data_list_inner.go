/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaTemplateGetV2ResponseDataListInner struct for DpaTemplateGetV2ResponseDataListInner
type DpaTemplateGetV2ResponseDataListInner struct {
	//
	Industry *int64 `json:"industry,omitempty"`
	//
	IsPublic *bool `json:"is_public,omitempty"`
	//
	TemplateDataList []*DpaTemplateGetV2ResponseDataListInnerTemplateDataListInner `json:"template_data_list,omitempty"`
	//
	TemplateId   *int64                                `json:"template_id,omitempty"`
	TemplateMode *DpaTemplateGetV2DataListTemplateMode `json:"template_mode,omitempty"`
	//
	TemplateName *string `json:"template_name,omitempty"`
}
