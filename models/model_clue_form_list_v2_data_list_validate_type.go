/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormListV2DataListValidateType
type ClueFormListV2DataListValidateType string

// List of clue_form_list_v2_data_list_validate_type
const (
	CLUE_PRIORITY_ClueFormListV2DataListValidateType     ClueFormListV2DataListValidateType = "CLUE_PRIORITY"
	NONE_VERIFICATION_ClueFormListV2DataListValidateType ClueFormListV2DataListValidateType = "NONE_VERIFICATION"
	ALL_VERIFICATION_ClueFormListV2DataListValidateType  ClueFormListV2DataListValidateType = "ALL_VERIFICATION"
	AUTO_VERIFICATION_ClueFormListV2DataListValidateType ClueFormListV2DataListValidateType = "AUTO_VERIFICATION"
	VALIDITY_PRIORITY_ClueFormListV2DataListValidateType ClueFormListV2DataListValidateType = "VALIDITY_PRIORITY"
)

// All allowed values of ClueFormListV2DataListValidateType enum
var AllowedClueFormListV2DataListValidateTypeEnumValues = []ClueFormListV2DataListValidateType{
	"CLUE_PRIORITY",
	"NONE_VERIFICATION",
	"ALL_VERIFICATION",
	"AUTO_VERIFICATION",
	"VALIDITY_PRIORITY",
}

// NewClueFormListV2DataListValidateTypeFromValue returns a pointer to a valid ClueFormListV2DataListValidateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormListV2DataListValidateTypeFromValue(v string) (*ClueFormListV2DataListValidateType, error) {
	ev := ClueFormListV2DataListValidateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormListV2DataListValidateType: valid values are %v", v, AllowedClueFormListV2DataListValidateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormListV2DataListValidateType) IsValid() bool {
	for _, existing := range AllowedClueFormListV2DataListValidateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_list_v2_data_list_validate_type value
func (v ClueFormListV2DataListValidateType) Ptr() *ClueFormListV2DataListValidateType {
	return &v
}