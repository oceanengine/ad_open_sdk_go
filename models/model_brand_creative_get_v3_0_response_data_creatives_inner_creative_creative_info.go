/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfo 基础创意信息
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfo struct {
	CreativeType         *BrandCreativeGetV30DataCreativesCreativeCreativeInfoCreativeType                      `json:"creative_type,omitempty"`
	GroupImage           *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoGroupImage           `json:"group_image,omitempty"`
	HorizontalImageVideo *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoHorizontalImageVideo `json:"horizontal_image_video,omitempty"`
	HorizontalVideo      *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoHorizontalVideo      `json:"horizontal_video,omitempty"`
	LargeImage           *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoLargeImage           `json:"large_image,omitempty"`
	SmallImage           *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoSmallImage           `json:"small_image,omitempty"`
	VariableImageVideo   *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVariableImageVideo   `json:"variable_image_video,omitempty"`
	VerticalImageVideo   *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalImageVideo   `json:"vertical_image_video,omitempty"`
	VerticalVideo        *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVerticalVideo        `json:"vertical_video,omitempty"`
	VideoCreative        *BrandCreativeGetV30ResponseDataCreativesInnerCreativeCreativeInfoVideoCreative        `json:"video_creative,omitempty"`
}
