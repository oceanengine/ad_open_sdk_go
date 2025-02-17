/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPromotionCreateV30VideoHpVisibility
type LocalPromotionCreateV30VideoHpVisibility string

// List of local_promotion_create_v3.0_video_hp_visibility
const (
	ALWAYS_VISIBLE_LocalPromotionCreateV30VideoHpVisibility   LocalPromotionCreateV30VideoHpVisibility = "ALWAYS_VISIBLE"
	HIDE_VIDEO_ON_HP_LocalPromotionCreateV30VideoHpVisibility LocalPromotionCreateV30VideoHpVisibility = "HIDE_VIDEO_ON_HP"
)

// Ptr returns reference to local_promotion_create_v3.0_video_hp_visibility value
func (v LocalPromotionCreateV30VideoHpVisibility) Ptr() *LocalPromotionCreateV30VideoHpVisibility {
	return &v
}
