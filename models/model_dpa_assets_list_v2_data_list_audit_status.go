/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaAssetsListV2DataListAuditStatus
type DpaAssetsListV2DataListAuditStatus string

// List of dpa_assets_list_v2_data_list_audit_status
const (
	ASSETS_AUDIT_DENY_DpaAssetsListV2DataListAuditStatus      DpaAssetsListV2DataListAuditStatus = "ASSETS_AUDIT_DENY"
	ASSETS_PENDING_AUDIT_DpaAssetsListV2DataListAuditStatus   DpaAssetsListV2DataListAuditStatus = "ASSETS_PENDING_AUDIT"
	ASSETS_AUDIT_SUCCESS_DpaAssetsListV2DataListAuditStatus   DpaAssetsListV2DataListAuditStatus = "ASSETS_AUDIT_SUCCESS"
	ASSETS_TO_SUBMIT_AUDIT_DpaAssetsListV2DataListAuditStatus DpaAssetsListV2DataListAuditStatus = "ASSETS_TO_SUBMIT_AUDIT"
)

// All allowed values of DpaAssetsListV2DataListAuditStatus enum
var AllowedDpaAssetsListV2DataListAuditStatusEnumValues = []DpaAssetsListV2DataListAuditStatus{
	"ASSETS_AUDIT_DENY",
	"ASSETS_PENDING_AUDIT",
	"ASSETS_AUDIT_SUCCESS",
	"ASSETS_TO_SUBMIT_AUDIT",
}

// NewDpaAssetsListV2DataListAuditStatusFromValue returns a pointer to a valid DpaAssetsListV2DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsListV2DataListAuditStatusFromValue(v string) (*DpaAssetsListV2DataListAuditStatus, error) {
	ev := DpaAssetsListV2DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsListV2DataListAuditStatus: valid values are %v", v, AllowedDpaAssetsListV2DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsListV2DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetsListV2DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_list_v2_data_list_audit_status value
func (v DpaAssetsListV2DataListAuditStatus) Ptr() *DpaAssetsListV2DataListAuditStatus {
	return &v
}