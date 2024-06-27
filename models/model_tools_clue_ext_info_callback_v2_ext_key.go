/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueExtInfoCallbackV2ExtKey
type ToolsClueExtInfoCallbackV2ExtKey string

// List of tools_clue_ext_info_callback_v2_ext_key
const (
	BUSINESS_STATE_ToolsClueExtInfoCallbackV2ExtKey      ToolsClueExtInfoCallbackV2ExtKey = "BUSINESS_STATE"
	CLUE_INTENTION_ToolsClueExtInfoCallbackV2ExtKey      ToolsClueExtInfoCallbackV2ExtKey = "CLUE_INTENTION"
	CLUE_TAGS_ToolsClueExtInfoCallbackV2ExtKey           ToolsClueExtInfoCallbackV2ExtKey = "CLUE_TAGS"
	CLUE_VALIDITY_ToolsClueExtInfoCallbackV2ExtKey       ToolsClueExtInfoCallbackV2ExtKey = "CLUE_VALIDITY"
	COMMUNICATION_STATE_ToolsClueExtInfoCallbackV2ExtKey ToolsClueExtInfoCallbackV2ExtKey = "COMMUNICATION_STATE"
	CUSTOMER_STATE_ToolsClueExtInfoCallbackV2ExtKey      ToolsClueExtInfoCallbackV2ExtKey = "CUSTOMER_STATE"
	CUSTOMER_TAGS_ToolsClueExtInfoCallbackV2ExtKey       ToolsClueExtInfoCallbackV2ExtKey = "CUSTOMER_TAGS"
	ORDER_STATE_ToolsClueExtInfoCallbackV2ExtKey         ToolsClueExtInfoCallbackV2ExtKey = "ORDER_STATE"
)

// All allowed values of ToolsClueExtInfoCallbackV2ExtKey enum
var AllowedToolsClueExtInfoCallbackV2ExtKeyEnumValues = []ToolsClueExtInfoCallbackV2ExtKey{
	"BUSINESS_STATE",
	"CLUE_INTENTION",
	"CLUE_TAGS",
	"CLUE_VALIDITY",
	"COMMUNICATION_STATE",
	"CUSTOMER_STATE",
	"CUSTOMER_TAGS",
	"ORDER_STATE",
}

// NewToolsClueExtInfoCallbackV2ExtKeyFromValue returns a pointer to a valid ToolsClueExtInfoCallbackV2ExtKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueExtInfoCallbackV2ExtKeyFromValue(v string) (*ToolsClueExtInfoCallbackV2ExtKey, error) {
	ev := ToolsClueExtInfoCallbackV2ExtKey(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueExtInfoCallbackV2ExtKey: valid values are %v", v, AllowedToolsClueExtInfoCallbackV2ExtKeyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueExtInfoCallbackV2ExtKey) IsValid() bool {
	for _, existing := range AllowedToolsClueExtInfoCallbackV2ExtKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_ext_info_callback_v2_ext_key value
func (v ToolsClueExtInfoCallbackV2ExtKey) Ptr() *ToolsClueExtInfoCallbackV2ExtKey {
	return &v
}
