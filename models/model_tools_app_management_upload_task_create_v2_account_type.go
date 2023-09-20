/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementUploadTaskCreateV2AccountType
type ToolsAppManagementUploadTaskCreateV2AccountType string

// List of tools_app_management_upload_task_create_v2_account_type
const (
	AD_ToolsAppManagementUploadTaskCreateV2AccountType ToolsAppManagementUploadTaskCreateV2AccountType = "AD"
	BP_ToolsAppManagementUploadTaskCreateV2AccountType ToolsAppManagementUploadTaskCreateV2AccountType = "BP"
)

// All allowed values of ToolsAppManagementUploadTaskCreateV2AccountType enum
var AllowedToolsAppManagementUploadTaskCreateV2AccountTypeEnumValues = []ToolsAppManagementUploadTaskCreateV2AccountType{
	"AD",
	"BP",
}

// NewToolsAppManagementUploadTaskCreateV2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementUploadTaskCreateV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementUploadTaskCreateV2AccountTypeFromValue(v string) (*ToolsAppManagementUploadTaskCreateV2AccountType, error) {
	ev := ToolsAppManagementUploadTaskCreateV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementUploadTaskCreateV2AccountType: valid values are %v", v, AllowedToolsAppManagementUploadTaskCreateV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementUploadTaskCreateV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementUploadTaskCreateV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_upload_task_create_v2_account_type value
func (v ToolsAppManagementUploadTaskCreateV2AccountType) Ptr() *ToolsAppManagementUploadTaskCreateV2AccountType {
	return &v
}
