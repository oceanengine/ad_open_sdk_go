/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormDetailV2DataFormIsDel
type ClueFormDetailV2DataFormIsDel string

// List of clue_form_detail_v2_data_form_is_del
const (
	Enum_0_ClueFormDetailV2DataFormIsDel ClueFormDetailV2DataFormIsDel = "0"
	Enum_1_ClueFormDetailV2DataFormIsDel ClueFormDetailV2DataFormIsDel = "1"
)

// All allowed values of ClueFormDetailV2DataFormIsDel enum
var AllowedClueFormDetailV2DataFormIsDelEnumValues = []ClueFormDetailV2DataFormIsDel{
	"0",
	"1",
}

// NewClueFormDetailV2DataFormIsDelFromValue returns a pointer to a valid ClueFormDetailV2DataFormIsDel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormIsDelFromValue(v string) (*ClueFormDetailV2DataFormIsDel, error) {
	ev := ClueFormDetailV2DataFormIsDel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormIsDel: valid values are %v", v, AllowedClueFormDetailV2DataFormIsDelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormIsDel) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormIsDelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_is_del value
func (v ClueFormDetailV2DataFormIsDel) Ptr() *ClueFormDetailV2DataFormIsDel {
	return &v
}
