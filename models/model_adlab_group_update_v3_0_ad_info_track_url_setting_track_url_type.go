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

// AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType
type AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType string

// List of adlab_group_update_v3.0_ad_info_track_url_setting_track_url_type
const (
	CUSTOM_AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType  AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType = "CUSTOM"
	EXISTED_AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType = "EXISTED"
)

// All allowed values of AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType enum
var AllowedAdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlTypeEnumValues = []AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType{
	"CUSTOM",
	"EXISTED",
}

// NewAdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlTypeFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlTypeFromValue(v string) (*AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType, error) {
	ev := AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_track_url_setting_track_url_type value
func (v AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType) Ptr() *AdlabGroupUpdateV30AdInfoTrackUrlSettingTrackUrlType {
	return &v
}
