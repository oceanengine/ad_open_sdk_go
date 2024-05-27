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

// AdlabGroupCreateV30AdInfoAudiencePlatform
type AdlabGroupCreateV30AdInfoAudiencePlatform string

// List of adlab_group_create_v3.0_ad_info_audience_platform
const (
	ANDROID_AdlabGroupCreateV30AdInfoAudiencePlatform AdlabGroupCreateV30AdInfoAudiencePlatform = "ANDROID"
	IOS_AdlabGroupCreateV30AdInfoAudiencePlatform     AdlabGroupCreateV30AdInfoAudiencePlatform = "IOS"
)

// All allowed values of AdlabGroupCreateV30AdInfoAudiencePlatform enum
var AllowedAdlabGroupCreateV30AdInfoAudiencePlatformEnumValues = []AdlabGroupCreateV30AdInfoAudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewAdlabGroupCreateV30AdInfoAudiencePlatformFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAudiencePlatformFromValue(v string) (*AdlabGroupCreateV30AdInfoAudiencePlatform, error) {
	ev := AdlabGroupCreateV30AdInfoAudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAudiencePlatform: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAudiencePlatform) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_audience_platform value
func (v AdlabGroupCreateV30AdInfoAudiencePlatform) Ptr() *AdlabGroupCreateV30AdInfoAudiencePlatform {
	return &v
}
