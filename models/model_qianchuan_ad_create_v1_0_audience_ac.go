/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceAc
type QianchuanAdCreateV10AudienceAc string

// List of qianchuan_ad_create_v1.0_audience_ac
const (
	Enum_2_G_QianchuanAdCreateV10AudienceAc QianchuanAdCreateV10AudienceAc = "2G"
	Enum_3_G_QianchuanAdCreateV10AudienceAc QianchuanAdCreateV10AudienceAc = "3G"
	Enum_4_G_QianchuanAdCreateV10AudienceAc QianchuanAdCreateV10AudienceAc = "4G"
	Enum_5_G_QianchuanAdCreateV10AudienceAc QianchuanAdCreateV10AudienceAc = "5G"
	WIFI_QianchuanAdCreateV10AudienceAc     QianchuanAdCreateV10AudienceAc = "WIFI"
)

// All allowed values of QianchuanAdCreateV10AudienceAc enum
var AllowedQianchuanAdCreateV10AudienceAcEnumValues = []QianchuanAdCreateV10AudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewQianchuanAdCreateV10AudienceAcFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceAcFromValue(v string) (*QianchuanAdCreateV10AudienceAc, error) {
	ev := QianchuanAdCreateV10AudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceAc: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceAc) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_ac value
func (v QianchuanAdCreateV10AudienceAc) Ptr() *QianchuanAdCreateV10AudienceAc {
	return &v
}