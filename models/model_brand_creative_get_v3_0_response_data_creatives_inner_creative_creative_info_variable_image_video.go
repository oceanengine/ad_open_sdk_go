/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVariableImageVideo 尺寸可变视频图文素材
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVariableImageVideo struct {
	Image *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVariableImageVideoImage `json:"image,omitempty"`
	// 标题
	Title *string                                                                                   `json:"title,omitempty"`
	Video *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVariableImageVideoVideo `json:"video,omitempty"`
}
