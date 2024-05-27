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

// AudiencePackageUpdateV2AndroidOsv
type AudiencePackageUpdateV2AndroidOsv string

// List of audience_package_update_v2_android_osv
const (
	Enum_9_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "9.0"
	Enum_10_0_AudiencePackageUpdateV2AndroidOsv AudiencePackageUpdateV2AndroidOsv = "10.0"
	Enum_5_1_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "5.1"
	Enum_3_1_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "3.1"
	Enum_11_0_AudiencePackageUpdateV2AndroidOsv AudiencePackageUpdateV2AndroidOsv = "11.0"
	Enum_4_3_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "4.3"
	Enum_7_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "7.0"
	Enum_2_1_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "2.1"
	NONE_AudiencePackageUpdateV2AndroidOsv      AudiencePackageUpdateV2AndroidOsv = "NONE"
	Enum_7_1_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "7.1"
	Enum_10_1_AudiencePackageUpdateV2AndroidOsv AudiencePackageUpdateV2AndroidOsv = "10.1"
	Enum_8_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "8.0"
	Enum_4_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "4.0"
	Enum_4_4_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "4.4"
	Enum_3_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "3.0"
	Enum_6_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "6.0"
	Enum_4_5_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "4.5"
	Enum_2_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "2.0"
	Enum_2_3_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "2.3"
	Enum_0_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "0.0"
	Enum_2_2_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "2.2"
	Enum_8_1_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "8.1"
	Enum_4_1_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "4.1"
	Enum_5_0_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "5.0"
	Enum_3_2_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "3.2"
	Enum_4_2_AudiencePackageUpdateV2AndroidOsv  AudiencePackageUpdateV2AndroidOsv = "4.2"
)

// All allowed values of AudiencePackageUpdateV2AndroidOsv enum
var AllowedAudiencePackageUpdateV2AndroidOsvEnumValues = []AudiencePackageUpdateV2AndroidOsv{
	"9.0",
	"10.0",
	"5.1",
	"3.1",
	"11.0",
	"4.3",
	"7.0",
	"2.1",
	"NONE",
	"7.1",
	"10.1",
	"8.0",
	"4.0",
	"4.4",
	"3.0",
	"6.0",
	"4.5",
	"2.0",
	"2.3",
	"0.0",
	"2.2",
	"8.1",
	"4.1",
	"5.0",
	"3.2",
	"4.2",
}

// NewAudiencePackageUpdateV2AndroidOsvFromValue returns a pointer to a valid AudiencePackageUpdateV2AndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2AndroidOsvFromValue(v string) (*AudiencePackageUpdateV2AndroidOsv, error) {
	ev := AudiencePackageUpdateV2AndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2AndroidOsv: valid values are %v", v, AllowedAudiencePackageUpdateV2AndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2AndroidOsv) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2AndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_android_osv value
func (v AudiencePackageUpdateV2AndroidOsv) Ptr() *AudiencePackageUpdateV2AndroidOsv {
	return &v
}
