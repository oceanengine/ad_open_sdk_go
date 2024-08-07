/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataAudienceLocationType
type QianchuanAdDetailGetV10DataAudienceLocationType string

// List of qianchuan_ad_detail_get_v1.0_data_audience_location_type
const (
	ALL_QianchuanAdDetailGetV10DataAudienceLocationType     QianchuanAdDetailGetV10DataAudienceLocationType = "ALL"
	CURRENT_QianchuanAdDetailGetV10DataAudienceLocationType QianchuanAdDetailGetV10DataAudienceLocationType = "CURRENT"
	HOME_QianchuanAdDetailGetV10DataAudienceLocationType    QianchuanAdDetailGetV10DataAudienceLocationType = "HOME"
	TRAVEL_QianchuanAdDetailGetV10DataAudienceLocationType  QianchuanAdDetailGetV10DataAudienceLocationType = "TRAVEL"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceLocationType enum
var AllowedQianchuanAdDetailGetV10DataAudienceLocationTypeEnumValues = []QianchuanAdDetailGetV10DataAudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewQianchuanAdDetailGetV10DataAudienceLocationTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceLocationTypeFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceLocationType, error) {
	ev := QianchuanAdDetailGetV10DataAudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceLocationType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceLocationType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_location_type value
func (v QianchuanAdDetailGetV10DataAudienceLocationType) Ptr() *QianchuanAdDetailGetV10DataAudienceLocationType {
	return &v
}
