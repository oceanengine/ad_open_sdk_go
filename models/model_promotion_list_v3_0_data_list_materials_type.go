/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListMaterialsType
type PromotionListV30DataListMaterialsType string

// List of promotion_list_v3.0_data_list_materials_type
const (
	LIVE_MATERIALS_PromotionListV30DataListMaterialsType      PromotionListV30DataListMaterialsType = "LIVE_MATERIALS"
	PROMOTION_MATERIALS_PromotionListV30DataListMaterialsType PromotionListV30DataListMaterialsType = "PROMOTION_MATERIALS"
)

// All allowed values of PromotionListV30DataListMaterialsType enum
var AllowedPromotionListV30DataListMaterialsTypeEnumValues = []PromotionListV30DataListMaterialsType{
	"LIVE_MATERIALS",
	"PROMOTION_MATERIALS",
}

// NewPromotionListV30DataListMaterialsTypeFromValue returns a pointer to a valid PromotionListV30DataListMaterialsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListMaterialsTypeFromValue(v string) (*PromotionListV30DataListMaterialsType, error) {
	ev := PromotionListV30DataListMaterialsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListMaterialsType: valid values are %v", v, AllowedPromotionListV30DataListMaterialsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListMaterialsType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListMaterialsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_materials_type value
func (v PromotionListV30DataListMaterialsType) Ptr() *PromotionListV30DataListMaterialsType {
	return &v
}
