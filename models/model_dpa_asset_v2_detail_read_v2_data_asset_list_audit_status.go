/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaAssetV2DetailReadV2DataAssetListAuditStatus
type DpaAssetV2DetailReadV2DataAssetListAuditStatus string

// List of dpa_asset_v2_detail_read_v2_data_asset_list_audit_status
const (
	REVIEW_FAIL_DpaAssetV2DetailReadV2DataAssetListAuditStatus    DpaAssetV2DetailReadV2DataAssetListAuditStatus = "REVIEW_FAIL"
	REVIEW_SUCCESS_DpaAssetV2DetailReadV2DataAssetListAuditStatus DpaAssetV2DetailReadV2DataAssetListAuditStatus = "REVIEW_SUCCESS"
	UNDER_REVIEW_DpaAssetV2DetailReadV2DataAssetListAuditStatus   DpaAssetV2DetailReadV2DataAssetListAuditStatus = "UNDER_REVIEW"
	UNREVIEWED_DpaAssetV2DetailReadV2DataAssetListAuditStatus     DpaAssetV2DetailReadV2DataAssetListAuditStatus = "UNREVIEWED"
)

// All allowed values of DpaAssetV2DetailReadV2DataAssetListAuditStatus enum
var AllowedDpaAssetV2DetailReadV2DataAssetListAuditStatusEnumValues = []DpaAssetV2DetailReadV2DataAssetListAuditStatus{
	"REVIEW_FAIL",
	"REVIEW_SUCCESS",
	"UNDER_REVIEW",
	"UNREVIEWED",
}

// NewDpaAssetV2DetailReadV2DataAssetListAuditStatusFromValue returns a pointer to a valid DpaAssetV2DetailReadV2DataAssetListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetV2DetailReadV2DataAssetListAuditStatusFromValue(v string) (*DpaAssetV2DetailReadV2DataAssetListAuditStatus, error) {
	ev := DpaAssetV2DetailReadV2DataAssetListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetV2DetailReadV2DataAssetListAuditStatus: valid values are %v", v, AllowedDpaAssetV2DetailReadV2DataAssetListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetV2DetailReadV2DataAssetListAuditStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetV2DetailReadV2DataAssetListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_asset_v2_detail_read_v2_data_asset_list_audit_status value
func (v DpaAssetV2DetailReadV2DataAssetListAuditStatus) Ptr() *DpaAssetV2DetailReadV2DataAssetListAuditStatus {
	return &v
}
