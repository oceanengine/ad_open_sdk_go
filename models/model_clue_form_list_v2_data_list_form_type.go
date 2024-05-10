/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormListV2DataListFormType
type ClueFormListV2DataListFormType string

// List of clue_form_list_v2_data_list_form_type
const (
	ADVANCED_CREATIVE_FORM_ClueFormListV2DataListFormType ClueFormListV2DataListFormType = "ADVANCED_CREATIVE_FORM"
	NATIVE_FORM_ClueFormListV2DataListFormType            ClueFormListV2DataListFormType = "NATIVE_FORM"
	NORMAL_FORM_ClueFormListV2DataListFormType            ClueFormListV2DataListFormType = "NORMAL_FORM"
)

// All allowed values of ClueFormListV2DataListFormType enum
var AllowedClueFormListV2DataListFormTypeEnumValues = []ClueFormListV2DataListFormType{
	"ADVANCED_CREATIVE_FORM",
	"NATIVE_FORM",
	"NORMAL_FORM",
}

// NewClueFormListV2DataListFormTypeFromValue returns a pointer to a valid ClueFormListV2DataListFormType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormListV2DataListFormTypeFromValue(v string) (*ClueFormListV2DataListFormType, error) {
	ev := ClueFormListV2DataListFormType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormListV2DataListFormType: valid values are %v", v, AllowedClueFormListV2DataListFormTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormListV2DataListFormType) IsValid() bool {
	for _, existing := range AllowedClueFormListV2DataListFormTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_list_v2_data_list_form_type value
func (v ClueFormListV2DataListFormType) Ptr() *ClueFormListV2DataListFormType {
	return &v
}
