/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoCard 图片磁贴或图片磁贴含搜索组件
type BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoCard struct {
	// 组件弹出时间
	AdvancedDuration *int64                                                                           `json:"advanced_duration,omitempty"`
	ImageInfo        *BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoCardImageInfo `json:"image_info,omitempty"`
	// 内容热推引流直播间抖音号
	LiveAwemeId *string `json:"live_aweme_id,omitempty"`
	// 搜索词数组
	SearchWords []string `json:"search_words,omitempty"`
}
