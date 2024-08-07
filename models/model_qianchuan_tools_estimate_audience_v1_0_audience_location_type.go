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

// QianchuanToolsEstimateAudienceV10AudienceLocationType
type QianchuanToolsEstimateAudienceV10AudienceLocationType string

// List of qianchuan_tools_estimate_audience_v1.0_audience_location_type
const (
	ALL_QianchuanToolsEstimateAudienceV10AudienceLocationType     QianchuanToolsEstimateAudienceV10AudienceLocationType = "ALL"
	CURRENT_QianchuanToolsEstimateAudienceV10AudienceLocationType QianchuanToolsEstimateAudienceV10AudienceLocationType = "CURRENT"
	HOME_QianchuanToolsEstimateAudienceV10AudienceLocationType    QianchuanToolsEstimateAudienceV10AudienceLocationType = "HOME"
	TRAVEL_QianchuanToolsEstimateAudienceV10AudienceLocationType  QianchuanToolsEstimateAudienceV10AudienceLocationType = "TRAVEL"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceLocationType enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceLocationTypeEnumValues = []QianchuanToolsEstimateAudienceV10AudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewQianchuanToolsEstimateAudienceV10AudienceLocationTypeFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceLocationTypeFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceLocationType, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceLocationType: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceLocationType) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_location_type value
func (v QianchuanToolsEstimateAudienceV10AudienceLocationType) Ptr() *QianchuanToolsEstimateAudienceV10AudienceLocationType {
	return &v
}
