/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseStopV30Request struct for ToolsPromotionRaiseStopV30Request
type ToolsPromotionRaiseStopV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 广告ID列表，目前仅支持1个
	PromotionIds []int64 `json:"promotion_ids"`
}
