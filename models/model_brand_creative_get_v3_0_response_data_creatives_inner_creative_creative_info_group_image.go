/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoGroupImage 组图
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoGroupImage struct {
	// 图片
	Image []*BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoGroupImageImageInner `json:"image,omitempty"`
	// 标题
	Title *string `json:"title,omitempty"`
}