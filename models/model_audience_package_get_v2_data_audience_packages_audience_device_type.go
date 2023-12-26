/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType
type AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType string

// List of audience_package_get_v2_data_audience_packages_audience_device_type
const (
	MOBILE_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType = "MOBILE"
	PAD_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType    AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType = "PAD"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDeviceTypeEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType{
	"MOBILE",
	"PAD",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceDeviceTypeFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceDeviceTypeFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_device_type value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceDeviceType {
	return &v
}
