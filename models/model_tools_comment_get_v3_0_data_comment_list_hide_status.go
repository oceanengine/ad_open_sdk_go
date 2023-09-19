/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCommentGetV30DataCommentListHideStatus
type ToolsCommentGetV30DataCommentListHideStatus string

// List of tools_comment_get_v3.0_data_comment_list_hide_status
const (
	ALL_ToolsCommentGetV30DataCommentListHideStatus      ToolsCommentGetV30DataCommentListHideStatus = "ALL"
	HIDE_ToolsCommentGetV30DataCommentListHideStatus     ToolsCommentGetV30DataCommentListHideStatus = "HIDE"
	NOT_HIDE_ToolsCommentGetV30DataCommentListHideStatus ToolsCommentGetV30DataCommentListHideStatus = "NOT_HIDE"
)

// All allowed values of ToolsCommentGetV30DataCommentListHideStatus enum
var AllowedToolsCommentGetV30DataCommentListHideStatusEnumValues = []ToolsCommentGetV30DataCommentListHideStatus{
	"ALL",
	"HIDE",
	"NOT_HIDE",
}

// NewToolsCommentGetV30DataCommentListHideStatusFromValue returns a pointer to a valid ToolsCommentGetV30DataCommentListHideStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30DataCommentListHideStatusFromValue(v string) (*ToolsCommentGetV30DataCommentListHideStatus, error) {
	ev := ToolsCommentGetV30DataCommentListHideStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30DataCommentListHideStatus: valid values are %v", v, AllowedToolsCommentGetV30DataCommentListHideStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30DataCommentListHideStatus) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30DataCommentListHideStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_data_comment_list_hide_status value
func (v ToolsCommentGetV30DataCommentListHideStatus) Ptr() *ToolsCommentGetV30DataCommentListHideStatus {
	return &v
}
