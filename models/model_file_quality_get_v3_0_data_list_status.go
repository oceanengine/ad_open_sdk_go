/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileQualityGetV30DataListStatus
type FileQualityGetV30DataListStatus string

// List of file_quality_get_v3.0_data_list_status
const (
	AUDITING_FileQualityGetV30DataListStatus       FileQualityGetV30DataListStatus = "AUDITING"
	AUDIT_ACCEPTED_FileQualityGetV30DataListStatus FileQualityGetV30DataListStatus = "AUDIT_ACCEPTED"
	AUDIT_FAILED_FileQualityGetV30DataListStatus   FileQualityGetV30DataListStatus = "AUDIT_FAILED"
	AUDIT_REJECT_FileQualityGetV30DataListStatus   FileQualityGetV30DataListStatus = "AUDIT_REJECT"
	AUDIT_SUBMIT_FileQualityGetV30DataListStatus   FileQualityGetV30DataListStatus = "AUDIT_SUBMIT"
	AUDIT_TIMEOUT_FileQualityGetV30DataListStatus  FileQualityGetV30DataListStatus = "AUDIT_TIMEOUT"
)

// All allowed values of FileQualityGetV30DataListStatus enum
var AllowedFileQualityGetV30DataListStatusEnumValues = []FileQualityGetV30DataListStatus{
	"AUDITING",
	"AUDIT_ACCEPTED",
	"AUDIT_FAILED",
	"AUDIT_REJECT",
	"AUDIT_SUBMIT",
	"AUDIT_TIMEOUT",
}

// NewFileQualityGetV30DataListStatusFromValue returns a pointer to a valid FileQualityGetV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileQualityGetV30DataListStatusFromValue(v string) (*FileQualityGetV30DataListStatus, error) {
	ev := FileQualityGetV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileQualityGetV30DataListStatus: valid values are %v", v, AllowedFileQualityGetV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileQualityGetV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedFileQualityGetV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_quality_get_v3.0_data_list_status value
func (v FileQualityGetV30DataListStatus) Ptr() *FileQualityGetV30DataListStatus {
	return &v
}
