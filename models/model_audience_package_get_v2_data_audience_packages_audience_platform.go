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

// AudiencePackageGetV2DataAudiencePackagesAudiencePlatform
type AudiencePackageGetV2DataAudiencePackagesAudiencePlatform string

// List of audience_package_get_v2_data_audience_packages_audience_platform
const (
	ANDROID_AudiencePackageGetV2DataAudiencePackagesAudiencePlatform AudiencePackageGetV2DataAudiencePackagesAudiencePlatform = "ANDROID"
	IOS_AudiencePackageGetV2DataAudiencePackagesAudiencePlatform     AudiencePackageGetV2DataAudiencePackagesAudiencePlatform = "IOS"
	IPAD_AudiencePackageGetV2DataAudiencePackagesAudiencePlatform    AudiencePackageGetV2DataAudiencePackagesAudiencePlatform = "IPAD"
	NONE_AudiencePackageGetV2DataAudiencePackagesAudiencePlatform    AudiencePackageGetV2DataAudiencePackagesAudiencePlatform = "NONE"
	PC_AudiencePackageGetV2DataAudiencePackagesAudiencePlatform      AudiencePackageGetV2DataAudiencePackagesAudiencePlatform = "PC"
	WAP_AudiencePackageGetV2DataAudiencePackagesAudiencePlatform     AudiencePackageGetV2DataAudiencePackagesAudiencePlatform = "WAP"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudiencePlatform enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudiencePlatformEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudiencePlatform{
	"ANDROID",
	"IOS",
	"IPAD",
	"NONE",
	"PC",
	"WAP",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudiencePlatformFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudiencePlatformFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudiencePlatform, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudiencePlatform: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudiencePlatform) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_platform value
func (v AudiencePackageGetV2DataAudiencePackagesAudiencePlatform) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudiencePlatform {
	return &v
}
