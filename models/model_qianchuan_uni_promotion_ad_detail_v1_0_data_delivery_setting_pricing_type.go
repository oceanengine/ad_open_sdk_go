/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType
type QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType string

// List of qianchuan_uni_promotion_ad_detail_v1.0_data_delivery_setting_pricing_type
const (
	OCPM_QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType = "OCPM"
)

// All allowed values of QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType enum
var AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingPricingTypeEnumValues = []QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType{
	"OCPM",
}

// NewQianchuanUniPromotionAdDetailV10DataDeliverySettingPricingTypeFromValue returns a pointer to a valid QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdDetailV10DataDeliverySettingPricingTypeFromValue(v string) (*QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType, error) {
	ev := QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType: valid values are %v", v, AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingPricingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdDetailV10DataDeliverySettingPricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_detail_v1.0_data_delivery_setting_pricing_type value
func (v QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType) Ptr() *QianchuanUniPromotionAdDetailV10DataDeliverySettingPricingType {
	return &v
}