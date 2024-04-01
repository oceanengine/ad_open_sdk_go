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

// FilePreauditSubmitV30MaterialChannel
type FilePreauditSubmitV30MaterialChannel string

// List of file_preaudit_submit_v3.0_material_channel
const (
	AD_FilePreauditSubmitV30MaterialChannel       FilePreauditSubmitV30MaterialChannel = "AD"
	DOU_PLUS_FilePreauditSubmitV30MaterialChannel FilePreauditSubmitV30MaterialChannel = "DOU_PLUS"
	QC_FilePreauditSubmitV30MaterialChannel       FilePreauditSubmitV30MaterialChannel = "QC"
)

// All allowed values of FilePreauditSubmitV30MaterialChannel enum
var AllowedFilePreauditSubmitV30MaterialChannelEnumValues = []FilePreauditSubmitV30MaterialChannel{
	"AD",
	"DOU_PLUS",
	"QC",
}

// NewFilePreauditSubmitV30MaterialChannelFromValue returns a pointer to a valid FilePreauditSubmitV30MaterialChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilePreauditSubmitV30MaterialChannelFromValue(v string) (*FilePreauditSubmitV30MaterialChannel, error) {
	ev := FilePreauditSubmitV30MaterialChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilePreauditSubmitV30MaterialChannel: valid values are %v", v, AllowedFilePreauditSubmitV30MaterialChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilePreauditSubmitV30MaterialChannel) IsValid() bool {
	for _, existing := range AllowedFilePreauditSubmitV30MaterialChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_preaudit_submit_v3.0_material_channel value
func (v FilePreauditSubmitV30MaterialChannel) Ptr() *FilePreauditSubmitV30MaterialChannel {
	return &v
}
