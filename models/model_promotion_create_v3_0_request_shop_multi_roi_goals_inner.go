/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestShopMultiRoiGoalsInner struct for PromotionCreateV30RequestShopMultiRoiGoalsInner
type PromotionCreateV30RequestShopMultiRoiGoalsInner struct {
	//
	RoiGoal      *float64                                         `json:"roi_goal,omitempty"`
	ShopPlatform *PromotionCreateV30ShopMultiRoiGoalsShopPlatform `json:"shop_platform,omitempty"`
}
