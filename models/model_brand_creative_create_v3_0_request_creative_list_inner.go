/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30RequestCreativeListInner struct for BrandCreativeCreateV30RequestCreativeListInner
type BrandCreativeCreateV30RequestCreativeListInner struct {
	AdvancedCreative *BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreative `json:"advanced_creative,omitempty"`
	// 抖音号ID
	AwemeId *string `json:"aweme_id,omitempty"`
	// 是否AI生成
	CreativeIsAIGC bool `json:"creative_is_AIGC"`
	// 直达链接
	ExternalOpenUrl *string `json:"external_open_url,omitempty"`
	// 落地页链接
	ExternalUrl *string `json:"external_url,omitempty"`
	// 落地页是否包含下载链接
	ExternalUrlIsDownload *bool `json:"external_url_is_download,omitempty"`
	// 落地页标题
	ExternalUrlTitle *string `json:"external_url_title,omitempty"`
	// 抖音号UID
	IesCoreUserId *int64 `json:"ies_core_user_id,omitempty"`
	// 抖音视频ID
	ItemId *int64 `json:"item_id,omitempty"`
	// 素材包ID
	MaterialPackageId *int64 `json:"material_package_id,omitempty"`
	// 购物车无分佣直达链接
	OpenUrl *string `json:"open_url,omitempty"`
	// 创意标题，选择素材包和抖音视频ID时无需填写
	Title *string `json:"title,omitempty"`
	// 视频ID
	VideoId *string `json:"video_id,omitempty"`
	// 购物车无分佣链接
	WebUrl *string `json:"web_url,omitempty"`
	// web_url是否是下载类落地页链接
	WebUrlIsDownload *bool `json:"web_url_is_download,omitempty"`
}
