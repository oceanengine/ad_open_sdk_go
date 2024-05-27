/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial
type PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial string

// List of promotion_list_v3.0_data_list_promotion_materials_video_material_list_is_carry_material
const (
	TRUE_PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial  PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial = "TRUE"
	FALSE_PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial = "FALSE"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial enum
var AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterialEnumValues = []PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial{
	"TRUE",
	"FALSE",
}

// NewPromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterialFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterialFromValue(v string) (*PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial, error) {
	ev := PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_video_material_list_is_carry_material value
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial) Ptr() *PromotionListV30DataListPromotionMaterialsVideoMaterialListIsCarryMaterial {
	return &v
}
