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

// ClueWechatPoolListV2FilterAuditStatus
type ClueWechatPoolListV2FilterAuditStatus string

// List of clue_wechat_pool_list_v2_filter_audit_status
const (
	AUDITING_ClueWechatPoolListV2FilterAuditStatus       ClueWechatPoolListV2FilterAuditStatus = "AUDITING"
	AUDIT_REJECTED_ClueWechatPoolListV2FilterAuditStatus ClueWechatPoolListV2FilterAuditStatus = "AUDIT_REJECTED"
	AUDIT_ACCEPTED_ClueWechatPoolListV2FilterAuditStatus ClueWechatPoolListV2FilterAuditStatus = "AUDIT_ACCEPTED"
)

// All allowed values of ClueWechatPoolListV2FilterAuditStatus enum
var AllowedClueWechatPoolListV2FilterAuditStatusEnumValues = []ClueWechatPoolListV2FilterAuditStatus{
	"AUDITING",
	"AUDIT_REJECTED",
	"AUDIT_ACCEPTED",
}

// NewClueWechatPoolListV2FilterAuditStatusFromValue returns a pointer to a valid ClueWechatPoolListV2FilterAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatPoolListV2FilterAuditStatusFromValue(v string) (*ClueWechatPoolListV2FilterAuditStatus, error) {
	ev := ClueWechatPoolListV2FilterAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatPoolListV2FilterAuditStatus: valid values are %v", v, AllowedClueWechatPoolListV2FilterAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatPoolListV2FilterAuditStatus) IsValid() bool {
	for _, existing := range AllowedClueWechatPoolListV2FilterAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_pool_list_v2_filter_audit_status value
func (v ClueWechatPoolListV2FilterAuditStatus) Ptr() *ClueWechatPoolListV2FilterAuditStatus {
	return &v
}
