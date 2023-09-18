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

// QianchuanToolsEstimateAudienceV10AudienceAudienceMode
type QianchuanToolsEstimateAudienceV10AudienceAudienceMode string

// List of qianchuan_tools_estimate_audience_v1.0_audience_audience_mode
const (
	AUTO_QianchuanToolsEstimateAudienceV10AudienceAudienceMode                QianchuanToolsEstimateAudienceV10AudienceAudienceMode = "AUTO"
	AUTO_ORIENTATION_QianchuanToolsEstimateAudienceV10AudienceAudienceMode    QianchuanToolsEstimateAudienceV10AudienceAudienceMode = "AUTO_ORIENTATION"
	CUSTOM_QianchuanToolsEstimateAudienceV10AudienceAudienceMode              QianchuanToolsEstimateAudienceV10AudienceAudienceMode = "CUSTOM"
	NONE_QianchuanToolsEstimateAudienceV10AudienceAudienceMode                QianchuanToolsEstimateAudienceV10AudienceAudienceMode = "NONE"
	ORIENTATION_PACKAGE_QianchuanToolsEstimateAudienceV10AudienceAudienceMode QianchuanToolsEstimateAudienceV10AudienceAudienceMode = "ORIENTATION_PACKAGE"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceAudienceMode enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceAudienceModeEnumValues = []QianchuanToolsEstimateAudienceV10AudienceAudienceMode{
	"AUTO",
	"AUTO_ORIENTATION",
	"CUSTOM",
	"NONE",
	"ORIENTATION_PACKAGE",
}

// NewQianchuanToolsEstimateAudienceV10AudienceAudienceModeFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceAudienceModeFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceAudienceMode, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceAudienceMode: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_audience_mode value
func (v QianchuanToolsEstimateAudienceV10AudienceAudienceMode) Ptr() *QianchuanToolsEstimateAudienceV10AudienceAudienceMode {
	return &v
}