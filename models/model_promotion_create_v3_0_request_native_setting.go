/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestNativeSetting
type PromotionCreateV30RequestNativeSetting struct {
	AnchorRelatedType *PromotionCreateV30NativeSettingAnchorRelatedType `json:"anchor_related_type,omitempty"`
	//
	AwemeId         *string                                         `json:"aweme_id,omitempty"`
	IsFeedAndFavSee *PromotionCreateV30NativeSettingIsFeedAndFavSee `json:"is_feed_and_fav_see,omitempty"`
}
