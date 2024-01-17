/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAppType
type AdGetV2DataAppType string

// List of ad_get_v2_data_app_type
const (
	APP_IOS_AdGetV2DataAppType     AdGetV2DataAppType = "APP_IOS"
	APP_ANDROID_AdGetV2DataAppType AdGetV2DataAppType = "APP_ANDROID"
)

// All allowed values of AdGetV2DataAppType enum
var AllowedAdGetV2DataAppTypeEnumValues = []AdGetV2DataAppType{
	"APP_IOS",
	"APP_ANDROID",
}

// NewAdGetV2DataAppTypeFromValue returns a pointer to a valid AdGetV2DataAppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAppTypeFromValue(v string) (*AdGetV2DataAppType, error) {
	ev := AdGetV2DataAppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAppType: valid values are %v", v, AllowedAdGetV2DataAppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAppType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_app_type value
func (v AdGetV2DataAppType) Ptr() *AdGetV2DataAppType {
	return &v
}
