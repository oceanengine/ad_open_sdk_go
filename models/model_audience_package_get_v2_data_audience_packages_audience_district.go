/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesAudienceDistrict
type AudiencePackageGetV2DataAudiencePackagesAudienceDistrict string

// List of audience_package_get_v2_data_audience_packages_audience_district
const (
	BUSINESS_DISTRICT_AudiencePackageGetV2DataAudiencePackagesAudienceDistrict AudiencePackageGetV2DataAudiencePackagesAudienceDistrict = "BUSINESS_DISTRICT"
	CITY_AudiencePackageGetV2DataAudiencePackagesAudienceDistrict              AudiencePackageGetV2DataAudiencePackagesAudienceDistrict = "CITY"
	COUNTY_AudiencePackageGetV2DataAudiencePackagesAudienceDistrict            AudiencePackageGetV2DataAudiencePackagesAudienceDistrict = "COUNTY"
	NONE_AudiencePackageGetV2DataAudiencePackagesAudienceDistrict              AudiencePackageGetV2DataAudiencePackagesAudienceDistrict = "NONE"
	OVERSEA_AudiencePackageGetV2DataAudiencePackagesAudienceDistrict           AudiencePackageGetV2DataAudiencePackagesAudienceDistrict = "OVERSEA"
	REGION_AudiencePackageGetV2DataAudiencePackagesAudienceDistrict            AudiencePackageGetV2DataAudiencePackagesAudienceDistrict = "REGION"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceDistrict enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDistrictEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceDistrict{
	"BUSINESS_DISTRICT",
	"CITY",
	"COUNTY",
	"NONE",
	"OVERSEA",
	"REGION",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceDistrictFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceDistrictFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceDistrict, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceDistrict: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_district value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceDistrict) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceDistrict {
	return &v
}