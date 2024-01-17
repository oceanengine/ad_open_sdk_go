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

// DpaClueProductListV2FilteringAuditStatus
type DpaClueProductListV2FilteringAuditStatus string

// List of dpa_clue_product_list_v2_filtering_audit_status
const (
	AUDIT_STATUS_APPROVE_DpaClueProductListV2FilteringAuditStatus DpaClueProductListV2FilteringAuditStatus = "AUDIT_STATUS_APPROVE"
	AUDIT_STATUS_INIT_DpaClueProductListV2FilteringAuditStatus    DpaClueProductListV2FilteringAuditStatus = "AUDIT_STATUS_INIT"
	AUDIT_STATUS_REJECT_DpaClueProductListV2FilteringAuditStatus  DpaClueProductListV2FilteringAuditStatus = "AUDIT_STATUS_REJECT"
)

// All allowed values of DpaClueProductListV2FilteringAuditStatus enum
var AllowedDpaClueProductListV2FilteringAuditStatusEnumValues = []DpaClueProductListV2FilteringAuditStatus{
	"AUDIT_STATUS_APPROVE",
	"AUDIT_STATUS_INIT",
	"AUDIT_STATUS_REJECT",
}

// NewDpaClueProductListV2FilteringAuditStatusFromValue returns a pointer to a valid DpaClueProductListV2FilteringAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductListV2FilteringAuditStatusFromValue(v string) (*DpaClueProductListV2FilteringAuditStatus, error) {
	ev := DpaClueProductListV2FilteringAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductListV2FilteringAuditStatus: valid values are %v", v, AllowedDpaClueProductListV2FilteringAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductListV2FilteringAuditStatus) IsValid() bool {
	for _, existing := range AllowedDpaClueProductListV2FilteringAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_list_v2_filtering_audit_status value
func (v DpaClueProductListV2FilteringAuditStatus) Ptr() *DpaClueProductListV2FilteringAuditStatus {
	return &v
}
