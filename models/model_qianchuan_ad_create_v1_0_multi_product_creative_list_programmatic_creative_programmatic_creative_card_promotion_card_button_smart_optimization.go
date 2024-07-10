/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization
type QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization int64

// List of qianchuan_ad_create_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_card_promotion_card_button_smart_optimization
const (
	Enum_0_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization = 0
	Enum_1_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization = 1
)

// All allowed values of QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization enum
var AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues = []QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization{
	0,
	1,
}

// NewQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimizationFromValue returns a pointer to a valid QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimizationFromValue(v int64) (*QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization, error) {
	ev := QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization: valid values are %v", v, AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_card_promotion_card_button_smart_optimization value
func (v QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization) Ptr() *QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization {
	return &v
}
