/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdRegionUpdateV10District
type QianchuanAdRegionUpdateV10District string

// List of qianchuan_ad_region_update_v1.0_district
const (
	CITY_QianchuanAdRegionUpdateV10District   QianchuanAdRegionUpdateV10District = "CITY"
	COUNTY_QianchuanAdRegionUpdateV10District QianchuanAdRegionUpdateV10District = "COUNTY"
	NONE_QianchuanAdRegionUpdateV10District   QianchuanAdRegionUpdateV10District = "NONE"
)

// All allowed values of QianchuanAdRegionUpdateV10District enum
var AllowedQianchuanAdRegionUpdateV10DistrictEnumValues = []QianchuanAdRegionUpdateV10District{
	"CITY",
	"COUNTY",
	"NONE",
}

// NewQianchuanAdRegionUpdateV10DistrictFromValue returns a pointer to a valid QianchuanAdRegionUpdateV10District
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRegionUpdateV10DistrictFromValue(v string) (*QianchuanAdRegionUpdateV10District, error) {
	ev := QianchuanAdRegionUpdateV10District(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRegionUpdateV10District: valid values are %v", v, AllowedQianchuanAdRegionUpdateV10DistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRegionUpdateV10District) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRegionUpdateV10DistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_region_update_v1.0_district value
func (v QianchuanAdRegionUpdateV10District) Ptr() *QianchuanAdRegionUpdateV10District {
	return &v
}
