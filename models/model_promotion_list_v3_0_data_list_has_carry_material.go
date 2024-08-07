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

// PromotionListV30DataListHasCarryMaterial
type PromotionListV30DataListHasCarryMaterial string

// List of promotion_list_v3.0_data_list_has_carry_material
const (
	TRUE_PromotionListV30DataListHasCarryMaterial  PromotionListV30DataListHasCarryMaterial = "TRUE"
	FALSE_PromotionListV30DataListHasCarryMaterial PromotionListV30DataListHasCarryMaterial = "FALSE"
)

// All allowed values of PromotionListV30DataListHasCarryMaterial enum
var AllowedPromotionListV30DataListHasCarryMaterialEnumValues = []PromotionListV30DataListHasCarryMaterial{
	"TRUE",
	"FALSE",
}

// NewPromotionListV30DataListHasCarryMaterialFromValue returns a pointer to a valid PromotionListV30DataListHasCarryMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListHasCarryMaterialFromValue(v string) (*PromotionListV30DataListHasCarryMaterial, error) {
	ev := PromotionListV30DataListHasCarryMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListHasCarryMaterial: valid values are %v", v, AllowedPromotionListV30DataListHasCarryMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListHasCarryMaterial) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListHasCarryMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_has_carry_material value
func (v PromotionListV30DataListHasCarryMaterial) Ptr() *PromotionListV30DataListHasCarryMaterial {
	return &v
}
