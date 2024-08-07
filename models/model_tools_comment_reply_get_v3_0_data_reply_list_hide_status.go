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

// ToolsCommentReplyGetV30DataReplyListHideStatus
type ToolsCommentReplyGetV30DataReplyListHideStatus string

// List of tools_comment_reply_get_v3.0_data_reply_list_hide_status
const (
	ALL_ToolsCommentReplyGetV30DataReplyListHideStatus      ToolsCommentReplyGetV30DataReplyListHideStatus = "ALL"
	HIDE_ToolsCommentReplyGetV30DataReplyListHideStatus     ToolsCommentReplyGetV30DataReplyListHideStatus = "HIDE"
	NOT_HIDE_ToolsCommentReplyGetV30DataReplyListHideStatus ToolsCommentReplyGetV30DataReplyListHideStatus = "NOT_HIDE"
)

// All allowed values of ToolsCommentReplyGetV30DataReplyListHideStatus enum
var AllowedToolsCommentReplyGetV30DataReplyListHideStatusEnumValues = []ToolsCommentReplyGetV30DataReplyListHideStatus{
	"ALL",
	"HIDE",
	"NOT_HIDE",
}

// NewToolsCommentReplyGetV30DataReplyListHideStatusFromValue returns a pointer to a valid ToolsCommentReplyGetV30DataReplyListHideStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentReplyGetV30DataReplyListHideStatusFromValue(v string) (*ToolsCommentReplyGetV30DataReplyListHideStatus, error) {
	ev := ToolsCommentReplyGetV30DataReplyListHideStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentReplyGetV30DataReplyListHideStatus: valid values are %v", v, AllowedToolsCommentReplyGetV30DataReplyListHideStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentReplyGetV30DataReplyListHideStatus) IsValid() bool {
	for _, existing := range AllowedToolsCommentReplyGetV30DataReplyListHideStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_reply_get_v3.0_data_reply_list_hide_status value
func (v ToolsCommentReplyGetV30DataReplyListHideStatus) Ptr() *ToolsCommentReplyGetV30DataReplyListHideStatus {
	return &v
}
