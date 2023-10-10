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

// QianchuanAdDetailGetV10DataAudienceAudienceMode
type QianchuanAdDetailGetV10DataAudienceAudienceMode string

// List of qianchuan_ad_detail_get_v1.0_data_audience_audience_mode
const (
	AUTO_QianchuanAdDetailGetV10DataAudienceAudienceMode                QianchuanAdDetailGetV10DataAudienceAudienceMode = "AUTO"
	AUTO_ORIENTATION_QianchuanAdDetailGetV10DataAudienceAudienceMode    QianchuanAdDetailGetV10DataAudienceAudienceMode = "AUTO_ORIENTATION"
	CUSTOM_QianchuanAdDetailGetV10DataAudienceAudienceMode              QianchuanAdDetailGetV10DataAudienceAudienceMode = "CUSTOM"
	NONE_QianchuanAdDetailGetV10DataAudienceAudienceMode                QianchuanAdDetailGetV10DataAudienceAudienceMode = "NONE"
	ORIENTATION_PACKAGE_QianchuanAdDetailGetV10DataAudienceAudienceMode QianchuanAdDetailGetV10DataAudienceAudienceMode = "ORIENTATION_PACKAGE"
	UNKNOWN_QianchuanAdDetailGetV10DataAudienceAudienceMode             QianchuanAdDetailGetV10DataAudienceAudienceMode = "UNKNOWN"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceAudienceMode enum
var AllowedQianchuanAdDetailGetV10DataAudienceAudienceModeEnumValues = []QianchuanAdDetailGetV10DataAudienceAudienceMode{
	"AUTO",
	"AUTO_ORIENTATION",
	"CUSTOM",
	"NONE",
	"ORIENTATION_PACKAGE",
	"UNKNOWN",
}

// NewQianchuanAdDetailGetV10DataAudienceAudienceModeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceAudienceModeFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceAudienceMode, error) {
	ev := QianchuanAdDetailGetV10DataAudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceAudienceMode: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_audience_mode value
func (v QianchuanAdDetailGetV10DataAudienceAudienceMode) Ptr() *QianchuanAdDetailGetV10DataAudienceAudienceMode {
	return &v
}
