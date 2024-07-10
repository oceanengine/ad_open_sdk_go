/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdGetV30DataAdsClassify
type BrandAdGetV30DataAdsClassify int64

// List of brand_ad_get_v3.0_data_ads_classify
const (
	Enum_1_BrandAdGetV30DataAdsClassify BrandAdGetV30DataAdsClassify = 1
	Enum_2_BrandAdGetV30DataAdsClassify BrandAdGetV30DataAdsClassify = 2
	Enum_3_BrandAdGetV30DataAdsClassify BrandAdGetV30DataAdsClassify = 3
	Enum_4_BrandAdGetV30DataAdsClassify BrandAdGetV30DataAdsClassify = 4
	Enum_5_BrandAdGetV30DataAdsClassify BrandAdGetV30DataAdsClassify = 5
	Enum_6_BrandAdGetV30DataAdsClassify BrandAdGetV30DataAdsClassify = 6
	Enum_7_BrandAdGetV30DataAdsClassify BrandAdGetV30DataAdsClassify = 7
)

// All allowed values of BrandAdGetV30DataAdsClassify enum
var AllowedBrandAdGetV30DataAdsClassifyEnumValues = []BrandAdGetV30DataAdsClassify{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
}

// NewBrandAdGetV30DataAdsClassifyFromValue returns a pointer to a valid BrandAdGetV30DataAdsClassify
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsClassifyFromValue(v int64) (*BrandAdGetV30DataAdsClassify, error) {
	ev := BrandAdGetV30DataAdsClassify(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsClassify: valid values are %v", v, AllowedBrandAdGetV30DataAdsClassifyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsClassify) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsClassifyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_classify value
func (v BrandAdGetV30DataAdsClassify) Ptr() *BrandAdGetV30DataAdsClassify {
	return &v
}
