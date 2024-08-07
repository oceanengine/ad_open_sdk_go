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

// DpaAssetsDetailReadV2DataListAuditStatus
type DpaAssetsDetailReadV2DataListAuditStatus string

// List of dpa_assets_detail_read_v2_data_list_audit_status
const (
	ASSETS_AUDIT_SUCCESS_DpaAssetsDetailReadV2DataListAuditStatus   DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_AUDIT_SUCCESS"
	ASSETS_TO_SUBMIT_AUDIT_DpaAssetsDetailReadV2DataListAuditStatus DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_TO_SUBMIT_AUDIT"
	ASSETS_PENDING_AUDIT_DpaAssetsDetailReadV2DataListAuditStatus   DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_PENDING_AUDIT"
	ASSETS_AUDIT_DENY_DpaAssetsDetailReadV2DataListAuditStatus      DpaAssetsDetailReadV2DataListAuditStatus = "ASSETS_AUDIT_DENY"
)

// All allowed values of DpaAssetsDetailReadV2DataListAuditStatus enum
var AllowedDpaAssetsDetailReadV2DataListAuditStatusEnumValues = []DpaAssetsDetailReadV2DataListAuditStatus{
	"ASSETS_AUDIT_SUCCESS",
	"ASSETS_TO_SUBMIT_AUDIT",
	"ASSETS_PENDING_AUDIT",
	"ASSETS_AUDIT_DENY",
}

// NewDpaAssetsDetailReadV2DataListAuditStatusFromValue returns a pointer to a valid DpaAssetsDetailReadV2DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsDetailReadV2DataListAuditStatusFromValue(v string) (*DpaAssetsDetailReadV2DataListAuditStatus, error) {
	ev := DpaAssetsDetailReadV2DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsDetailReadV2DataListAuditStatus: valid values are %v", v, AllowedDpaAssetsDetailReadV2DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsDetailReadV2DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetsDetailReadV2DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_detail_read_v2_data_list_audit_status value
func (v DpaAssetsDetailReadV2DataListAuditStatus) Ptr() *DpaAssetsDetailReadV2DataListAuditStatus {
	return &v
}
