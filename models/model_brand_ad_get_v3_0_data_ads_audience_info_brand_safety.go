/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdGetV30DataAdsAudienceInfoBrandSafety
type BrandAdGetV30DataAdsAudienceInfoBrandSafety int64

// List of brand_ad_get_v3.0_data_ads_audience_info_brand_safety
const (
	Enum_0_BrandAdGetV30DataAdsAudienceInfoBrandSafety BrandAdGetV30DataAdsAudienceInfoBrandSafety = 0
	Enum_1_BrandAdGetV30DataAdsAudienceInfoBrandSafety BrandAdGetV30DataAdsAudienceInfoBrandSafety = 1
	Enum_2_BrandAdGetV30DataAdsAudienceInfoBrandSafety BrandAdGetV30DataAdsAudienceInfoBrandSafety = 2
	Enum_3_BrandAdGetV30DataAdsAudienceInfoBrandSafety BrandAdGetV30DataAdsAudienceInfoBrandSafety = 3
	Enum_4_BrandAdGetV30DataAdsAudienceInfoBrandSafety BrandAdGetV30DataAdsAudienceInfoBrandSafety = 4
)

// All allowed values of BrandAdGetV30DataAdsAudienceInfoBrandSafety enum
var AllowedBrandAdGetV30DataAdsAudienceInfoBrandSafetyEnumValues = []BrandAdGetV30DataAdsAudienceInfoBrandSafety{
	0,
	1,
	2,
	3,
	4,
}

// NewBrandAdGetV30DataAdsAudienceInfoBrandSafetyFromValue returns a pointer to a valid BrandAdGetV30DataAdsAudienceInfoBrandSafety
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAudienceInfoBrandSafetyFromValue(v int64) (*BrandAdGetV30DataAdsAudienceInfoBrandSafety, error) {
	ev := BrandAdGetV30DataAdsAudienceInfoBrandSafety(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAudienceInfoBrandSafety: valid values are %v", v, AllowedBrandAdGetV30DataAdsAudienceInfoBrandSafetyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAudienceInfoBrandSafety) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAudienceInfoBrandSafetyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_audience_info_brand_safety value
func (v BrandAdGetV30DataAdsAudienceInfoBrandSafety) Ptr() *BrandAdGetV30DataAdsAudienceInfoBrandSafety {
	return &v
}
