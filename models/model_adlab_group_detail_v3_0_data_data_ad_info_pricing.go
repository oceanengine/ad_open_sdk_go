/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoPricing
type AdlabGroupDetailV30DataDataAdInfoPricing string

// List of adlab_group_detail_v3.0_data_data_ad_info_pricing
const (
	PRICING_OCPM_AdlabGroupDetailV30DataDataAdInfoPricing AdlabGroupDetailV30DataDataAdInfoPricing = "PRICING_OCPM"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoPricing enum
var AllowedAdlabGroupDetailV30DataDataAdInfoPricingEnumValues = []AdlabGroupDetailV30DataDataAdInfoPricing{
	"PRICING_OCPM",
}

// NewAdlabGroupDetailV30DataDataAdInfoPricingFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoPricingFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoPricing, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoPricing: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoPricing) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_pricing value
func (v AdlabGroupDetailV30DataDataAdInfoPricing) Ptr() *AdlabGroupDetailV30DataDataAdInfoPricing {
	return &v
}
