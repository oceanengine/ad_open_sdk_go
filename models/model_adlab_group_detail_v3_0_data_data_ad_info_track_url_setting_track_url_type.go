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

// AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType
type AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType string

// List of adlab_group_detail_v3.0_data_data_ad_info_track_url_setting_track_url_type
const (
	CUSTOM_AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType  AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType = "CUSTOM"
	EXISTED_AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType = "EXISTED"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType{
	"CUSTOM",
	"EXISTED",
}

// NewAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_track_url_setting_track_url_type value
func (v AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType) Ptr() *AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType {
	return &v
}
