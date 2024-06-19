/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormDetailV2DataFormValidateType
type ClueFormDetailV2DataFormValidateType string

// List of clue_form_detail_v2_data_form_validate_type
const (
	VALIDITY_PRIORITY_ClueFormDetailV2DataFormValidateType ClueFormDetailV2DataFormValidateType = "VALIDITY_PRIORITY"
	NONE_VERIFICATION_ClueFormDetailV2DataFormValidateType ClueFormDetailV2DataFormValidateType = "NONE_VERIFICATION"
	ALL_VERIFICATION_ClueFormDetailV2DataFormValidateType  ClueFormDetailV2DataFormValidateType = "ALL_VERIFICATION"
	CLUE_PRIORITY_ClueFormDetailV2DataFormValidateType     ClueFormDetailV2DataFormValidateType = "CLUE_PRIORITY"
	AUTO_VERIFICATION_ClueFormDetailV2DataFormValidateType ClueFormDetailV2DataFormValidateType = "AUTO_VERIFICATION"
)

// All allowed values of ClueFormDetailV2DataFormValidateType enum
var AllowedClueFormDetailV2DataFormValidateTypeEnumValues = []ClueFormDetailV2DataFormValidateType{
	"VALIDITY_PRIORITY",
	"NONE_VERIFICATION",
	"ALL_VERIFICATION",
	"CLUE_PRIORITY",
	"AUTO_VERIFICATION",
}

// NewClueFormDetailV2DataFormValidateTypeFromValue returns a pointer to a valid ClueFormDetailV2DataFormValidateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormValidateTypeFromValue(v string) (*ClueFormDetailV2DataFormValidateType, error) {
	ev := ClueFormDetailV2DataFormValidateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormValidateType: valid values are %v", v, AllowedClueFormDetailV2DataFormValidateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormValidateType) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormValidateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_validate_type value
func (v ClueFormDetailV2DataFormValidateType) Ptr() *ClueFormDetailV2DataFormValidateType {
	return &v
}
