/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPioneerProgramAttachmentUploadV2DataFileType
type ToolsPioneerProgramAttachmentUploadV2DataFileType string

// List of tools_pioneer_program_attachment_upload_v2_data_file_type
const (
	AD_SETTING_DATA_ToolsPioneerProgramAttachmentUploadV2DataFileType       ToolsPioneerProgramAttachmentUploadV2DataFileType = "AD_SETTING_DATA"
	BASIC_STAT_DATA_ToolsPioneerProgramAttachmentUploadV2DataFileType       ToolsPioneerProgramAttachmentUploadV2DataFileType = "BASIC_STAT_DATA"
	DEEP_STAT_DAILY_DATA_ToolsPioneerProgramAttachmentUploadV2DataFileType  ToolsPioneerProgramAttachmentUploadV2DataFileType = "DEEP_STAT_DAILY_DATA"
	DEEP_STAT_HOURLY_DATA_ToolsPioneerProgramAttachmentUploadV2DataFileType ToolsPioneerProgramAttachmentUploadV2DataFileType = "DEEP_STAT_HOURLY_DATA"
	MATERIAL_MAPPING_DATA_ToolsPioneerProgramAttachmentUploadV2DataFileType ToolsPioneerProgramAttachmentUploadV2DataFileType = "MATERIAL_MAPPING_DATA"
)

// All allowed values of ToolsPioneerProgramAttachmentUploadV2DataFileType enum
var AllowedToolsPioneerProgramAttachmentUploadV2DataFileTypeEnumValues = []ToolsPioneerProgramAttachmentUploadV2DataFileType{
	"AD_SETTING_DATA",
	"BASIC_STAT_DATA",
	"DEEP_STAT_DAILY_DATA",
	"DEEP_STAT_HOURLY_DATA",
	"MATERIAL_MAPPING_DATA",
}

// NewToolsPioneerProgramAttachmentUploadV2DataFileTypeFromValue returns a pointer to a valid ToolsPioneerProgramAttachmentUploadV2DataFileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPioneerProgramAttachmentUploadV2DataFileTypeFromValue(v string) (*ToolsPioneerProgramAttachmentUploadV2DataFileType, error) {
	ev := ToolsPioneerProgramAttachmentUploadV2DataFileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPioneerProgramAttachmentUploadV2DataFileType: valid values are %v", v, AllowedToolsPioneerProgramAttachmentUploadV2DataFileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPioneerProgramAttachmentUploadV2DataFileType) IsValid() bool {
	for _, existing := range AllowedToolsPioneerProgramAttachmentUploadV2DataFileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_pioneer_program_attachment_upload_v2_data_file_type value
func (v ToolsPioneerProgramAttachmentUploadV2DataFileType) Ptr() *ToolsPioneerProgramAttachmentUploadV2DataFileType {
	return &v
}
