/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeListInnerImageMaterialsInner struct for CreativeDetailGetV30ResponseDataCreativeListInnerImageMaterialsInner
type CreativeDetailGetV30ResponseDataCreativeListInnerImageMaterialsInner struct {
	// 图片素材信息
	ImageInfo     []*CreativeDetailGetV30ResponseDataCreativeListInnerImageMaterialsInnerImageInfoInner `json:"image_info,omitempty"`
	TemplateImage *CreativeDetailGetV30ResponseDataCreativeListInnerImageMaterialsInnerTemplateImage    `json:"template_image,omitempty"`
}
