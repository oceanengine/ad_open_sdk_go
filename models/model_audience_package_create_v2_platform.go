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

// AudiencePackageCreateV2Platform
type AudiencePackageCreateV2Platform string

// List of audience_package_create_v2_platform
const (
	WAP_AudiencePackageCreateV2Platform     AudiencePackageCreateV2Platform = "WAP"
	IPAD_AudiencePackageCreateV2Platform    AudiencePackageCreateV2Platform = "IPAD"
	IOS_AudiencePackageCreateV2Platform     AudiencePackageCreateV2Platform = "IOS"
	PC_AudiencePackageCreateV2Platform      AudiencePackageCreateV2Platform = "PC"
	ANDROID_AudiencePackageCreateV2Platform AudiencePackageCreateV2Platform = "ANDROID"
)

// All allowed values of AudiencePackageCreateV2Platform enum
var AllowedAudiencePackageCreateV2PlatformEnumValues = []AudiencePackageCreateV2Platform{
	"WAP",
	"IPAD",
	"IOS",
	"PC",
	"ANDROID",
}

// NewAudiencePackageCreateV2PlatformFromValue returns a pointer to a valid AudiencePackageCreateV2Platform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2PlatformFromValue(v string) (*AudiencePackageCreateV2Platform, error) {
	ev := AudiencePackageCreateV2Platform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2Platform: valid values are %v", v, AllowedAudiencePackageCreateV2PlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2Platform) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2PlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_platform value
func (v AudiencePackageCreateV2Platform) Ptr() *AudiencePackageCreateV2Platform {
	return &v
}
