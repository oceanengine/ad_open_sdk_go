/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileUploadTaskCreateV2AccountType
type FileUploadTaskCreateV2AccountType string

// List of file_upload_task_create_v2_account_type
const (
	ADVERTISER_FileUploadTaskCreateV2AccountType FileUploadTaskCreateV2AccountType = "ADVERTISER"
	AGENT_FileUploadTaskCreateV2AccountType      FileUploadTaskCreateV2AccountType = "AGENT"
)

// All allowed values of FileUploadTaskCreateV2AccountType enum
var AllowedFileUploadTaskCreateV2AccountTypeEnumValues = []FileUploadTaskCreateV2AccountType{
	"ADVERTISER",
	"AGENT",
}

// NewFileUploadTaskCreateV2AccountTypeFromValue returns a pointer to a valid FileUploadTaskCreateV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileUploadTaskCreateV2AccountTypeFromValue(v string) (*FileUploadTaskCreateV2AccountType, error) {
	ev := FileUploadTaskCreateV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileUploadTaskCreateV2AccountType: valid values are %v", v, AllowedFileUploadTaskCreateV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileUploadTaskCreateV2AccountType) IsValid() bool {
	for _, existing := range AllowedFileUploadTaskCreateV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_upload_task_create_v2_account_type value
func (v FileUploadTaskCreateV2AccountType) Ptr() *FileUploadTaskCreateV2AccountType {
	return &v
}
