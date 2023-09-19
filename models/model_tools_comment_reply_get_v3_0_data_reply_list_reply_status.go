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

// ToolsCommentReplyGetV30DataReplyListReplyStatus
type ToolsCommentReplyGetV30DataReplyListReplyStatus string

// List of tools_comment_reply_get_v3.0_data_reply_list_reply_status
const (
	NO_REPLY_ToolsCommentReplyGetV30DataReplyListReplyStatus            ToolsCommentReplyGetV30DataReplyListReplyStatus = "NO_REPLY"
	REPLY_AUDIT_FAIL_ToolsCommentReplyGetV30DataReplyListReplyStatus    ToolsCommentReplyGetV30DataReplyListReplyStatus = "REPLY_AUDIT_FAIL"
	REPLY_AUDIT_SUCCESS_ToolsCommentReplyGetV30DataReplyListReplyStatus ToolsCommentReplyGetV30DataReplyListReplyStatus = "REPLY_AUDIT_SUCCESS"
	REPLY_TO_AUDIT_ToolsCommentReplyGetV30DataReplyListReplyStatus      ToolsCommentReplyGetV30DataReplyListReplyStatus = "REPLY_TO_AUDIT"
)

// All allowed values of ToolsCommentReplyGetV30DataReplyListReplyStatus enum
var AllowedToolsCommentReplyGetV30DataReplyListReplyStatusEnumValues = []ToolsCommentReplyGetV30DataReplyListReplyStatus{
	"NO_REPLY",
	"REPLY_AUDIT_FAIL",
	"REPLY_AUDIT_SUCCESS",
	"REPLY_TO_AUDIT",
}

// NewToolsCommentReplyGetV30DataReplyListReplyStatusFromValue returns a pointer to a valid ToolsCommentReplyGetV30DataReplyListReplyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentReplyGetV30DataReplyListReplyStatusFromValue(v string) (*ToolsCommentReplyGetV30DataReplyListReplyStatus, error) {
	ev := ToolsCommentReplyGetV30DataReplyListReplyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentReplyGetV30DataReplyListReplyStatus: valid values are %v", v, AllowedToolsCommentReplyGetV30DataReplyListReplyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentReplyGetV30DataReplyListReplyStatus) IsValid() bool {
	for _, existing := range AllowedToolsCommentReplyGetV30DataReplyListReplyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_reply_get_v3.0_data_reply_list_reply_status value
func (v ToolsCommentReplyGetV30DataReplyListReplyStatus) Ptr() *ToolsCommentReplyGetV30DataReplyListReplyStatus {
	return &v
}
