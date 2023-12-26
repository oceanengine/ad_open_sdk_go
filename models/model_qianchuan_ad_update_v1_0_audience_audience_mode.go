/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceAudienceMode
type QianchuanAdUpdateV10AudienceAudienceMode string

// List of qianchuan_ad_update_v1.0_audience_audience_mode
const (
	AUTO_QianchuanAdUpdateV10AudienceAudienceMode                QianchuanAdUpdateV10AudienceAudienceMode = "AUTO"
	AUTO_ORIENTATION_QianchuanAdUpdateV10AudienceAudienceMode    QianchuanAdUpdateV10AudienceAudienceMode = "AUTO_ORIENTATION"
	CUSTOM_QianchuanAdUpdateV10AudienceAudienceMode              QianchuanAdUpdateV10AudienceAudienceMode = "CUSTOM"
	NONE_QianchuanAdUpdateV10AudienceAudienceMode                QianchuanAdUpdateV10AudienceAudienceMode = "NONE"
	ORIENTATION_PACKAGE_QianchuanAdUpdateV10AudienceAudienceMode QianchuanAdUpdateV10AudienceAudienceMode = "ORIENTATION_PACKAGE"
)

// All allowed values of QianchuanAdUpdateV10AudienceAudienceMode enum
var AllowedQianchuanAdUpdateV10AudienceAudienceModeEnumValues = []QianchuanAdUpdateV10AudienceAudienceMode{
	"AUTO",
	"AUTO_ORIENTATION",
	"CUSTOM",
	"NONE",
	"ORIENTATION_PACKAGE",
}

// NewQianchuanAdUpdateV10AudienceAudienceModeFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceAudienceModeFromValue(v string) (*QianchuanAdUpdateV10AudienceAudienceMode, error) {
	ev := QianchuanAdUpdateV10AudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceAudienceMode: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_audience_mode value
func (v QianchuanAdUpdateV10AudienceAudienceMode) Ptr() *QianchuanAdUpdateV10AudienceAudienceMode {
	return &v
}
