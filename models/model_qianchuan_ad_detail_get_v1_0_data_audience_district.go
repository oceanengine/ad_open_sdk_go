/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataAudienceDistrict
type QianchuanAdDetailGetV10DataAudienceDistrict string

// List of qianchuan_ad_detail_get_v1.0_data_audience_district
const (
	BUSINESS_DISTRICT_QianchuanAdDetailGetV10DataAudienceDistrict QianchuanAdDetailGetV10DataAudienceDistrict = "BUSINESS_DISTRICT"
	CITY_QianchuanAdDetailGetV10DataAudienceDistrict              QianchuanAdDetailGetV10DataAudienceDistrict = "CITY"
	COUNTY_QianchuanAdDetailGetV10DataAudienceDistrict            QianchuanAdDetailGetV10DataAudienceDistrict = "COUNTY"
	NONE_QianchuanAdDetailGetV10DataAudienceDistrict              QianchuanAdDetailGetV10DataAudienceDistrict = "NONE"
	OVERSEA_QianchuanAdDetailGetV10DataAudienceDistrict           QianchuanAdDetailGetV10DataAudienceDistrict = "OVERSEA"
	REGION_QianchuanAdDetailGetV10DataAudienceDistrict            QianchuanAdDetailGetV10DataAudienceDistrict = "REGION"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceDistrict enum
var AllowedQianchuanAdDetailGetV10DataAudienceDistrictEnumValues = []QianchuanAdDetailGetV10DataAudienceDistrict{
	"BUSINESS_DISTRICT",
	"CITY",
	"COUNTY",
	"NONE",
	"OVERSEA",
	"REGION",
}

// NewQianchuanAdDetailGetV10DataAudienceDistrictFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceDistrictFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceDistrict, error) {
	ev := QianchuanAdDetailGetV10DataAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceDistrict: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_district value
func (v QianchuanAdDetailGetV10DataAudienceDistrict) Ptr() *QianchuanAdDetailGetV10DataAudienceDistrict {
	return &v
}
