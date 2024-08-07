/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormDetailV2DataFormElementsIsUnique
type ClueFormDetailV2DataFormElementsIsUnique string

// List of clue_form_detail_v2_data_form_elements_is_unique
const (
	Enum_0_ClueFormDetailV2DataFormElementsIsUnique ClueFormDetailV2DataFormElementsIsUnique = "0"
	Enum_1_ClueFormDetailV2DataFormElementsIsUnique ClueFormDetailV2DataFormElementsIsUnique = "1"
)

// All allowed values of ClueFormDetailV2DataFormElementsIsUnique enum
var AllowedClueFormDetailV2DataFormElementsIsUniqueEnumValues = []ClueFormDetailV2DataFormElementsIsUnique{
	"0",
	"1",
}

// NewClueFormDetailV2DataFormElementsIsUniqueFromValue returns a pointer to a valid ClueFormDetailV2DataFormElementsIsUnique
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormElementsIsUniqueFromValue(v string) (*ClueFormDetailV2DataFormElementsIsUnique, error) {
	ev := ClueFormDetailV2DataFormElementsIsUnique(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormElementsIsUnique: valid values are %v", v, AllowedClueFormDetailV2DataFormElementsIsUniqueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormElementsIsUnique) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormElementsIsUniqueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_elements_is_unique value
func (v ClueFormDetailV2DataFormElementsIsUnique) Ptr() *ClueFormDetailV2DataFormElementsIsUnique {
	return &v
}
