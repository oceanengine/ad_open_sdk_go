/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataCreativeInfoImageMaterialsInner struct for AdlabGroupDetailV30ResponseDataDataCreativeInfoImageMaterialsInner
type AdlabGroupDetailV30ResponseDataDataCreativeInfoImageMaterialsInner struct {
	// 图片信息
	ImageInfo []*AdlabGroupDetailV30ResponseDataDataCreativeInfoImageMaterialsInnerImageInfoInner `json:"image_info,omitempty"`
	ImageMode *AdlabGroupDetailV30DataDataCreativeInfoImageMaterialsImageMode                     `json:"image_mode,omitempty"`
}
