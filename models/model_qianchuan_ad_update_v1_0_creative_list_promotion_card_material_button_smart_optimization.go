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

// QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization
type QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization int64

// List of qianchuan_ad_update_v1.0_creative_list_promotion_card_material_button_smart_optimization
const (
	Enum_0_QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization = 0
	Enum_1_QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization = 1
)

// All allowed values of QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization enum
var AllowedQianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues = []QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization{
	0,
	1,
}

// NewQianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimizationFromValue returns a pointer to a valid QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimizationFromValue(v int64) (*QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization, error) {
	ev := QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization: valid values are %v", v, AllowedQianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_creative_list_promotion_card_material_button_smart_optimization value
func (v QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization) Ptr() *QianchuanAdUpdateV10CreativeListPromotionCardMaterialButtonSmartOptimization {
	return &v
}
