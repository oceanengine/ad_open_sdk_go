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

// AdGetV2DataLaunchTargetType
type AdGetV2DataLaunchTargetType string

// List of ad_get_v2_data_launch_target_type
const (
	LIVE_CONVERT_AdGetV2DataLaunchTargetType AdGetV2DataLaunchTargetType = "LIVE_CONVERT"
	APP_AdGetV2DataLaunchTargetType          AdGetV2DataLaunchTargetType = "APP"
	EXTERNAL_AdGetV2DataLaunchTargetType     AdGetV2DataLaunchTargetType = "EXTERNAL"
)

// All allowed values of AdGetV2DataLaunchTargetType enum
var AllowedAdGetV2DataLaunchTargetTypeEnumValues = []AdGetV2DataLaunchTargetType{
	"LIVE_CONVERT",
	"APP",
	"EXTERNAL",
}

// NewAdGetV2DataLaunchTargetTypeFromValue returns a pointer to a valid AdGetV2DataLaunchTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataLaunchTargetTypeFromValue(v string) (*AdGetV2DataLaunchTargetType, error) {
	ev := AdGetV2DataLaunchTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataLaunchTargetType: valid values are %v", v, AllowedAdGetV2DataLaunchTargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataLaunchTargetType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataLaunchTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_launch_target_type value
func (v AdGetV2DataLaunchTargetType) Ptr() *AdGetV2DataLaunchTargetType {
	return &v
}
