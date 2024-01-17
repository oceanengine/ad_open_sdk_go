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

// AudiencePackageCreateV2District
type AudiencePackageCreateV2District string

// List of audience_package_create_v2_district
const (
	BUSINESS_DISTRICT_AudiencePackageCreateV2District AudiencePackageCreateV2District = "BUSINESS_DISTRICT"
	OVERSEA_AudiencePackageCreateV2District           AudiencePackageCreateV2District = "OVERSEA"
	REGION_AudiencePackageCreateV2District            AudiencePackageCreateV2District = "REGION"
	NONE_AudiencePackageCreateV2District              AudiencePackageCreateV2District = "NONE"
)

// All allowed values of AudiencePackageCreateV2District enum
var AllowedAudiencePackageCreateV2DistrictEnumValues = []AudiencePackageCreateV2District{
	"BUSINESS_DISTRICT",
	"OVERSEA",
	"REGION",
	"NONE",
}

// NewAudiencePackageCreateV2DistrictFromValue returns a pointer to a valid AudiencePackageCreateV2District
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2DistrictFromValue(v string) (*AudiencePackageCreateV2District, error) {
	ev := AudiencePackageCreateV2District(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2District: valid values are %v", v, AllowedAudiencePackageCreateV2DistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2District) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2DistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_district value
func (v AudiencePackageCreateV2District) Ptr() *AudiencePackageCreateV2District {
	return &v
}
