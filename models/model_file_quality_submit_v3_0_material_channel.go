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

// FileQualitySubmitV30MaterialChannel
type FileQualitySubmitV30MaterialChannel string

// List of file_quality_submit_v3.0_material_channel
const (
	AD_FileQualitySubmitV30MaterialChannel       FileQualitySubmitV30MaterialChannel = "AD"
	DOU_PLUS_FileQualitySubmitV30MaterialChannel FileQualitySubmitV30MaterialChannel = "DOU_PLUS"
	QC_FileQualitySubmitV30MaterialChannel       FileQualitySubmitV30MaterialChannel = "QC"
)

// All allowed values of FileQualitySubmitV30MaterialChannel enum
var AllowedFileQualitySubmitV30MaterialChannelEnumValues = []FileQualitySubmitV30MaterialChannel{
	"AD",
	"DOU_PLUS",
	"QC",
}

// NewFileQualitySubmitV30MaterialChannelFromValue returns a pointer to a valid FileQualitySubmitV30MaterialChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileQualitySubmitV30MaterialChannelFromValue(v string) (*FileQualitySubmitV30MaterialChannel, error) {
	ev := FileQualitySubmitV30MaterialChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileQualitySubmitV30MaterialChannel: valid values are %v", v, AllowedFileQualitySubmitV30MaterialChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileQualitySubmitV30MaterialChannel) IsValid() bool {
	for _, existing := range AllowedFileQualitySubmitV30MaterialChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_quality_submit_v3.0_material_channel value
func (v FileQualitySubmitV30MaterialChannel) Ptr() *FileQualitySubmitV30MaterialChannel {
	return &v
}
