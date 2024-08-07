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

// BrandAdGetV30DataAdsAudienceInfoDistrictType
type BrandAdGetV30DataAdsAudienceInfoDistrictType int64

// List of brand_ad_get_v3.0_data_ads_audience_info_district_type
const (
	Enum_0_BrandAdGetV30DataAdsAudienceInfoDistrictType BrandAdGetV30DataAdsAudienceInfoDistrictType = 0
	Enum_1_BrandAdGetV30DataAdsAudienceInfoDistrictType BrandAdGetV30DataAdsAudienceInfoDistrictType = 1
	Enum_2_BrandAdGetV30DataAdsAudienceInfoDistrictType BrandAdGetV30DataAdsAudienceInfoDistrictType = 2
	Enum_3_BrandAdGetV30DataAdsAudienceInfoDistrictType BrandAdGetV30DataAdsAudienceInfoDistrictType = 3
)

// All allowed values of BrandAdGetV30DataAdsAudienceInfoDistrictType enum
var AllowedBrandAdGetV30DataAdsAudienceInfoDistrictTypeEnumValues = []BrandAdGetV30DataAdsAudienceInfoDistrictType{
	0,
	1,
	2,
	3,
}

// NewBrandAdGetV30DataAdsAudienceInfoDistrictTypeFromValue returns a pointer to a valid BrandAdGetV30DataAdsAudienceInfoDistrictType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAudienceInfoDistrictTypeFromValue(v int64) (*BrandAdGetV30DataAdsAudienceInfoDistrictType, error) {
	ev := BrandAdGetV30DataAdsAudienceInfoDistrictType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAudienceInfoDistrictType: valid values are %v", v, AllowedBrandAdGetV30DataAdsAudienceInfoDistrictTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAudienceInfoDistrictType) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAudienceInfoDistrictTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_audience_info_district_type value
func (v BrandAdGetV30DataAdsAudienceInfoDistrictType) Ptr() *BrandAdGetV30DataAdsAudienceInfoDistrictType {
	return &v
}
