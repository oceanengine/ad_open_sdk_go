/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestShopMultiRoiGoalsInner struct for PromotionUpdateV30RequestShopMultiRoiGoalsInner
type PromotionUpdateV30RequestShopMultiRoiGoalsInner struct {
	//
	RoiGoal      *float64                                         `json:"roi_goal,omitempty"`
	ShopPlatform *PromotionUpdateV30ShopMultiRoiGoalsShopPlatform `json:"shop_platform,omitempty"`
}
