/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30MaterialsType
type PromotionCreateV30MaterialsType string

// List of promotion_create_v3.0_materials_type
const (
	LIVE_MATERIALS_PromotionCreateV30MaterialsType      PromotionCreateV30MaterialsType = "LIVE_MATERIALS"
	PROMOTION_MATERIALS_PromotionCreateV30MaterialsType PromotionCreateV30MaterialsType = "PROMOTION_MATERIALS"
)

// All allowed values of PromotionCreateV30MaterialsType enum
var AllowedPromotionCreateV30MaterialsTypeEnumValues = []PromotionCreateV30MaterialsType{
	"LIVE_MATERIALS",
	"PROMOTION_MATERIALS",
}

// NewPromotionCreateV30MaterialsTypeFromValue returns a pointer to a valid PromotionCreateV30MaterialsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30MaterialsTypeFromValue(v string) (*PromotionCreateV30MaterialsType, error) {
	ev := PromotionCreateV30MaterialsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30MaterialsType: valid values are %v", v, AllowedPromotionCreateV30MaterialsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30MaterialsType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30MaterialsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_materials_type value
func (v PromotionCreateV30MaterialsType) Ptr() *PromotionCreateV30MaterialsType {
	return &v
}
