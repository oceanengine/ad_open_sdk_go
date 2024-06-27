/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType
type PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType string

// List of promotion_update_v3.0_promotion_materials_product_info_product_selling_point_type
const (
	CUSTOM_PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType = "CUSTOM"
	DPA_PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType    PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType = "DPA"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType enum
var AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointTypeEnumValues = []PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType{
	"CUSTOM",
	"DPA",
}

// NewPromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointTypeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointTypeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType, error) {
	ev := PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_product_info_product_selling_point_type value
func (v PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType) Ptr() *PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType {
	return &v
}
