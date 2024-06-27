/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdGetV30DataAdsAudienceInfoDistrictSplit
type BrandAdGetV30DataAdsAudienceInfoDistrictSplit int64

// List of brand_ad_get_v3.0_data_ads_audience_info_district_split
const (
	Enum_0_BrandAdGetV30DataAdsAudienceInfoDistrictSplit BrandAdGetV30DataAdsAudienceInfoDistrictSplit = 0
	Enum_1_BrandAdGetV30DataAdsAudienceInfoDistrictSplit BrandAdGetV30DataAdsAudienceInfoDistrictSplit = 1
)

// All allowed values of BrandAdGetV30DataAdsAudienceInfoDistrictSplit enum
var AllowedBrandAdGetV30DataAdsAudienceInfoDistrictSplitEnumValues = []BrandAdGetV30DataAdsAudienceInfoDistrictSplit{
	0,
	1,
}

// NewBrandAdGetV30DataAdsAudienceInfoDistrictSplitFromValue returns a pointer to a valid BrandAdGetV30DataAdsAudienceInfoDistrictSplit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAudienceInfoDistrictSplitFromValue(v int64) (*BrandAdGetV30DataAdsAudienceInfoDistrictSplit, error) {
	ev := BrandAdGetV30DataAdsAudienceInfoDistrictSplit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAudienceInfoDistrictSplit: valid values are %v", v, AllowedBrandAdGetV30DataAdsAudienceInfoDistrictSplitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAudienceInfoDistrictSplit) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAudienceInfoDistrictSplitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_audience_info_district_split value
func (v BrandAdGetV30DataAdsAudienceInfoDistrictSplit) Ptr() *BrandAdGetV30DataAdsAudienceInfoDistrictSplit {
	return &v
}
