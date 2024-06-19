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

// AdlabGroupListV30DataGroupsAdInfoDownloadType
type AdlabGroupListV30DataGroupsAdInfoDownloadType string

// List of adlab_group_list_v3.0_data_groups_ad_info_download_type
const (
	DOWNLOAD_URL_AdlabGroupListV30DataGroupsAdInfoDownloadType AdlabGroupListV30DataGroupsAdInfoDownloadType = "DOWNLOAD_URL"
	EXTERNAL_URL_AdlabGroupListV30DataGroupsAdInfoDownloadType AdlabGroupListV30DataGroupsAdInfoDownloadType = "EXTERNAL_URL"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoDownloadType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoDownloadTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoDownloadType{
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
}

// NewAdlabGroupListV30DataGroupsAdInfoDownloadTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoDownloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoDownloadTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoDownloadType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoDownloadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoDownloadType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoDownloadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoDownloadType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoDownloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_download_type value
func (v AdlabGroupListV30DataGroupsAdInfoDownloadType) Ptr() *AdlabGroupListV30DataGroupsAdInfoDownloadType {
	return &v
}
