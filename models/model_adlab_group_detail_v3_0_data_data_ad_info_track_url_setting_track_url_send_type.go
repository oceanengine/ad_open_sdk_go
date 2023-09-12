/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType
type AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType string

// List of adlab_group_detail_v3.0_data_data_ad_info_track_url_setting_track_url_send_type
const (
	CLIENT_SEND_AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType = "CLIENT_SEND"
	SERVER_SEND_AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType = "SERVER_SEND"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType{
	"CLIENT_SEND",
	"SERVER_SEND",
}

// NewAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_track_url_setting_track_url_send_type value
func (v AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType) Ptr() *AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType {
	return &v
}
