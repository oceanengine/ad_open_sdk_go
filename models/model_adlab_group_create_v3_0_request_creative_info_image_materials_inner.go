/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestCreativeInfoImageMaterialsInner struct for AdlabGroupCreateV30RequestCreativeInfoImageMaterialsInner
type AdlabGroupCreateV30RequestCreativeInfoImageMaterialsInner struct {
	// 图片信息
	ImageInfo []*AdlabGroupCreateV30RequestCreativeInfoImageMaterialsInnerImageInfoInner `json:"image_info"`
	ImageMode AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode                     `json:"image_mode"`
}
