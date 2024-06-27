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

// ClueFormCreateV2ValidateType
type ClueFormCreateV2ValidateType string

// List of clue_form_create_v2_validate_type
const (
	NONE_VERIFICATION_ClueFormCreateV2ValidateType ClueFormCreateV2ValidateType = "NONE_VERIFICATION"
	AUTO_VERIFICATION_ClueFormCreateV2ValidateType ClueFormCreateV2ValidateType = "AUTO_VERIFICATION"
	CLUE_PRIORITY_ClueFormCreateV2ValidateType     ClueFormCreateV2ValidateType = "CLUE_PRIORITY"
	VALIDITY_PRIORITY_ClueFormCreateV2ValidateType ClueFormCreateV2ValidateType = "VALIDITY_PRIORITY"
	ALL_VERIFICATION_ClueFormCreateV2ValidateType  ClueFormCreateV2ValidateType = "ALL_VERIFICATION"
)

// All allowed values of ClueFormCreateV2ValidateType enum
var AllowedClueFormCreateV2ValidateTypeEnumValues = []ClueFormCreateV2ValidateType{
	"NONE_VERIFICATION",
	"AUTO_VERIFICATION",
	"CLUE_PRIORITY",
	"VALIDITY_PRIORITY",
	"ALL_VERIFICATION",
}

// NewClueFormCreateV2ValidateTypeFromValue returns a pointer to a valid ClueFormCreateV2ValidateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2ValidateTypeFromValue(v string) (*ClueFormCreateV2ValidateType, error) {
	ev := ClueFormCreateV2ValidateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2ValidateType: valid values are %v", v, AllowedClueFormCreateV2ValidateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2ValidateType) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2ValidateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_validate_type value
func (v ClueFormCreateV2ValidateType) Ptr() *ClueFormCreateV2ValidateType {
	return &v
}
