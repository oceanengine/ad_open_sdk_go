/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceAc
type QianchuanToolsEstimateAudienceV10AudienceAc string

// List of qianchuan_tools_estimate_audience_v1.0_audience_ac
const (
	Enum_2_G_QianchuanToolsEstimateAudienceV10AudienceAc QianchuanToolsEstimateAudienceV10AudienceAc = "2G"
	Enum_3_G_QianchuanToolsEstimateAudienceV10AudienceAc QianchuanToolsEstimateAudienceV10AudienceAc = "3G"
	Enum_4_G_QianchuanToolsEstimateAudienceV10AudienceAc QianchuanToolsEstimateAudienceV10AudienceAc = "4G"
	Enum_5_G_QianchuanToolsEstimateAudienceV10AudienceAc QianchuanToolsEstimateAudienceV10AudienceAc = "5G"
	WIFI_QianchuanToolsEstimateAudienceV10AudienceAc     QianchuanToolsEstimateAudienceV10AudienceAc = "WIFI"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceAc enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceAcEnumValues = []QianchuanToolsEstimateAudienceV10AudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewQianchuanToolsEstimateAudienceV10AudienceAcFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceAcFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceAc, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceAc: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceAc) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_ac value
func (v QianchuanToolsEstimateAudienceV10AudienceAc) Ptr() *QianchuanToolsEstimateAudienceV10AudienceAc {
	return &v
}
