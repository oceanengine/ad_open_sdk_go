/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceAudienceMode
type QianchuanAdCreateV10AudienceAudienceMode string

// List of qianchuan_ad_create_v1.0_audience_audience_mode
const (
	AUTO_QianchuanAdCreateV10AudienceAudienceMode                QianchuanAdCreateV10AudienceAudienceMode = "AUTO"
	AUTO_ORIENTATION_QianchuanAdCreateV10AudienceAudienceMode    QianchuanAdCreateV10AudienceAudienceMode = "AUTO_ORIENTATION"
	CUSTOM_QianchuanAdCreateV10AudienceAudienceMode              QianchuanAdCreateV10AudienceAudienceMode = "CUSTOM"
	NONE_QianchuanAdCreateV10AudienceAudienceMode                QianchuanAdCreateV10AudienceAudienceMode = "NONE"
	ORIENTATION_PACKAGE_QianchuanAdCreateV10AudienceAudienceMode QianchuanAdCreateV10AudienceAudienceMode = "ORIENTATION_PACKAGE"
)

// All allowed values of QianchuanAdCreateV10AudienceAudienceMode enum
var AllowedQianchuanAdCreateV10AudienceAudienceModeEnumValues = []QianchuanAdCreateV10AudienceAudienceMode{
	"AUTO",
	"AUTO_ORIENTATION",
	"CUSTOM",
	"NONE",
	"ORIENTATION_PACKAGE",
}

// NewQianchuanAdCreateV10AudienceAudienceModeFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceAudienceModeFromValue(v string) (*QianchuanAdCreateV10AudienceAudienceMode, error) {
	ev := QianchuanAdCreateV10AudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceAudienceMode: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_audience_mode value
func (v QianchuanAdCreateV10AudienceAudienceMode) Ptr() *QianchuanAdCreateV10AudienceAudienceMode {
	return &v
}
