/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceLocationType
type QianchuanAdUpdateV10AudienceLocationType string

// List of qianchuan_ad_update_v1.0_audience_location_type
const (
	ALL_QianchuanAdUpdateV10AudienceLocationType     QianchuanAdUpdateV10AudienceLocationType = "ALL"
	CURRENT_QianchuanAdUpdateV10AudienceLocationType QianchuanAdUpdateV10AudienceLocationType = "CURRENT"
	HOME_QianchuanAdUpdateV10AudienceLocationType    QianchuanAdUpdateV10AudienceLocationType = "HOME"
	TRAVEL_QianchuanAdUpdateV10AudienceLocationType  QianchuanAdUpdateV10AudienceLocationType = "TRAVEL"
)

// All allowed values of QianchuanAdUpdateV10AudienceLocationType enum
var AllowedQianchuanAdUpdateV10AudienceLocationTypeEnumValues = []QianchuanAdUpdateV10AudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewQianchuanAdUpdateV10AudienceLocationTypeFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceLocationTypeFromValue(v string) (*QianchuanAdUpdateV10AudienceLocationType, error) {
	ev := QianchuanAdUpdateV10AudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceLocationType: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceLocationType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_location_type value
func (v QianchuanAdUpdateV10AudienceLocationType) Ptr() *QianchuanAdUpdateV10AudienceLocationType {
	return &v
}
