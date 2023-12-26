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

// QianchuanAwemeOrderCreateV10AudienceAudienceMode
type QianchuanAwemeOrderCreateV10AudienceAudienceMode string

// List of qianchuan_aweme_order_create_v1.0_audience_audience_mode
const (
	ATUO_QianchuanAwemeOrderCreateV10AudienceAudienceMode     QianchuanAwemeOrderCreateV10AudienceAudienceMode = "ATUO"
	CUSTOM_QianchuanAwemeOrderCreateV10AudienceAudienceMode   QianchuanAwemeOrderCreateV10AudienceAudienceMode = "CUSTOM"
	FANS_QianchuanAwemeOrderCreateV10AudienceAudienceMode     QianchuanAwemeOrderCreateV10AudienceAudienceMode = "FANS"
	LIVEFANS_QianchuanAwemeOrderCreateV10AudienceAudienceMode QianchuanAwemeOrderCreateV10AudienceAudienceMode = "LIVEFANS"
)

// All allowed values of QianchuanAwemeOrderCreateV10AudienceAudienceMode enum
var AllowedQianchuanAwemeOrderCreateV10AudienceAudienceModeEnumValues = []QianchuanAwemeOrderCreateV10AudienceAudienceMode{
	"ATUO",
	"CUSTOM",
	"FANS",
	"LIVEFANS",
}

// NewQianchuanAwemeOrderCreateV10AudienceAudienceModeFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10AudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10AudienceAudienceModeFromValue(v string) (*QianchuanAwemeOrderCreateV10AudienceAudienceMode, error) {
	ev := QianchuanAwemeOrderCreateV10AudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10AudienceAudienceMode: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10AudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10AudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10AudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_audience_audience_mode value
func (v QianchuanAwemeOrderCreateV10AudienceAudienceMode) Ptr() *QianchuanAwemeOrderCreateV10AudienceAudienceMode {
	return &v
}
