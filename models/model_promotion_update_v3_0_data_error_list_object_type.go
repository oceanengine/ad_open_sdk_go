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

// PromotionUpdateV30DataErrorListObjectType
type PromotionUpdateV30DataErrorListObjectType string

// List of promotion_update_v3.0_data_error_list_object_type
const (
	BUDGET_PromotionUpdateV30DataErrorListObjectType             PromotionUpdateV30DataErrorListObjectType = "BUDGET"
	KEYWORDS_PromotionUpdateV30DataErrorListObjectType           PromotionUpdateV30DataErrorListObjectType = "KEYWORDS"
	PROJECT_SETTING_PromotionUpdateV30DataErrorListObjectType    PromotionUpdateV30DataErrorListObjectType = "PROJECT_SETTING"
	PROMOTION_BUDGET_PromotionUpdateV30DataErrorListObjectType   PromotionUpdateV30DataErrorListObjectType = "PROMOTION_BUDGET"
	PROMOTION_MATERIAL_PromotionUpdateV30DataErrorListObjectType PromotionUpdateV30DataErrorListObjectType = "PROMOTION_MATERIAL"
	PROMOTION_SETTING_PromotionUpdateV30DataErrorListObjectType  PromotionUpdateV30DataErrorListObjectType = "PROMOTION_SETTING"
)

// All allowed values of PromotionUpdateV30DataErrorListObjectType enum
var AllowedPromotionUpdateV30DataErrorListObjectTypeEnumValues = []PromotionUpdateV30DataErrorListObjectType{
	"BUDGET",
	"KEYWORDS",
	"PROJECT_SETTING",
	"PROMOTION_BUDGET",
	"PROMOTION_MATERIAL",
	"PROMOTION_SETTING",
}

// NewPromotionUpdateV30DataErrorListObjectTypeFromValue returns a pointer to a valid PromotionUpdateV30DataErrorListObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30DataErrorListObjectTypeFromValue(v string) (*PromotionUpdateV30DataErrorListObjectType, error) {
	ev := PromotionUpdateV30DataErrorListObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30DataErrorListObjectType: valid values are %v", v, AllowedPromotionUpdateV30DataErrorListObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30DataErrorListObjectType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30DataErrorListObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_data_error_list_object_type value
func (v PromotionUpdateV30DataErrorListObjectType) Ptr() *PromotionUpdateV30DataErrorListObjectType {
	return &v
}
