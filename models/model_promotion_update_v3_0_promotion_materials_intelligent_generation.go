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

// PromotionUpdateV30PromotionMaterialsIntelligentGeneration
type PromotionUpdateV30PromotionMaterialsIntelligentGeneration string

// List of promotion_update_v3.0_promotion_materials_intelligent_generation
const (
	OFF_PromotionUpdateV30PromotionMaterialsIntelligentGeneration PromotionUpdateV30PromotionMaterialsIntelligentGeneration = "OFF"
	ON_PromotionUpdateV30PromotionMaterialsIntelligentGeneration  PromotionUpdateV30PromotionMaterialsIntelligentGeneration = "ON"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsIntelligentGeneration enum
var AllowedPromotionUpdateV30PromotionMaterialsIntelligentGenerationEnumValues = []PromotionUpdateV30PromotionMaterialsIntelligentGeneration{
	"OFF",
	"ON",
}

// NewPromotionUpdateV30PromotionMaterialsIntelligentGenerationFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsIntelligentGeneration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsIntelligentGenerationFromValue(v string) (*PromotionUpdateV30PromotionMaterialsIntelligentGeneration, error) {
	ev := PromotionUpdateV30PromotionMaterialsIntelligentGeneration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsIntelligentGeneration: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsIntelligentGenerationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsIntelligentGeneration) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsIntelligentGenerationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_intelligent_generation value
func (v PromotionUpdateV30PromotionMaterialsIntelligentGeneration) Ptr() *PromotionUpdateV30PromotionMaterialsIntelligentGeneration {
	return &v
}
