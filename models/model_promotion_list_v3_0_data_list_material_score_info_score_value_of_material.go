/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial
type PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial string

// List of promotion_list_v3.0_data_list_material_score_info_score_value_of_material
const (
	LEVEL1_PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial = "LEVEL1"
	LEVEL2_PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial = "LEVEL2"
	LEVEL3_PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial = "LEVEL3"
	LEVEL4_PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial = "LEVEL4"
	LEVEL5_PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial = "LEVEL5"
)

// All allowed values of PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial enum
var AllowedPromotionListV30DataListMaterialScoreInfoScoreValueOfMaterialEnumValues = []PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial{
	"LEVEL1",
	"LEVEL2",
	"LEVEL3",
	"LEVEL4",
	"LEVEL5",
}

// NewPromotionListV30DataListMaterialScoreInfoScoreValueOfMaterialFromValue returns a pointer to a valid PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListMaterialScoreInfoScoreValueOfMaterialFromValue(v string) (*PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial, error) {
	ev := PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial: valid values are %v", v, AllowedPromotionListV30DataListMaterialScoreInfoScoreValueOfMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListMaterialScoreInfoScoreValueOfMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_material_score_info_score_value_of_material value
func (v PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial) Ptr() *PromotionListV30DataListMaterialScoreInfoScoreValueOfMaterial {
	return &v
}
