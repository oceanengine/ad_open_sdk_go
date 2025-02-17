/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoExploreOneTimeComponent 探索一下组件
type BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoExploreOneTimeComponent struct {
	// 组件弹出时间
	AdvancedDuration *int64                                                                                                    `json:"advanced_duration,omitempty"`
	WipeAfterImage   *BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoExploreOneTimeComponentWipeAfterImage  `json:"wipe_after_image,omitempty"`
	WipeBeforeImage  *BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoExploreOneTimeComponentWipeBeforeImage `json:"wipe_before_image,omitempty"`
	WipeImage        *BrandCreativeCreateV30RequestCreativeListInnerAdvancedCreativeInfoExploreOneTimeComponentWipeImage       `json:"wipe_image,omitempty"`
	// 擦一擦文案
	WipeText *string `json:"wipe_text,omitempty"`
}
