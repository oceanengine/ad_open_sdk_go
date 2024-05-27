/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageUpdateV2District
type AudiencePackageUpdateV2District string

// List of audience_package_update_v2_district
const (
	NONE_AudiencePackageUpdateV2District              AudiencePackageUpdateV2District = "NONE"
	OVERSEA_AudiencePackageUpdateV2District           AudiencePackageUpdateV2District = "OVERSEA"
	BUSINESS_DISTRICT_AudiencePackageUpdateV2District AudiencePackageUpdateV2District = "BUSINESS_DISTRICT"
	REGION_AudiencePackageUpdateV2District            AudiencePackageUpdateV2District = "REGION"
)

// All allowed values of AudiencePackageUpdateV2District enum
var AllowedAudiencePackageUpdateV2DistrictEnumValues = []AudiencePackageUpdateV2District{
	"NONE",
	"OVERSEA",
	"BUSINESS_DISTRICT",
	"REGION",
}

// NewAudiencePackageUpdateV2DistrictFromValue returns a pointer to a valid AudiencePackageUpdateV2District
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2DistrictFromValue(v string) (*AudiencePackageUpdateV2District, error) {
	ev := AudiencePackageUpdateV2District(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2District: valid values are %v", v, AllowedAudiencePackageUpdateV2DistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2District) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2DistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_district value
func (v AudiencePackageUpdateV2District) Ptr() *AudiencePackageUpdateV2District {
	return &v
}
