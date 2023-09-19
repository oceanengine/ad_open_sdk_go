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

// ToolsCommentGetV30DataCommentListReplyStatus
type ToolsCommentGetV30DataCommentListReplyStatus string

// List of tools_comment_get_v3.0_data_comment_list_reply_status
const (
	NO_REPLY_ToolsCommentGetV30DataCommentListReplyStatus            ToolsCommentGetV30DataCommentListReplyStatus = "NO_REPLY"
	REPLY_AUDIT_FAIL_ToolsCommentGetV30DataCommentListReplyStatus    ToolsCommentGetV30DataCommentListReplyStatus = "REPLY_AUDIT_FAIL"
	REPLY_AUDIT_SUCCESS_ToolsCommentGetV30DataCommentListReplyStatus ToolsCommentGetV30DataCommentListReplyStatus = "REPLY_AUDIT_SUCCESS"
	REPLY_TO_AUDIT_ToolsCommentGetV30DataCommentListReplyStatus      ToolsCommentGetV30DataCommentListReplyStatus = "REPLY_TO_AUDIT"
)

// All allowed values of ToolsCommentGetV30DataCommentListReplyStatus enum
var AllowedToolsCommentGetV30DataCommentListReplyStatusEnumValues = []ToolsCommentGetV30DataCommentListReplyStatus{
	"NO_REPLY",
	"REPLY_AUDIT_FAIL",
	"REPLY_AUDIT_SUCCESS",
	"REPLY_TO_AUDIT",
}

// NewToolsCommentGetV30DataCommentListReplyStatusFromValue returns a pointer to a valid ToolsCommentGetV30DataCommentListReplyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30DataCommentListReplyStatusFromValue(v string) (*ToolsCommentGetV30DataCommentListReplyStatus, error) {
	ev := ToolsCommentGetV30DataCommentListReplyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30DataCommentListReplyStatus: valid values are %v", v, AllowedToolsCommentGetV30DataCommentListReplyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30DataCommentListReplyStatus) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30DataCommentListReplyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_data_comment_list_reply_status value
func (v ToolsCommentGetV30DataCommentListReplyStatus) Ptr() *ToolsCommentGetV30DataCommentListReplyStatus {
	return &v
}
