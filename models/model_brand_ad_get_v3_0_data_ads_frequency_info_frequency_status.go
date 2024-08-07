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

// BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus
type BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus int64

// List of brand_ad_get_v3.0_data_ads_frequency_info_frequency_status
const (
	Enum_0_BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus = 0
	Enum_1_BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus = 1
)

// All allowed values of BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus enum
var AllowedBrandAdGetV30DataAdsFrequencyInfoFrequencyStatusEnumValues = []BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus{
	0,
	1,
}

// NewBrandAdGetV30DataAdsFrequencyInfoFrequencyStatusFromValue returns a pointer to a valid BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsFrequencyInfoFrequencyStatusFromValue(v int64) (*BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus, error) {
	ev := BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus: valid values are %v", v, AllowedBrandAdGetV30DataAdsFrequencyInfoFrequencyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsFrequencyInfoFrequencyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_frequency_info_frequency_status value
func (v BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus) Ptr() *BrandAdGetV30DataAdsFrequencyInfoFrequencyStatus {
	return &v
}
