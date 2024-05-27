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

// CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus
type CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus int64

// List of creative_custom_creative_update_v2_ad_data_ad_download_status
const (
	Enum_0_CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus = 0
	Enum_1_CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus = 1
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus enum
var AllowedCreativeCustomCreativeUpdateV2AdDataAdDownloadStatusEnumValues = []CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus{
	0,
	1,
}

// NewCreativeCustomCreativeUpdateV2AdDataAdDownloadStatusFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataAdDownloadStatusFromValue(v int64) (*CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataAdDownloadStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataAdDownloadStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_ad_download_status value
func (v CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus) Ptr() *CreativeCustomCreativeUpdateV2AdDataAdDownloadStatus {
	return &v
}
