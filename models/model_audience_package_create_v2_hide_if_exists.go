/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2HideIfExists
type AudiencePackageCreateV2HideIfExists int64

// List of audience_package_create_v2_hide_if_exists
const (
	Enum_0_AudiencePackageCreateV2HideIfExists AudiencePackageCreateV2HideIfExists = 0
	Enum_1_AudiencePackageCreateV2HideIfExists AudiencePackageCreateV2HideIfExists = 1
	Enum_2_AudiencePackageCreateV2HideIfExists AudiencePackageCreateV2HideIfExists = 2
)

// All allowed values of AudiencePackageCreateV2HideIfExists enum
var AllowedAudiencePackageCreateV2HideIfExistsEnumValues = []AudiencePackageCreateV2HideIfExists{
	0,
	1,
	2,
}

// NewAudiencePackageCreateV2HideIfExistsFromValue returns a pointer to a valid AudiencePackageCreateV2HideIfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2HideIfExistsFromValue(v int64) (*AudiencePackageCreateV2HideIfExists, error) {
	ev := AudiencePackageCreateV2HideIfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2HideIfExists: valid values are %v", v, AllowedAudiencePackageCreateV2HideIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2HideIfExists) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2HideIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_hide_if_exists value
func (v AudiencePackageCreateV2HideIfExists) Ptr() *AudiencePackageCreateV2HideIfExists {
	return &v
}
