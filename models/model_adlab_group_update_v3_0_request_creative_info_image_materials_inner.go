/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateV30RequestCreativeInfoImageMaterialsInner struct for AdlabGroupUpdateV30RequestCreativeInfoImageMaterialsInner
type AdlabGroupUpdateV30RequestCreativeInfoImageMaterialsInner struct {
	// 图片信息
	ImageInfo []*AdlabGroupUpdateV30RequestCreativeInfoImageMaterialsInnerImageInfoInner `json:"image_info"`
	ImageMode AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode                     `json:"image_mode"`
}
