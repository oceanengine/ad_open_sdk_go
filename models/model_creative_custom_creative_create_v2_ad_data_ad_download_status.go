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

// CreativeCustomCreativeCreateV2AdDataAdDownloadStatus
type CreativeCustomCreativeCreateV2AdDataAdDownloadStatus int64

// List of creative_custom_creative_create_v2_ad_data_ad_download_status
const (
	Enum_0_CreativeCustomCreativeCreateV2AdDataAdDownloadStatus CreativeCustomCreativeCreateV2AdDataAdDownloadStatus = 0
	Enum_1_CreativeCustomCreativeCreateV2AdDataAdDownloadStatus CreativeCustomCreativeCreateV2AdDataAdDownloadStatus = 1
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataAdDownloadStatus enum
var AllowedCreativeCustomCreativeCreateV2AdDataAdDownloadStatusEnumValues = []CreativeCustomCreativeCreateV2AdDataAdDownloadStatus{
	0,
	1,
}

// NewCreativeCustomCreativeCreateV2AdDataAdDownloadStatusFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataAdDownloadStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataAdDownloadStatusFromValue(v int64) (*CreativeCustomCreativeCreateV2AdDataAdDownloadStatus, error) {
	ev := CreativeCustomCreativeCreateV2AdDataAdDownloadStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataAdDownloadStatus: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataAdDownloadStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataAdDownloadStatus) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataAdDownloadStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_ad_download_status value
func (v CreativeCustomCreativeCreateV2AdDataAdDownloadStatus) Ptr() *CreativeCustomCreativeCreateV2AdDataAdDownloadStatus {
	return &v
}
