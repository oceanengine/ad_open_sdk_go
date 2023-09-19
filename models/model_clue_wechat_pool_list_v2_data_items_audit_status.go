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

// ClueWechatPoolListV2DataItemsAuditStatus
type ClueWechatPoolListV2DataItemsAuditStatus string

// List of clue_wechat_pool_list_v2_data_items_audit_status
const (
	AUDITING_ClueWechatPoolListV2DataItemsAuditStatus       ClueWechatPoolListV2DataItemsAuditStatus = "AUDITING"
	AUDIT_REJECTED_ClueWechatPoolListV2DataItemsAuditStatus ClueWechatPoolListV2DataItemsAuditStatus = "AUDIT_REJECTED"
	AUDIT_ACCEPTED_ClueWechatPoolListV2DataItemsAuditStatus ClueWechatPoolListV2DataItemsAuditStatus = "AUDIT_ACCEPTED"
)

// All allowed values of ClueWechatPoolListV2DataItemsAuditStatus enum
var AllowedClueWechatPoolListV2DataItemsAuditStatusEnumValues = []ClueWechatPoolListV2DataItemsAuditStatus{
	"AUDITING",
	"AUDIT_REJECTED",
	"AUDIT_ACCEPTED",
}

// NewClueWechatPoolListV2DataItemsAuditStatusFromValue returns a pointer to a valid ClueWechatPoolListV2DataItemsAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatPoolListV2DataItemsAuditStatusFromValue(v string) (*ClueWechatPoolListV2DataItemsAuditStatus, error) {
	ev := ClueWechatPoolListV2DataItemsAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatPoolListV2DataItemsAuditStatus: valid values are %v", v, AllowedClueWechatPoolListV2DataItemsAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatPoolListV2DataItemsAuditStatus) IsValid() bool {
	for _, existing := range AllowedClueWechatPoolListV2DataItemsAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_pool_list_v2_data_items_audit_status value
func (v ClueWechatPoolListV2DataItemsAuditStatus) Ptr() *ClueWechatPoolListV2DataItemsAuditStatus {
	return &v
}
