/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionBidUpdateV30Request struct for PromotionBidUpdateV30Request
type PromotionBidUpdateV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 批量更新列表
	Data []*PromotionBidUpdateV30RequestDataInner `json:"data"`
}
