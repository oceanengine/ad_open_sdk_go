/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoLiveCard 直播卡片
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoLiveCard struct {
	// 副标题
	AdvancedSubtitle *string `json:"advanced_subtitle,omitempty"`
	// 标题
	AdvancedTitle *string                                                                                     `json:"advanced_title,omitempty"`
	ImageInfo     *BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoLiveCardImageInfo `json:"image_info,omitempty"`
}
