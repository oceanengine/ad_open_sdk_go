/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalImageVideo 竖版带图视频
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalImageVideo struct {
	Image *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalImageVideoImage `json:"image,omitempty"`
	// 标题
	Title *string                                                                                   `json:"title,omitempty"`
	Video *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalImageVideoVideo `json:"video,omitempty"`
}
