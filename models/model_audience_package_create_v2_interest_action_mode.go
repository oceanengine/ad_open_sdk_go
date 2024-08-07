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

// AudiencePackageCreateV2InterestActionMode
type AudiencePackageCreateV2InterestActionMode string

// List of audience_package_create_v2_interest_action_mode
const (
	CUSTOM_AudiencePackageCreateV2InterestActionMode    AudiencePackageCreateV2InterestActionMode = "CUSTOM"
	UNLIMITED_AudiencePackageCreateV2InterestActionMode AudiencePackageCreateV2InterestActionMode = "UNLIMITED"
	RECOMMEND_AudiencePackageCreateV2InterestActionMode AudiencePackageCreateV2InterestActionMode = "RECOMMEND"
)

// All allowed values of AudiencePackageCreateV2InterestActionMode enum
var AllowedAudiencePackageCreateV2InterestActionModeEnumValues = []AudiencePackageCreateV2InterestActionMode{
	"CUSTOM",
	"UNLIMITED",
	"RECOMMEND",
}

// NewAudiencePackageCreateV2InterestActionModeFromValue returns a pointer to a valid AudiencePackageCreateV2InterestActionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2InterestActionModeFromValue(v string) (*AudiencePackageCreateV2InterestActionMode, error) {
	ev := AudiencePackageCreateV2InterestActionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2InterestActionMode: valid values are %v", v, AllowedAudiencePackageCreateV2InterestActionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2InterestActionMode) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2InterestActionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_interest_action_mode value
func (v AudiencePackageCreateV2InterestActionMode) Ptr() *AudiencePackageCreateV2InterestActionMode {
	return &v
}
