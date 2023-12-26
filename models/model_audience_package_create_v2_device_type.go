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

// AudiencePackageCreateV2DeviceType
type AudiencePackageCreateV2DeviceType string

// List of audience_package_create_v2_device_type
const (
	MOBILE_AudiencePackageCreateV2DeviceType AudiencePackageCreateV2DeviceType = "MOBILE"
	PAD_AudiencePackageCreateV2DeviceType    AudiencePackageCreateV2DeviceType = "PAD"
)

// All allowed values of AudiencePackageCreateV2DeviceType enum
var AllowedAudiencePackageCreateV2DeviceTypeEnumValues = []AudiencePackageCreateV2DeviceType{
	"MOBILE",
	"PAD",
}

// NewAudiencePackageCreateV2DeviceTypeFromValue returns a pointer to a valid AudiencePackageCreateV2DeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2DeviceTypeFromValue(v string) (*AudiencePackageCreateV2DeviceType, error) {
	ev := AudiencePackageCreateV2DeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2DeviceType: valid values are %v", v, AllowedAudiencePackageCreateV2DeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2DeviceType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2DeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_device_type value
func (v AudiencePackageCreateV2DeviceType) Ptr() *AudiencePackageCreateV2DeviceType {
	return &v
}
