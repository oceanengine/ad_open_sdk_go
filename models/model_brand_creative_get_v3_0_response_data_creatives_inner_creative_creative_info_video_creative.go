/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVideoCreative 视频不区分横竖版
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVideoCreative struct {
	Image *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVideoCreativeImage `json:"image,omitempty"`
	// 标题
	Title *string                                                                              `json:"title,omitempty"`
	Video *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVideoCreativeVideo `json:"video,omitempty"`
}
