/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoAudienceDistrict
type AdlabGroupCreateV30AdInfoAudienceDistrict string

// List of adlab_group_create_v3.0_ad_info_audience_district
const (
	BUSINESS_DISTRICT_AdlabGroupCreateV30AdInfoAudienceDistrict AdlabGroupCreateV30AdInfoAudienceDistrict = "BUSINESS_DISTRICT"
	NONE_AdlabGroupCreateV30AdInfoAudienceDistrict              AdlabGroupCreateV30AdInfoAudienceDistrict = "NONE"
	REGION_AdlabGroupCreateV30AdInfoAudienceDistrict            AdlabGroupCreateV30AdInfoAudienceDistrict = "REGION"
)

// All allowed values of AdlabGroupCreateV30AdInfoAudienceDistrict enum
var AllowedAdlabGroupCreateV30AdInfoAudienceDistrictEnumValues = []AdlabGroupCreateV30AdInfoAudienceDistrict{
	"BUSINESS_DISTRICT",
	"NONE",
	"REGION",
}

// NewAdlabGroupCreateV30AdInfoAudienceDistrictFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAudienceDistrictFromValue(v string) (*AdlabGroupCreateV30AdInfoAudienceDistrict, error) {
	ev := AdlabGroupCreateV30AdInfoAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAudienceDistrict: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_audience_district value
func (v AdlabGroupCreateV30AdInfoAudienceDistrict) Ptr() *AdlabGroupCreateV30AdInfoAudienceDistrict {
	return &v
}