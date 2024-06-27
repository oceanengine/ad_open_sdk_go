/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30ShopMultiRoiGoalsShopPlatform
type PromotionCreateV30ShopMultiRoiGoalsShopPlatform string

// List of promotion_create_v3.0_shop_multi_roi_goals_shop_platform
const (
	JD_PromotionCreateV30ShopMultiRoiGoalsShopPlatform    PromotionCreateV30ShopMultiRoiGoalsShopPlatform = "JD"
	OTHER_PromotionCreateV30ShopMultiRoiGoalsShopPlatform PromotionCreateV30ShopMultiRoiGoalsShopPlatform = "OTHER"
	PDD_PromotionCreateV30ShopMultiRoiGoalsShopPlatform   PromotionCreateV30ShopMultiRoiGoalsShopPlatform = "PDD"
	TB_PromotionCreateV30ShopMultiRoiGoalsShopPlatform    PromotionCreateV30ShopMultiRoiGoalsShopPlatform = "TB"
)

// All allowed values of PromotionCreateV30ShopMultiRoiGoalsShopPlatform enum
var AllowedPromotionCreateV30ShopMultiRoiGoalsShopPlatformEnumValues = []PromotionCreateV30ShopMultiRoiGoalsShopPlatform{
	"JD",
	"OTHER",
	"PDD",
	"TB",
}

// NewPromotionCreateV30ShopMultiRoiGoalsShopPlatformFromValue returns a pointer to a valid PromotionCreateV30ShopMultiRoiGoalsShopPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30ShopMultiRoiGoalsShopPlatformFromValue(v string) (*PromotionCreateV30ShopMultiRoiGoalsShopPlatform, error) {
	ev := PromotionCreateV30ShopMultiRoiGoalsShopPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30ShopMultiRoiGoalsShopPlatform: valid values are %v", v, AllowedPromotionCreateV30ShopMultiRoiGoalsShopPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30ShopMultiRoiGoalsShopPlatform) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30ShopMultiRoiGoalsShopPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_shop_multi_roi_goals_shop_platform value
func (v PromotionCreateV30ShopMultiRoiGoalsShopPlatform) Ptr() *PromotionCreateV30ShopMultiRoiGoalsShopPlatform {
	return &v
}
