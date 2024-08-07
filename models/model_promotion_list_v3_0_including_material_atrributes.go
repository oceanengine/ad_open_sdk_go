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

// PromotionListV30IncludingMaterialAtrributes
type PromotionListV30IncludingMaterialAtrributes string

// List of promotion_list_v3.0_including_material_atrributes
const (
	RETURN_CARRY_DATA_PromotionListV30IncludingMaterialAtrributes PromotionListV30IncludingMaterialAtrributes = "RETURN_CARRY_DATA"
)

// All allowed values of PromotionListV30IncludingMaterialAtrributes enum
var AllowedPromotionListV30IncludingMaterialAtrributesEnumValues = []PromotionListV30IncludingMaterialAtrributes{
	"RETURN_CARRY_DATA",
}

// NewPromotionListV30IncludingMaterialAtrributesFromValue returns a pointer to a valid PromotionListV30IncludingMaterialAtrributes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30IncludingMaterialAtrributesFromValue(v string) (*PromotionListV30IncludingMaterialAtrributes, error) {
	ev := PromotionListV30IncludingMaterialAtrributes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30IncludingMaterialAtrributes: valid values are %v", v, AllowedPromotionListV30IncludingMaterialAtrributesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30IncludingMaterialAtrributes) IsValid() bool {
	for _, existing := range AllowedPromotionListV30IncludingMaterialAtrributesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_including_material_atrributes value
func (v PromotionListV30IncludingMaterialAtrributes) Ptr() *PromotionListV30IncludingMaterialAtrributes {
	return &v
}
