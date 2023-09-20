/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudiencePlatform
type QianchuanToolsEstimateAudienceV10AudiencePlatform string

// List of qianchuan_tools_estimate_audience_v1.0_audience_platform
const (
	ANDROID_QianchuanToolsEstimateAudienceV10AudiencePlatform QianchuanToolsEstimateAudienceV10AudiencePlatform = "ANDROID"
	IOS_QianchuanToolsEstimateAudienceV10AudiencePlatform     QianchuanToolsEstimateAudienceV10AudiencePlatform = "IOS"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudiencePlatform enum
var AllowedQianchuanToolsEstimateAudienceV10AudiencePlatformEnumValues = []QianchuanToolsEstimateAudienceV10AudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewQianchuanToolsEstimateAudienceV10AudiencePlatformFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudiencePlatformFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudiencePlatform, error) {
	ev := QianchuanToolsEstimateAudienceV10AudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudiencePlatform: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudiencePlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_platform value
func (v QianchuanToolsEstimateAudienceV10AudiencePlatform) Ptr() *QianchuanToolsEstimateAudienceV10AudiencePlatform {
	return &v
}
