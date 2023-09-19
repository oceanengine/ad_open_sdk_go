/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial
type PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial string

// List of promotion_list_v3.0_data_list_material_score_info_score_num_of_material
const (
	LEVEL1_PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial = "LEVEL1"
	LEVEL2_PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial = "LEVEL2"
	LEVEL3_PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial = "LEVEL3"
	LEVEL4_PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial = "LEVEL4"
	LEVEL5_PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial = "LEVEL5"
)

// All allowed values of PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial enum
var AllowedPromotionListV30DataListMaterialScoreInfoScoreNumOfMaterialEnumValues = []PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial{
	"LEVEL1",
	"LEVEL2",
	"LEVEL3",
	"LEVEL4",
	"LEVEL5",
}

// NewPromotionListV30DataListMaterialScoreInfoScoreNumOfMaterialFromValue returns a pointer to a valid PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListMaterialScoreInfoScoreNumOfMaterialFromValue(v string) (*PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial, error) {
	ev := PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial: valid values are %v", v, AllowedPromotionListV30DataListMaterialScoreInfoScoreNumOfMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListMaterialScoreInfoScoreNumOfMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_material_score_info_score_num_of_material value
func (v PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial) Ptr() *PromotionListV30DataListMaterialScoreInfoScoreNumOfMaterial {
	return &v
}
