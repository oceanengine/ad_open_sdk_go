/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsDownloadPackageGetV2DataPackageStatus
type ToolsDownloadPackageGetV2DataPackageStatus string

// List of tools_download_package_get_v2_data_package_status
const (
	NEW_PACKAGE_ToolsDownloadPackageGetV2DataPackageStatus ToolsDownloadPackageGetV2DataPackageStatus = "NEW_PACKAGE"
	PARSING_ToolsDownloadPackageGetV2DataPackageStatus     ToolsDownloadPackageGetV2DataPackageStatus = "PARSING"
	SUCCESS_ToolsDownloadPackageGetV2DataPackageStatus     ToolsDownloadPackageGetV2DataPackageStatus = "SUCCESS"
	FAIL_ToolsDownloadPackageGetV2DataPackageStatus        ToolsDownloadPackageGetV2DataPackageStatus = "FAIL"
)

// All allowed values of ToolsDownloadPackageGetV2DataPackageStatus enum
var AllowedToolsDownloadPackageGetV2DataPackageStatusEnumValues = []ToolsDownloadPackageGetV2DataPackageStatus{
	"NEW_PACKAGE",
	"PARSING",
	"SUCCESS",
	"FAIL",
}

// NewToolsDownloadPackageGetV2DataPackageStatusFromValue returns a pointer to a valid ToolsDownloadPackageGetV2DataPackageStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsDownloadPackageGetV2DataPackageStatusFromValue(v string) (*ToolsDownloadPackageGetV2DataPackageStatus, error) {
	ev := ToolsDownloadPackageGetV2DataPackageStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsDownloadPackageGetV2DataPackageStatus: valid values are %v", v, AllowedToolsDownloadPackageGetV2DataPackageStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsDownloadPackageGetV2DataPackageStatus) IsValid() bool {
	for _, existing := range AllowedToolsDownloadPackageGetV2DataPackageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_download_package_get_v2_data_package_status value
func (v ToolsDownloadPackageGetV2DataPackageStatus) Ptr() *ToolsDownloadPackageGetV2DataPackageStatus {
	return &v
}
