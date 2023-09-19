/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataAudiencePlatform
type QianchuanAdDetailGetV10DataAudiencePlatform string

// List of qianchuan_ad_detail_get_v1.0_data_audience_platform
const (
	ANDROID_QianchuanAdDetailGetV10DataAudiencePlatform QianchuanAdDetailGetV10DataAudiencePlatform = "ANDROID"
	IOS_QianchuanAdDetailGetV10DataAudiencePlatform     QianchuanAdDetailGetV10DataAudiencePlatform = "IOS"
	IPAD_QianchuanAdDetailGetV10DataAudiencePlatform    QianchuanAdDetailGetV10DataAudiencePlatform = "IPAD"
	NONE_QianchuanAdDetailGetV10DataAudiencePlatform    QianchuanAdDetailGetV10DataAudiencePlatform = "NONE"
	PC_QianchuanAdDetailGetV10DataAudiencePlatform      QianchuanAdDetailGetV10DataAudiencePlatform = "PC"
	WAP_QianchuanAdDetailGetV10DataAudiencePlatform     QianchuanAdDetailGetV10DataAudiencePlatform = "WAP"
)

// All allowed values of QianchuanAdDetailGetV10DataAudiencePlatform enum
var AllowedQianchuanAdDetailGetV10DataAudiencePlatformEnumValues = []QianchuanAdDetailGetV10DataAudiencePlatform{
	"ANDROID",
	"IOS",
	"IPAD",
	"NONE",
	"PC",
	"WAP",
}

// NewQianchuanAdDetailGetV10DataAudiencePlatformFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudiencePlatformFromValue(v string) (*QianchuanAdDetailGetV10DataAudiencePlatform, error) {
	ev := QianchuanAdDetailGetV10DataAudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudiencePlatform: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudiencePlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_platform value
func (v QianchuanAdDetailGetV10DataAudiencePlatform) Ptr() *QianchuanAdDetailGetV10DataAudiencePlatform {
	return &v
}
