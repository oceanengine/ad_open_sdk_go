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

// AudiencePackageUpdateV2FilterOwnAwemeFans
type AudiencePackageUpdateV2FilterOwnAwemeFans int64

// List of audience_package_update_v2_filter_own_aweme_fans
const (
	Enum_0_AudiencePackageUpdateV2FilterOwnAwemeFans AudiencePackageUpdateV2FilterOwnAwemeFans = 0
	Enum_1_AudiencePackageUpdateV2FilterOwnAwemeFans AudiencePackageUpdateV2FilterOwnAwemeFans = 1
)

// All allowed values of AudiencePackageUpdateV2FilterOwnAwemeFans enum
var AllowedAudiencePackageUpdateV2FilterOwnAwemeFansEnumValues = []AudiencePackageUpdateV2FilterOwnAwemeFans{
	0,
	1,
}

// NewAudiencePackageUpdateV2FilterOwnAwemeFansFromValue returns a pointer to a valid AudiencePackageUpdateV2FilterOwnAwemeFans
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2FilterOwnAwemeFansFromValue(v int64) (*AudiencePackageUpdateV2FilterOwnAwemeFans, error) {
	ev := AudiencePackageUpdateV2FilterOwnAwemeFans(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2FilterOwnAwemeFans: valid values are %v", v, AllowedAudiencePackageUpdateV2FilterOwnAwemeFansEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2FilterOwnAwemeFans) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2FilterOwnAwemeFansEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_filter_own_aweme_fans value
func (v AudiencePackageUpdateV2FilterOwnAwemeFans) Ptr() *AudiencePackageUpdateV2FilterOwnAwemeFans {
	return &v
}
