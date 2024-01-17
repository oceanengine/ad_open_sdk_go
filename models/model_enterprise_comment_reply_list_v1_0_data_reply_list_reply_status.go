/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseCommentReplyListV10DataReplyListReplyStatus
type EnterpriseCommentReplyListV10DataReplyListReplyStatus string

// List of enterprise_comment_reply_list_v1.0_data_reply_list_reply_status
const (
	REPLY_AUDIT_SUCCESS_EnterpriseCommentReplyListV10DataReplyListReplyStatus EnterpriseCommentReplyListV10DataReplyListReplyStatus = "REPLY_AUDIT_SUCCESS"
	REPLY_TO_AUDIT_EnterpriseCommentReplyListV10DataReplyListReplyStatus      EnterpriseCommentReplyListV10DataReplyListReplyStatus = "REPLY_TO_AUDIT"
	NO_REPLY_EnterpriseCommentReplyListV10DataReplyListReplyStatus            EnterpriseCommentReplyListV10DataReplyListReplyStatus = "NO_REPLY"
	REPLY_AUDIT_FAIL_EnterpriseCommentReplyListV10DataReplyListReplyStatus    EnterpriseCommentReplyListV10DataReplyListReplyStatus = "REPLY_AUDIT_FAIL"
)

// All allowed values of EnterpriseCommentReplyListV10DataReplyListReplyStatus enum
var AllowedEnterpriseCommentReplyListV10DataReplyListReplyStatusEnumValues = []EnterpriseCommentReplyListV10DataReplyListReplyStatus{
	"REPLY_AUDIT_SUCCESS",
	"REPLY_TO_AUDIT",
	"NO_REPLY",
	"REPLY_AUDIT_FAIL",
}

// NewEnterpriseCommentReplyListV10DataReplyListReplyStatusFromValue returns a pointer to a valid EnterpriseCommentReplyListV10DataReplyListReplyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentReplyListV10DataReplyListReplyStatusFromValue(v string) (*EnterpriseCommentReplyListV10DataReplyListReplyStatus, error) {
	ev := EnterpriseCommentReplyListV10DataReplyListReplyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentReplyListV10DataReplyListReplyStatus: valid values are %v", v, AllowedEnterpriseCommentReplyListV10DataReplyListReplyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentReplyListV10DataReplyListReplyStatus) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentReplyListV10DataReplyListReplyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_reply_list_v1.0_data_reply_list_reply_status value
func (v EnterpriseCommentReplyListV10DataReplyListReplyStatus) Ptr() *EnterpriseCommentReplyListV10DataReplyListReplyStatus {
	return &v
}
