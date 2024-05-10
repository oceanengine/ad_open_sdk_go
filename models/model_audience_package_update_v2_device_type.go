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

// AudiencePackageUpdateV2DeviceType
type AudiencePackageUpdateV2DeviceType string

// List of audience_package_update_v2_device_type
const (
	MOBILE_AudiencePackageUpdateV2DeviceType AudiencePackageUpdateV2DeviceType = "MOBILE"
	PAD_AudiencePackageUpdateV2DeviceType    AudiencePackageUpdateV2DeviceType = "PAD"
)

// All allowed values of AudiencePackageUpdateV2DeviceType enum
var AllowedAudiencePackageUpdateV2DeviceTypeEnumValues = []AudiencePackageUpdateV2DeviceType{
	"MOBILE",
	"PAD",
}

// NewAudiencePackageUpdateV2DeviceTypeFromValue returns a pointer to a valid AudiencePackageUpdateV2DeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2DeviceTypeFromValue(v string) (*AudiencePackageUpdateV2DeviceType, error) {
	ev := AudiencePackageUpdateV2DeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2DeviceType: valid values are %v", v, AllowedAudiencePackageUpdateV2DeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2DeviceType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2DeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_device_type value
func (v AudiencePackageUpdateV2DeviceType) Ptr() *AudiencePackageUpdateV2DeviceType {
	return &v
}
