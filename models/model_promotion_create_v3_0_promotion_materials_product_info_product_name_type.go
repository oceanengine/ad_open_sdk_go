/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsProductInfoProductNameType
type PromotionCreateV30PromotionMaterialsProductInfoProductNameType string

// List of promotion_create_v3.0_promotion_materials_product_info_product_name_type
const (
	CUSTOM_PromotionCreateV30PromotionMaterialsProductInfoProductNameType PromotionCreateV30PromotionMaterialsProductInfoProductNameType = "CUSTOM"
	DPA_PromotionCreateV30PromotionMaterialsProductInfoProductNameType    PromotionCreateV30PromotionMaterialsProductInfoProductNameType = "DPA"
)

// All allowed values of PromotionCreateV30PromotionMaterialsProductInfoProductNameType enum
var AllowedPromotionCreateV30PromotionMaterialsProductInfoProductNameTypeEnumValues = []PromotionCreateV30PromotionMaterialsProductInfoProductNameType{
	"CUSTOM",
	"DPA",
}

// NewPromotionCreateV30PromotionMaterialsProductInfoProductNameTypeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsProductInfoProductNameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsProductInfoProductNameTypeFromValue(v string) (*PromotionCreateV30PromotionMaterialsProductInfoProductNameType, error) {
	ev := PromotionCreateV30PromotionMaterialsProductInfoProductNameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsProductInfoProductNameType: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsProductInfoProductNameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsProductInfoProductNameType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsProductInfoProductNameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_product_info_product_name_type value
func (v PromotionCreateV30PromotionMaterialsProductInfoProductNameType) Ptr() *PromotionCreateV30PromotionMaterialsProductInfoProductNameType {
	return &v
}