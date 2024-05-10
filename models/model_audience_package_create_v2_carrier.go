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

// AudiencePackageCreateV2Carrier
type AudiencePackageCreateV2Carrier string

// List of audience_package_create_v2_carrier
const (
	UNICOM_AudiencePackageCreateV2Carrier AudiencePackageCreateV2Carrier = "UNICOM"
	TELCOM_AudiencePackageCreateV2Carrier AudiencePackageCreateV2Carrier = "TELCOM"
	MOBILE_AudiencePackageCreateV2Carrier AudiencePackageCreateV2Carrier = "MOBILE"
)

// All allowed values of AudiencePackageCreateV2Carrier enum
var AllowedAudiencePackageCreateV2CarrierEnumValues = []AudiencePackageCreateV2Carrier{
	"UNICOM",
	"TELCOM",
	"MOBILE",
}

// NewAudiencePackageCreateV2CarrierFromValue returns a pointer to a valid AudiencePackageCreateV2Carrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2CarrierFromValue(v string) (*AudiencePackageCreateV2Carrier, error) {
	ev := AudiencePackageCreateV2Carrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2Carrier: valid values are %v", v, AllowedAudiencePackageCreateV2CarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2Carrier) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2CarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_carrier value
func (v AudiencePackageCreateV2Carrier) Ptr() *AudiencePackageCreateV2Carrier {
	return &v
}
