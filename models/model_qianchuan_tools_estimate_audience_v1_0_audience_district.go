/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceDistrict
type QianchuanToolsEstimateAudienceV10AudienceDistrict string

// List of qianchuan_tools_estimate_audience_v1.0_audience_district
const (
	CITY_QianchuanToolsEstimateAudienceV10AudienceDistrict   QianchuanToolsEstimateAudienceV10AudienceDistrict = "CITY"
	COUNTY_QianchuanToolsEstimateAudienceV10AudienceDistrict QianchuanToolsEstimateAudienceV10AudienceDistrict = "COUNTY"
	NONE_QianchuanToolsEstimateAudienceV10AudienceDistrict   QianchuanToolsEstimateAudienceV10AudienceDistrict = "NONE"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceDistrict enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceDistrictEnumValues = []QianchuanToolsEstimateAudienceV10AudienceDistrict{
	"CITY",
	"COUNTY",
	"NONE",
}

// NewQianchuanToolsEstimateAudienceV10AudienceDistrictFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceDistrictFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceDistrict, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceDistrict: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_district value
func (v QianchuanToolsEstimateAudienceV10AudienceDistrict) Ptr() *QianchuanToolsEstimateAudienceV10AudienceDistrict {
	return &v
}
