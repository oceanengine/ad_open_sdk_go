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

// AudiencePackageGetV2DataAudiencePackagesAudienceAc
type AudiencePackageGetV2DataAudiencePackagesAudienceAc string

// List of audience_package_get_v2_data_audience_packages_audience_ac
const (
	Enum_2_G_AudiencePackageGetV2DataAudiencePackagesAudienceAc AudiencePackageGetV2DataAudiencePackagesAudienceAc = "2G"
	Enum_3_G_AudiencePackageGetV2DataAudiencePackagesAudienceAc AudiencePackageGetV2DataAudiencePackagesAudienceAc = "3G"
	Enum_4_G_AudiencePackageGetV2DataAudiencePackagesAudienceAc AudiencePackageGetV2DataAudiencePackagesAudienceAc = "4G"
	Enum_5_G_AudiencePackageGetV2DataAudiencePackagesAudienceAc AudiencePackageGetV2DataAudiencePackagesAudienceAc = "5G"
	WIFI_AudiencePackageGetV2DataAudiencePackagesAudienceAc     AudiencePackageGetV2DataAudiencePackagesAudienceAc = "WIFI"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceAc enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAcEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceAcFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceAcFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceAc, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceAc: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceAc) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_ac value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceAc) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceAc {
	return &v
}