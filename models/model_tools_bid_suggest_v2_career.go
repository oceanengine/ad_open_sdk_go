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

// ToolsBidSuggestV2Career
type ToolsBidSuggestV2Career string

// List of tools_bid_suggest_v2_career
const (
	CIVIL_SERVANTS_ToolsBidSuggestV2Career  ToolsBidSuggestV2Career = "CIVIL_SERVANTS"
	IT_ToolsBidSuggestV2Career              ToolsBidSuggestV2Career = "IT"
	MEDICAL_STAFF_ToolsBidSuggestV2Career   ToolsBidSuggestV2Career = "MEDICAL_STAFF"
	COLLEGE_STUDENT_ToolsBidSuggestV2Career ToolsBidSuggestV2Career = "COLLEGE_STUDENT"
	FINANCIAL_ToolsBidSuggestV2Career       ToolsBidSuggestV2Career = "FINANCIAL"
	TEACHER_ToolsBidSuggestV2Career         ToolsBidSuggestV2Career = "TEACHER"
)

// All allowed values of ToolsBidSuggestV2Career enum
var AllowedToolsBidSuggestV2CareerEnumValues = []ToolsBidSuggestV2Career{
	"CIVIL_SERVANTS",
	"IT",
	"MEDICAL_STAFF",
	"COLLEGE_STUDENT",
	"FINANCIAL",
	"TEACHER",
}

// NewToolsBidSuggestV2CareerFromValue returns a pointer to a valid ToolsBidSuggestV2Career
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2CareerFromValue(v string) (*ToolsBidSuggestV2Career, error) {
	ev := ToolsBidSuggestV2Career(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2Career: valid values are %v", v, AllowedToolsBidSuggestV2CareerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2Career) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2CareerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_career value
func (v ToolsBidSuggestV2Career) Ptr() *ToolsBidSuggestV2Career {
	return &v
}
