/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionNewcustomerCreateV30RequestPromotionMaterialsVideoMaterialListInner struct for PromotionNewcustomerCreateV30RequestPromotionMaterialsVideoMaterialListInner
type PromotionNewcustomerCreateV30RequestPromotionMaterialsVideoMaterialListInner struct {
	ImageMode *PromotionNewcustomerCreateV30PromotionMaterialsVideoMaterialListImageMode `json:"image_mode,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	// 视频封面图片uri
	VideoCoverUri *string `json:"video_cover_uri,omitempty"`
	// 视频id
	VideoId *string `json:"video_id,omitempty"`
}
