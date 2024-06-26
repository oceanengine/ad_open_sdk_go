/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileMaterialAuditResultGetV2DataListAuditStatus
type FileMaterialAuditResultGetV2DataListAuditStatus string

// List of file_material_audit_result_get_v2_data_list_audit_status
const (
	PROCESS_FileMaterialAuditResultGetV2DataListAuditStatus FileMaterialAuditResultGetV2DataListAuditStatus = "PROCESS"
	PASS_FileMaterialAuditResultGetV2DataListAuditStatus    FileMaterialAuditResultGetV2DataListAuditStatus = "PASS"
	RISK_FileMaterialAuditResultGetV2DataListAuditStatus    FileMaterialAuditResultGetV2DataListAuditStatus = "RISK"
	REJECT_FileMaterialAuditResultGetV2DataListAuditStatus  FileMaterialAuditResultGetV2DataListAuditStatus = "REJECT"
	ERROR_FileMaterialAuditResultGetV2DataListAuditStatus   FileMaterialAuditResultGetV2DataListAuditStatus = "ERROR"
)

// All allowed values of FileMaterialAuditResultGetV2DataListAuditStatus enum
var AllowedFileMaterialAuditResultGetV2DataListAuditStatusEnumValues = []FileMaterialAuditResultGetV2DataListAuditStatus{
	"PROCESS",
	"PASS",
	"RISK",
	"REJECT",
	"ERROR",
}

// NewFileMaterialAuditResultGetV2DataListAuditStatusFromValue returns a pointer to a valid FileMaterialAuditResultGetV2DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialAuditResultGetV2DataListAuditStatusFromValue(v string) (*FileMaterialAuditResultGetV2DataListAuditStatus, error) {
	ev := FileMaterialAuditResultGetV2DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialAuditResultGetV2DataListAuditStatus: valid values are %v", v, AllowedFileMaterialAuditResultGetV2DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialAuditResultGetV2DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedFileMaterialAuditResultGetV2DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_audit_result_get_v2_data_list_audit_status value
func (v FileMaterialAuditResultGetV2DataListAuditStatus) Ptr() *FileMaterialAuditResultGetV2DataListAuditStatus {
	return &v
}
