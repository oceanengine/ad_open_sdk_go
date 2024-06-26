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

// AdlabGroupListV30DataGroupsAdInfoAudienceDistrict
type AdlabGroupListV30DataGroupsAdInfoAudienceDistrict string

// List of adlab_group_list_v3.0_data_groups_ad_info_audience_district
const (
	BUSINESS_DISTRICT_AdlabGroupListV30DataGroupsAdInfoAudienceDistrict AdlabGroupListV30DataGroupsAdInfoAudienceDistrict = "BUSINESS_DISTRICT"
	CITY_AdlabGroupListV30DataGroupsAdInfoAudienceDistrict              AdlabGroupListV30DataGroupsAdInfoAudienceDistrict = "CITY"
	COUNTY_AdlabGroupListV30DataGroupsAdInfoAudienceDistrict            AdlabGroupListV30DataGroupsAdInfoAudienceDistrict = "COUNTY"
	NONE_AdlabGroupListV30DataGroupsAdInfoAudienceDistrict              AdlabGroupListV30DataGroupsAdInfoAudienceDistrict = "NONE"
	REGION_AdlabGroupListV30DataGroupsAdInfoAudienceDistrict            AdlabGroupListV30DataGroupsAdInfoAudienceDistrict = "REGION"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAudienceDistrict enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAudienceDistrictEnumValues = []AdlabGroupListV30DataGroupsAdInfoAudienceDistrict{
	"BUSINESS_DISTRICT",
	"CITY",
	"COUNTY",
	"NONE",
	"REGION",
}

// NewAdlabGroupListV30DataGroupsAdInfoAudienceDistrictFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAudienceDistrictFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAudienceDistrict, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAudienceDistrict: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_audience_district value
func (v AdlabGroupListV30DataGroupsAdInfoAudienceDistrict) Ptr() *AdlabGroupListV30DataGroupsAdInfoAudienceDistrict {
	return &v
}
