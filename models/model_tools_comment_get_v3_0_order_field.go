/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCommentGetV30OrderField
type ToolsCommentGetV30OrderField string

// List of tools_comment_get_v3.0_order_field
const (
	CREATE_TIME_ToolsCommentGetV30OrderField ToolsCommentGetV30OrderField = "CREATE_TIME"
	LIKE_COUNT_ToolsCommentGetV30OrderField  ToolsCommentGetV30OrderField = "LIKE_COUNT"
	REPLY_COUNT_ToolsCommentGetV30OrderField ToolsCommentGetV30OrderField = "REPLY_COUNT"
)

// All allowed values of ToolsCommentGetV30OrderField enum
var AllowedToolsCommentGetV30OrderFieldEnumValues = []ToolsCommentGetV30OrderField{
	"CREATE_TIME",
	"LIKE_COUNT",
	"REPLY_COUNT",
}

// NewToolsCommentGetV30OrderFieldFromValue returns a pointer to a valid ToolsCommentGetV30OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30OrderFieldFromValue(v string) (*ToolsCommentGetV30OrderField, error) {
	ev := ToolsCommentGetV30OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30OrderField: valid values are %v", v, AllowedToolsCommentGetV30OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30OrderField) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_order_field value
func (v ToolsCommentGetV30OrderField) Ptr() *ToolsCommentGetV30OrderField {
	return &v
}
