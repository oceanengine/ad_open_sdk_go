/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormListV2IsDel
type ClueFormListV2IsDel string

// List of clue_form_list_v2_is_del
const (
	Enum_0_ClueFormListV2IsDel ClueFormListV2IsDel = "0"
	Enum_1_ClueFormListV2IsDel ClueFormListV2IsDel = "1"
)

// All allowed values of ClueFormListV2IsDel enum
var AllowedClueFormListV2IsDelEnumValues = []ClueFormListV2IsDel{
	"0",
	"1",
}

// NewClueFormListV2IsDelFromValue returns a pointer to a valid ClueFormListV2IsDel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormListV2IsDelFromValue(v string) (*ClueFormListV2IsDel, error) {
	ev := ClueFormListV2IsDel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormListV2IsDel: valid values are %v", v, AllowedClueFormListV2IsDelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormListV2IsDel) IsValid() bool {
	for _, existing := range AllowedClueFormListV2IsDelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_list_v2_is_del value
func (v ClueFormListV2IsDel) Ptr() *ClueFormListV2IsDel {
	return &v
}
