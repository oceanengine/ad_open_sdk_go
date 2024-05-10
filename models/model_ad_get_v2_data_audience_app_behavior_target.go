/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceAppBehaviorTarget
type AdGetV2DataAudienceAppBehaviorTarget string

// List of ad_get_v2_data_audience_app_behavior_target
const (
	NONE_AdGetV2DataAudienceAppBehaviorTarget     AdGetV2DataAudienceAppBehaviorTarget = "NONE"
	CATEGORY_AdGetV2DataAudienceAppBehaviorTarget AdGetV2DataAudienceAppBehaviorTarget = "CATEGORY"
	APP_AdGetV2DataAudienceAppBehaviorTarget      AdGetV2DataAudienceAppBehaviorTarget = "APP"
)

// All allowed values of AdGetV2DataAudienceAppBehaviorTarget enum
var AllowedAdGetV2DataAudienceAppBehaviorTargetEnumValues = []AdGetV2DataAudienceAppBehaviorTarget{
	"NONE",
	"CATEGORY",
	"APP",
}

// NewAdGetV2DataAudienceAppBehaviorTargetFromValue returns a pointer to a valid AdGetV2DataAudienceAppBehaviorTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceAppBehaviorTargetFromValue(v string) (*AdGetV2DataAudienceAppBehaviorTarget, error) {
	ev := AdGetV2DataAudienceAppBehaviorTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceAppBehaviorTarget: valid values are %v", v, AllowedAdGetV2DataAudienceAppBehaviorTargetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceAppBehaviorTarget) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceAppBehaviorTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_app_behavior_target value
func (v AdGetV2DataAudienceAppBehaviorTarget) Ptr() *AdGetV2DataAudienceAppBehaviorTarget {
	return &v
}
