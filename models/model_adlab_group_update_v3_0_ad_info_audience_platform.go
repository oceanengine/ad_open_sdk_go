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

// AdlabGroupUpdateV30AdInfoAudiencePlatform
type AdlabGroupUpdateV30AdInfoAudiencePlatform string

// List of adlab_group_update_v3.0_ad_info_audience_platform
const (
	ANDROID_AdlabGroupUpdateV30AdInfoAudiencePlatform AdlabGroupUpdateV30AdInfoAudiencePlatform = "ANDROID"
	IOS_AdlabGroupUpdateV30AdInfoAudiencePlatform     AdlabGroupUpdateV30AdInfoAudiencePlatform = "IOS"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudiencePlatform enum
var AllowedAdlabGroupUpdateV30AdInfoAudiencePlatformEnumValues = []AdlabGroupUpdateV30AdInfoAudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewAdlabGroupUpdateV30AdInfoAudiencePlatformFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudiencePlatformFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudiencePlatform, error) {
	ev := AdlabGroupUpdateV30AdInfoAudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudiencePlatform: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudiencePlatform) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_platform value
func (v AdlabGroupUpdateV30AdInfoAudiencePlatform) Ptr() *AdlabGroupUpdateV30AdInfoAudiencePlatform {
	return &v
}
