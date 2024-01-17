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

// BrandAdGetV30DataAdsAudienceInfoAcV2
type BrandAdGetV30DataAdsAudienceInfoAcV2 int64

// List of brand_ad_get_v3.0_data_ads_audience_info_ac_v2
const (
	Enum_1_BrandAdGetV30DataAdsAudienceInfoAcV2 BrandAdGetV30DataAdsAudienceInfoAcV2 = 1
	Enum_2_BrandAdGetV30DataAdsAudienceInfoAcV2 BrandAdGetV30DataAdsAudienceInfoAcV2 = 2
	Enum_3_BrandAdGetV30DataAdsAudienceInfoAcV2 BrandAdGetV30DataAdsAudienceInfoAcV2 = 3
	Enum_4_BrandAdGetV30DataAdsAudienceInfoAcV2 BrandAdGetV30DataAdsAudienceInfoAcV2 = 4
)

// All allowed values of BrandAdGetV30DataAdsAudienceInfoAcV2 enum
var AllowedBrandAdGetV30DataAdsAudienceInfoAcV2EnumValues = []BrandAdGetV30DataAdsAudienceInfoAcV2{
	1,
	2,
	3,
	4,
}

// NewBrandAdGetV30DataAdsAudienceInfoAcV2FromValue returns a pointer to a valid BrandAdGetV30DataAdsAudienceInfoAcV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAudienceInfoAcV2FromValue(v int64) (*BrandAdGetV30DataAdsAudienceInfoAcV2, error) {
	ev := BrandAdGetV30DataAdsAudienceInfoAcV2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAudienceInfoAcV2: valid values are %v", v, AllowedBrandAdGetV30DataAdsAudienceInfoAcV2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAudienceInfoAcV2) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAudienceInfoAcV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_audience_info_ac_v2 value
func (v BrandAdGetV30DataAdsAudienceInfoAcV2) Ptr() *BrandAdGetV30DataAdsAudienceInfoAcV2 {
	return &v
}
