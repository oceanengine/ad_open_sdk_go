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

// PromotionListV30DataListPromotionMaterialsProductInfoProductImageType
type PromotionListV30DataListPromotionMaterialsProductInfoProductImageType string

// List of promotion_list_v3.0_data_list_promotion_materials_product_info_product_image_type
const (
	CUSTOM_PromotionListV30DataListPromotionMaterialsProductInfoProductImageType PromotionListV30DataListPromotionMaterialsProductInfoProductImageType = "CUSTOM"
	DPA_PromotionListV30DataListPromotionMaterialsProductInfoProductImageType    PromotionListV30DataListPromotionMaterialsProductInfoProductImageType = "DPA"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsProductInfoProductImageType enum
var AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductImageTypeEnumValues = []PromotionListV30DataListPromotionMaterialsProductInfoProductImageType{
	"CUSTOM",
	"DPA",
}

// NewPromotionListV30DataListPromotionMaterialsProductInfoProductImageTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsProductInfoProductImageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsProductInfoProductImageTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsProductInfoProductImageType, error) {
	ev := PromotionListV30DataListPromotionMaterialsProductInfoProductImageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsProductInfoProductImageType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductImageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsProductInfoProductImageType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductImageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_product_info_product_image_type value
func (v PromotionListV30DataListPromotionMaterialsProductInfoProductImageType) Ptr() *PromotionListV30DataListPromotionMaterialsProductInfoProductImageType {
	return &v
}
