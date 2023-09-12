/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCommentStickOnTopV30StickType
type ToolsCommentStickOnTopV30StickType string

// List of tools_comment_stick_on_top_v3.0_stick_type
const (
	CANCEL_STICK_ToolsCommentStickOnTopV30StickType ToolsCommentStickOnTopV30StickType = "CANCEL_STICK"
	STICK_ON_TOP_ToolsCommentStickOnTopV30StickType ToolsCommentStickOnTopV30StickType = "STICK_ON_TOP"
)

// All allowed values of ToolsCommentStickOnTopV30StickType enum
var AllowedToolsCommentStickOnTopV30StickTypeEnumValues = []ToolsCommentStickOnTopV30StickType{
	"CANCEL_STICK",
	"STICK_ON_TOP",
}

// NewToolsCommentStickOnTopV30StickTypeFromValue returns a pointer to a valid ToolsCommentStickOnTopV30StickType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentStickOnTopV30StickTypeFromValue(v string) (*ToolsCommentStickOnTopV30StickType, error) {
	ev := ToolsCommentStickOnTopV30StickType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentStickOnTopV30StickType: valid values are %v", v, AllowedToolsCommentStickOnTopV30StickTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentStickOnTopV30StickType) IsValid() bool {
	for _, existing := range AllowedToolsCommentStickOnTopV30StickTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_stick_on_top_v3.0_stick_type value
func (v ToolsCommentStickOnTopV30StickType) Ptr() *ToolsCommentStickOnTopV30StickType {
	return &v
}
