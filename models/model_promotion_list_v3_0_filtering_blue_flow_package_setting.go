/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30FilteringBlueFlowPackageSetting
type PromotionListV30FilteringBlueFlowPackageSetting string

// List of promotion_list_v3.0_filtering_blue_flow_package_setting
const (
	OFF_PromotionListV30FilteringBlueFlowPackageSetting PromotionListV30FilteringBlueFlowPackageSetting = "OFF"
	ON_PromotionListV30FilteringBlueFlowPackageSetting  PromotionListV30FilteringBlueFlowPackageSetting = "ON"
)

// All allowed values of PromotionListV30FilteringBlueFlowPackageSetting enum
var AllowedPromotionListV30FilteringBlueFlowPackageSettingEnumValues = []PromotionListV30FilteringBlueFlowPackageSetting{
	"OFF",
	"ON",
}

// NewPromotionListV30FilteringBlueFlowPackageSettingFromValue returns a pointer to a valid PromotionListV30FilteringBlueFlowPackageSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringBlueFlowPackageSettingFromValue(v string) (*PromotionListV30FilteringBlueFlowPackageSetting, error) {
	ev := PromotionListV30FilteringBlueFlowPackageSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringBlueFlowPackageSetting: valid values are %v", v, AllowedPromotionListV30FilteringBlueFlowPackageSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringBlueFlowPackageSetting) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringBlueFlowPackageSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_blue_flow_package_setting value
func (v PromotionListV30FilteringBlueFlowPackageSetting) Ptr() *PromotionListV30FilteringBlueFlowPackageSetting {
	return &v
}
