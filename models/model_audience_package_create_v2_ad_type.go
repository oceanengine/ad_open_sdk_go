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

// AudiencePackageCreateV2AdType
type AudiencePackageCreateV2AdType string

// List of audience_package_create_v2_ad_type
const (
	SEARCH_AudiencePackageCreateV2AdType AudiencePackageCreateV2AdType = "SEARCH"
	ALL_AudiencePackageCreateV2AdType    AudiencePackageCreateV2AdType = "ALL"
)

// All allowed values of AudiencePackageCreateV2AdType enum
var AllowedAudiencePackageCreateV2AdTypeEnumValues = []AudiencePackageCreateV2AdType{
	"SEARCH",
	"ALL",
}

// NewAudiencePackageCreateV2AdTypeFromValue returns a pointer to a valid AudiencePackageCreateV2AdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2AdTypeFromValue(v string) (*AudiencePackageCreateV2AdType, error) {
	ev := AudiencePackageCreateV2AdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2AdType: valid values are %v", v, AllowedAudiencePackageCreateV2AdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2AdType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2AdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_ad_type value
func (v AudiencePackageCreateV2AdType) Ptr() *AudiencePackageCreateV2AdType {
	return &v
}
