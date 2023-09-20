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

// QianchuanAdUpdateV10AudienceAc
type QianchuanAdUpdateV10AudienceAc string

// List of qianchuan_ad_update_v1.0_audience_ac
const (
	Enum_2_G_QianchuanAdUpdateV10AudienceAc QianchuanAdUpdateV10AudienceAc = "2G"
	Enum_3_G_QianchuanAdUpdateV10AudienceAc QianchuanAdUpdateV10AudienceAc = "3G"
	Enum_4_G_QianchuanAdUpdateV10AudienceAc QianchuanAdUpdateV10AudienceAc = "4G"
	Enum_5_G_QianchuanAdUpdateV10AudienceAc QianchuanAdUpdateV10AudienceAc = "5G"
	WIFI_QianchuanAdUpdateV10AudienceAc     QianchuanAdUpdateV10AudienceAc = "WIFI"
)

// All allowed values of QianchuanAdUpdateV10AudienceAc enum
var AllowedQianchuanAdUpdateV10AudienceAcEnumValues = []QianchuanAdUpdateV10AudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewQianchuanAdUpdateV10AudienceAcFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceAcFromValue(v string) (*QianchuanAdUpdateV10AudienceAc, error) {
	ev := QianchuanAdUpdateV10AudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceAc: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceAc) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_ac value
func (v QianchuanAdUpdateV10AudienceAc) Ptr() *QianchuanAdUpdateV10AudienceAc {
	return &v
}
