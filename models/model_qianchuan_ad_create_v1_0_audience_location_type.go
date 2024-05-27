/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceLocationType
type QianchuanAdCreateV10AudienceLocationType string

// List of qianchuan_ad_create_v1.0_audience_location_type
const (
	ALL_QianchuanAdCreateV10AudienceLocationType     QianchuanAdCreateV10AudienceLocationType = "ALL"
	CURRENT_QianchuanAdCreateV10AudienceLocationType QianchuanAdCreateV10AudienceLocationType = "CURRENT"
	HOME_QianchuanAdCreateV10AudienceLocationType    QianchuanAdCreateV10AudienceLocationType = "HOME"
	TRAVEL_QianchuanAdCreateV10AudienceLocationType  QianchuanAdCreateV10AudienceLocationType = "TRAVEL"
)

// All allowed values of QianchuanAdCreateV10AudienceLocationType enum
var AllowedQianchuanAdCreateV10AudienceLocationTypeEnumValues = []QianchuanAdCreateV10AudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewQianchuanAdCreateV10AudienceLocationTypeFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceLocationTypeFromValue(v string) (*QianchuanAdCreateV10AudienceLocationType, error) {
	ev := QianchuanAdCreateV10AudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceLocationType: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceLocationType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_location_type value
func (v QianchuanAdCreateV10AudienceLocationType) Ptr() *QianchuanAdCreateV10AudienceLocationType {
	return &v
}
