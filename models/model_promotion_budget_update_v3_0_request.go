/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionBudgetUpdateV30Request struct for PromotionBudgetUpdateV30Request
type PromotionBudgetUpdateV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 批量修改预算列表
	Data []*PromotionBudgetUpdateV30RequestDataInner `json:"data"`
}
