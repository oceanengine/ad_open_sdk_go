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

// AudiencePackageGetV2FilteringAdType
type AudiencePackageGetV2FilteringAdType string

// List of audience_package_get_v2_filtering_ad_type
const (
	ALL_AudiencePackageGetV2FilteringAdType    AudiencePackageGetV2FilteringAdType = "ALL"
	SEARCH_AudiencePackageGetV2FilteringAdType AudiencePackageGetV2FilteringAdType = "SEARCH"
)

// All allowed values of AudiencePackageGetV2FilteringAdType enum
var AllowedAudiencePackageGetV2FilteringAdTypeEnumValues = []AudiencePackageGetV2FilteringAdType{
	"ALL",
	"SEARCH",
}

// NewAudiencePackageGetV2FilteringAdTypeFromValue returns a pointer to a valid AudiencePackageGetV2FilteringAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2FilteringAdTypeFromValue(v string) (*AudiencePackageGetV2FilteringAdType, error) {
	ev := AudiencePackageGetV2FilteringAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2FilteringAdType: valid values are %v", v, AllowedAudiencePackageGetV2FilteringAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2FilteringAdType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2FilteringAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_filtering_ad_type value
func (v AudiencePackageGetV2FilteringAdType) Ptr() *AudiencePackageGetV2FilteringAdType {
	return &v
}
