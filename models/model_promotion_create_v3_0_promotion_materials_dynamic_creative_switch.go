/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch
type PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch string

// List of promotion_create_v3.0_promotion_materials_dynamic_creative_switch
const (
	OFF_PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch = "OFF"
	ON_PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch  PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch = "ON"
)

// All allowed values of PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch enum
var AllowedPromotionCreateV30PromotionMaterialsDynamicCreativeSwitchEnumValues = []PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch{
	"OFF",
	"ON",
}

// NewPromotionCreateV30PromotionMaterialsDynamicCreativeSwitchFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsDynamicCreativeSwitchFromValue(v string) (*PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch, error) {
	ev := PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_dynamic_creative_switch value
func (v PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch) Ptr() *PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch {
	return &v
}
