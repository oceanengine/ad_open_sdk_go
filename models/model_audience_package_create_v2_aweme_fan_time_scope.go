/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2AwemeFanTimeScope
type AudiencePackageCreateV2AwemeFanTimeScope string

// List of audience_package_create_v2_aweme_fan_time_scope
const (
	SIXTY_DAYS_AudiencePackageCreateV2AwemeFanTimeScope   AudiencePackageCreateV2AwemeFanTimeScope = "SIXTY_DAYS"
	THIRTY_DAYS_AudiencePackageCreateV2AwemeFanTimeScope  AudiencePackageCreateV2AwemeFanTimeScope = "THIRTY_DAYS"
	FIFTEEN_DAYS_AudiencePackageCreateV2AwemeFanTimeScope AudiencePackageCreateV2AwemeFanTimeScope = "FIFTEEN_DAYS"
)

// All allowed values of AudiencePackageCreateV2AwemeFanTimeScope enum
var AllowedAudiencePackageCreateV2AwemeFanTimeScopeEnumValues = []AudiencePackageCreateV2AwemeFanTimeScope{
	"SIXTY_DAYS",
	"THIRTY_DAYS",
	"FIFTEEN_DAYS",
}

// NewAudiencePackageCreateV2AwemeFanTimeScopeFromValue returns a pointer to a valid AudiencePackageCreateV2AwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2AwemeFanTimeScopeFromValue(v string) (*AudiencePackageCreateV2AwemeFanTimeScope, error) {
	ev := AudiencePackageCreateV2AwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2AwemeFanTimeScope: valid values are %v", v, AllowedAudiencePackageCreateV2AwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2AwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2AwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_aweme_fan_time_scope value
func (v AudiencePackageCreateV2AwemeFanTimeScope) Ptr() *AudiencePackageCreateV2AwemeFanTimeScope {
	return &v
}
