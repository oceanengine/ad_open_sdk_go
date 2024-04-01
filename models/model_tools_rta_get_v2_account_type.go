/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRtaGetV2AccountType
type ToolsRtaGetV2AccountType string

// List of tools_rta_get_v2_account_type
const (
	AD_ToolsRtaGetV2AccountType   ToolsRtaGetV2AccountType = "AD"
	STAR_ToolsRtaGetV2AccountType ToolsRtaGetV2AccountType = "STAR"
)

// All allowed values of ToolsRtaGetV2AccountType enum
var AllowedToolsRtaGetV2AccountTypeEnumValues = []ToolsRtaGetV2AccountType{
	"AD",
	"STAR",
}

// NewToolsRtaGetV2AccountTypeFromValue returns a pointer to a valid ToolsRtaGetV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRtaGetV2AccountTypeFromValue(v string) (*ToolsRtaGetV2AccountType, error) {
	ev := ToolsRtaGetV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRtaGetV2AccountType: valid values are %v", v, AllowedToolsRtaGetV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRtaGetV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsRtaGetV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rta_get_v2_account_type value
func (v ToolsRtaGetV2AccountType) Ptr() *ToolsRtaGetV2AccountType {
	return &v
}
