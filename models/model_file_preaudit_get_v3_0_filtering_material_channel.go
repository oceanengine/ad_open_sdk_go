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

// FilePreauditGetV30FilteringMaterialChannel
type FilePreauditGetV30FilteringMaterialChannel string

// List of file_preaudit_get_v3.0_filtering_material_channel
const (
	AD_FilePreauditGetV30FilteringMaterialChannel       FilePreauditGetV30FilteringMaterialChannel = "AD"
	DOU_PLUS_FilePreauditGetV30FilteringMaterialChannel FilePreauditGetV30FilteringMaterialChannel = "DOU_PLUS"
	QC_FilePreauditGetV30FilteringMaterialChannel       FilePreauditGetV30FilteringMaterialChannel = "QC"
)

// All allowed values of FilePreauditGetV30FilteringMaterialChannel enum
var AllowedFilePreauditGetV30FilteringMaterialChannelEnumValues = []FilePreauditGetV30FilteringMaterialChannel{
	"AD",
	"DOU_PLUS",
	"QC",
}

// NewFilePreauditGetV30FilteringMaterialChannelFromValue returns a pointer to a valid FilePreauditGetV30FilteringMaterialChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilePreauditGetV30FilteringMaterialChannelFromValue(v string) (*FilePreauditGetV30FilteringMaterialChannel, error) {
	ev := FilePreauditGetV30FilteringMaterialChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilePreauditGetV30FilteringMaterialChannel: valid values are %v", v, AllowedFilePreauditGetV30FilteringMaterialChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilePreauditGetV30FilteringMaterialChannel) IsValid() bool {
	for _, existing := range AllowedFilePreauditGetV30FilteringMaterialChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_preaudit_get_v3.0_filtering_material_channel value
func (v FilePreauditGetV30FilteringMaterialChannel) Ptr() *FilePreauditGetV30FilteringMaterialChannel {
	return &v
}
