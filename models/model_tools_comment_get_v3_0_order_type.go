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

// ToolsCommentGetV30OrderType
type ToolsCommentGetV30OrderType string

// List of tools_comment_get_v3.0_order_type
const (
	ASC_ToolsCommentGetV30OrderType  ToolsCommentGetV30OrderType = "ASC"
	DESC_ToolsCommentGetV30OrderType ToolsCommentGetV30OrderType = "DESC"
)

// All allowed values of ToolsCommentGetV30OrderType enum
var AllowedToolsCommentGetV30OrderTypeEnumValues = []ToolsCommentGetV30OrderType{
	"ASC",
	"DESC",
}

// NewToolsCommentGetV30OrderTypeFromValue returns a pointer to a valid ToolsCommentGetV30OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30OrderTypeFromValue(v string) (*ToolsCommentGetV30OrderType, error) {
	ev := ToolsCommentGetV30OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30OrderType: valid values are %v", v, AllowedToolsCommentGetV30OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30OrderType) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_order_type value
func (v ToolsCommentGetV30OrderType) Ptr() *ToolsCommentGetV30OrderType {
	return &v
}
