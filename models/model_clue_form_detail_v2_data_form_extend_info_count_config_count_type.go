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

// ClueFormDetailV2DataFormExtendInfoCountConfigCountType
type ClueFormDetailV2DataFormExtendInfoCountConfigCountType string

// List of clue_form_detail_v2_data_form_extend_info_count_config_count_type
const (
	COUNT_TYPE_INCREMENT_ClueFormDetailV2DataFormExtendInfoCountConfigCountType ClueFormDetailV2DataFormExtendInfoCountConfigCountType = "COUNT_TYPE_INCREMENT"
	COUNT_TYPE_DECREMENT_ClueFormDetailV2DataFormExtendInfoCountConfigCountType ClueFormDetailV2DataFormExtendInfoCountConfigCountType = "COUNT_TYPE_DECREMENT"
)

// All allowed values of ClueFormDetailV2DataFormExtendInfoCountConfigCountType enum
var AllowedClueFormDetailV2DataFormExtendInfoCountConfigCountTypeEnumValues = []ClueFormDetailV2DataFormExtendInfoCountConfigCountType{
	"COUNT_TYPE_INCREMENT",
	"COUNT_TYPE_DECREMENT",
}

// NewClueFormDetailV2DataFormExtendInfoCountConfigCountTypeFromValue returns a pointer to a valid ClueFormDetailV2DataFormExtendInfoCountConfigCountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormExtendInfoCountConfigCountTypeFromValue(v string) (*ClueFormDetailV2DataFormExtendInfoCountConfigCountType, error) {
	ev := ClueFormDetailV2DataFormExtendInfoCountConfigCountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormExtendInfoCountConfigCountType: valid values are %v", v, AllowedClueFormDetailV2DataFormExtendInfoCountConfigCountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormExtendInfoCountConfigCountType) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormExtendInfoCountConfigCountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_extend_info_count_config_count_type value
func (v ClueFormDetailV2DataFormExtendInfoCountConfigCountType) Ptr() *ClueFormDetailV2DataFormExtendInfoCountConfigCountType {
	return &v
}