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

// QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization
type QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization int64

// List of qianchuan_ad_detail_get_v1.0_data_creative_list_promotion_card_material_button_smart_optimization
const (
	Enum_0_QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization = 0
	Enum_1_QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization = 1
)

// All allowed values of QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization enum
var AllowedQianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues = []QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimizationFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimizationFromValue(v int64) (*QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization, error) {
	ev := QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_creative_list_promotion_card_material_button_smart_optimization value
func (v QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization) Ptr() *QianchuanAdDetailGetV10DataCreativeListPromotionCardMaterialButtonSmartOptimization {
	return &v
}
