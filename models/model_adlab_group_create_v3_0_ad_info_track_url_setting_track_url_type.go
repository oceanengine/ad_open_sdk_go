/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType
type AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType string

// List of adlab_group_create_v3.0_ad_info_track_url_setting_track_url_type
const (
	CUSTOM_AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType  AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType = "CUSTOM"
	EXISTED_AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType = "EXISTED"
)

// All allowed values of AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType enum
var AllowedAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlTypeEnumValues = []AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType{
	"CUSTOM",
	"EXISTED",
}

// NewAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType, error) {
	ev := AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_track_url_setting_track_url_type value
func (v AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType) Ptr() *AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType {
	return &v
}