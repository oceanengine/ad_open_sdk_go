/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseCommentListGetV10FilterReplyStatusByEDouyin
type EnterpriseCommentListGetV10FilterReplyStatusByEDouyin string

// List of enterprise_comment_list_get_v1.0_filter_reply_status_by_e_douyin
const (
	ALL_EnterpriseCommentListGetV10FilterReplyStatusByEDouyin      EnterpriseCommentListGetV10FilterReplyStatusByEDouyin = "ALL"
	REPLY_EnterpriseCommentListGetV10FilterReplyStatusByEDouyin    EnterpriseCommentListGetV10FilterReplyStatusByEDouyin = "REPLY"
	NO_REPLY_EnterpriseCommentListGetV10FilterReplyStatusByEDouyin EnterpriseCommentListGetV10FilterReplyStatusByEDouyin = "NO_REPLY"
)

// All allowed values of EnterpriseCommentListGetV10FilterReplyStatusByEDouyin enum
var AllowedEnterpriseCommentListGetV10FilterReplyStatusByEDouyinEnumValues = []EnterpriseCommentListGetV10FilterReplyStatusByEDouyin{
	"ALL",
	"REPLY",
	"NO_REPLY",
}

// NewEnterpriseCommentListGetV10FilterReplyStatusByEDouyinFromValue returns a pointer to a valid EnterpriseCommentListGetV10FilterReplyStatusByEDouyin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10FilterReplyStatusByEDouyinFromValue(v string) (*EnterpriseCommentListGetV10FilterReplyStatusByEDouyin, error) {
	ev := EnterpriseCommentListGetV10FilterReplyStatusByEDouyin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10FilterReplyStatusByEDouyin: valid values are %v", v, AllowedEnterpriseCommentListGetV10FilterReplyStatusByEDouyinEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10FilterReplyStatusByEDouyin) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10FilterReplyStatusByEDouyinEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_reply_status_by_e_douyin value
func (v EnterpriseCommentListGetV10FilterReplyStatusByEDouyin) Ptr() *EnterpriseCommentListGetV10FilterReplyStatusByEDouyin {
	return &v
}
