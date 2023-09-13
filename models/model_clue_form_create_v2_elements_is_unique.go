/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormCreateV2ElementsIsUnique
type ClueFormCreateV2ElementsIsUnique string

// List of clue_form_create_v2_elements_is_unique
const (
	Enum_0_ClueFormCreateV2ElementsIsUnique ClueFormCreateV2ElementsIsUnique = "0"
	Enum_1_ClueFormCreateV2ElementsIsUnique ClueFormCreateV2ElementsIsUnique = "1"
)

// All allowed values of ClueFormCreateV2ElementsIsUnique enum
var AllowedClueFormCreateV2ElementsIsUniqueEnumValues = []ClueFormCreateV2ElementsIsUnique{
	"0",
	"1",
}

// NewClueFormCreateV2ElementsIsUniqueFromValue returns a pointer to a valid ClueFormCreateV2ElementsIsUnique
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2ElementsIsUniqueFromValue(v string) (*ClueFormCreateV2ElementsIsUnique, error) {
	ev := ClueFormCreateV2ElementsIsUnique(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2ElementsIsUnique: valid values are %v", v, AllowedClueFormCreateV2ElementsIsUniqueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2ElementsIsUnique) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2ElementsIsUniqueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_elements_is_unique value
func (v ClueFormCreateV2ElementsIsUnique) Ptr() *ClueFormCreateV2ElementsIsUnique {
	return &v
}
