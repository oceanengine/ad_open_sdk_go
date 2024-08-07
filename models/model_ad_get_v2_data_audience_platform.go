/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudiencePlatform
type AdGetV2DataAudiencePlatform string

// List of ad_get_v2_data_audience_platform
const (
	IPAD_AdGetV2DataAudiencePlatform    AdGetV2DataAudiencePlatform = "IPAD"
	NONE_AdGetV2DataAudiencePlatform    AdGetV2DataAudiencePlatform = "NONE"
	ANDROID_AdGetV2DataAudiencePlatform AdGetV2DataAudiencePlatform = "ANDROID"
	PC_AdGetV2DataAudiencePlatform      AdGetV2DataAudiencePlatform = "PC"
	IOS_AdGetV2DataAudiencePlatform     AdGetV2DataAudiencePlatform = "IOS"
	WAP_AdGetV2DataAudiencePlatform     AdGetV2DataAudiencePlatform = "WAP"
)

// All allowed values of AdGetV2DataAudiencePlatform enum
var AllowedAdGetV2DataAudiencePlatformEnumValues = []AdGetV2DataAudiencePlatform{
	"IPAD",
	"NONE",
	"ANDROID",
	"PC",
	"IOS",
	"WAP",
}

// NewAdGetV2DataAudiencePlatformFromValue returns a pointer to a valid AdGetV2DataAudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudiencePlatformFromValue(v string) (*AdGetV2DataAudiencePlatform, error) {
	ev := AdGetV2DataAudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudiencePlatform: valid values are %v", v, AllowedAdGetV2DataAudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudiencePlatform) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_platform value
func (v AdGetV2DataAudiencePlatform) Ptr() *AdGetV2DataAudiencePlatform {
	return &v
}
