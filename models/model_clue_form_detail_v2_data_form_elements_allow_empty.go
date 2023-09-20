/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormDetailV2DataFormElementsAllowEmpty
type ClueFormDetailV2DataFormElementsAllowEmpty string

// List of clue_form_detail_v2_data_form_elements_allow_empty
const (
	Enum_0_ClueFormDetailV2DataFormElementsAllowEmpty ClueFormDetailV2DataFormElementsAllowEmpty = "0"
	Enum_1_ClueFormDetailV2DataFormElementsAllowEmpty ClueFormDetailV2DataFormElementsAllowEmpty = "1"
)

// All allowed values of ClueFormDetailV2DataFormElementsAllowEmpty enum
var AllowedClueFormDetailV2DataFormElementsAllowEmptyEnumValues = []ClueFormDetailV2DataFormElementsAllowEmpty{
	"0",
	"1",
}

// NewClueFormDetailV2DataFormElementsAllowEmptyFromValue returns a pointer to a valid ClueFormDetailV2DataFormElementsAllowEmpty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormElementsAllowEmptyFromValue(v string) (*ClueFormDetailV2DataFormElementsAllowEmpty, error) {
	ev := ClueFormDetailV2DataFormElementsAllowEmpty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormElementsAllowEmpty: valid values are %v", v, AllowedClueFormDetailV2DataFormElementsAllowEmptyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormElementsAllowEmpty) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormElementsAllowEmptyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_elements_allow_empty value
func (v ClueFormDetailV2DataFormElementsAllowEmpty) Ptr() *ClueFormDetailV2DataFormElementsAllowEmpty {
	return &v
}
