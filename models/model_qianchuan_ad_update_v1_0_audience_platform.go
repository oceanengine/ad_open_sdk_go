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

// QianchuanAdUpdateV10AudiencePlatform
type QianchuanAdUpdateV10AudiencePlatform string

// List of qianchuan_ad_update_v1.0_audience_platform
const (
	ANDROID_QianchuanAdUpdateV10AudiencePlatform QianchuanAdUpdateV10AudiencePlatform = "ANDROID"
	IOS_QianchuanAdUpdateV10AudiencePlatform     QianchuanAdUpdateV10AudiencePlatform = "IOS"
)

// All allowed values of QianchuanAdUpdateV10AudiencePlatform enum
var AllowedQianchuanAdUpdateV10AudiencePlatformEnumValues = []QianchuanAdUpdateV10AudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewQianchuanAdUpdateV10AudiencePlatformFromValue returns a pointer to a valid QianchuanAdUpdateV10AudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudiencePlatformFromValue(v string) (*QianchuanAdUpdateV10AudiencePlatform, error) {
	ev := QianchuanAdUpdateV10AudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudiencePlatform: valid values are %v", v, AllowedQianchuanAdUpdateV10AudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudiencePlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_platform value
func (v QianchuanAdUpdateV10AudiencePlatform) Ptr() *QianchuanAdUpdateV10AudiencePlatform {
	return &v
}
