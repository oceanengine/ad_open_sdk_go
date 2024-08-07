/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch
type PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch string

// List of promotion_list_v3.0_data_list_promotion_materials_dynamic_creative_switch
const (
	OFF_PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch = "OFF"
	ON_PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch  PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch = "ON"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch enum
var AllowedPromotionListV30DataListPromotionMaterialsDynamicCreativeSwitchEnumValues = []PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListPromotionMaterialsDynamicCreativeSwitchFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsDynamicCreativeSwitchFromValue(v string) (*PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch, error) {
	ev := PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_dynamic_creative_switch value
func (v PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch) Ptr() *PromotionListV30DataListPromotionMaterialsDynamicCreativeSwitch {
	return &v
}
