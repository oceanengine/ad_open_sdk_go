/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2AndroidOsv
type AudiencePackageCreateV2AndroidOsv string

// List of audience_package_create_v2_android_osv
const (
	Enum_3_1_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "3.1"
	Enum_7_1_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "7.1"
	Enum_2_3_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "2.3"
	Enum_2_1_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "2.1"
	Enum_9_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "9.0"
	Enum_8_1_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "8.1"
	Enum_2_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "2.0"
	Enum_4_4_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "4.4"
	Enum_3_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "3.0"
	Enum_7_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "7.0"
	Enum_6_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "6.0"
	Enum_4_1_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "4.1"
	Enum_10_0_AudiencePackageCreateV2AndroidOsv AudiencePackageCreateV2AndroidOsv = "10.0"
	Enum_2_2_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "2.2"
	Enum_0_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "0.0"
	Enum_4_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "4.0"
	Enum_4_2_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "4.2"
	Enum_3_2_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "3.2"
	Enum_5_1_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "5.1"
	Enum_11_0_AudiencePackageCreateV2AndroidOsv AudiencePackageCreateV2AndroidOsv = "11.0"
	Enum_4_3_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "4.3"
	Enum_10_1_AudiencePackageCreateV2AndroidOsv AudiencePackageCreateV2AndroidOsv = "10.1"
	Enum_5_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "5.0"
	NONE_AudiencePackageCreateV2AndroidOsv      AudiencePackageCreateV2AndroidOsv = "NONE"
	Enum_4_5_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "4.5"
	Enum_8_0_AudiencePackageCreateV2AndroidOsv  AudiencePackageCreateV2AndroidOsv = "8.0"
)

// All allowed values of AudiencePackageCreateV2AndroidOsv enum
var AllowedAudiencePackageCreateV2AndroidOsvEnumValues = []AudiencePackageCreateV2AndroidOsv{
	"3.1",
	"7.1",
	"2.3",
	"2.1",
	"9.0",
	"8.1",
	"2.0",
	"4.4",
	"3.0",
	"7.0",
	"6.0",
	"4.1",
	"10.0",
	"2.2",
	"0.0",
	"4.0",
	"4.2",
	"3.2",
	"5.1",
	"11.0",
	"4.3",
	"10.1",
	"5.0",
	"NONE",
	"4.5",
	"8.0",
}

// NewAudiencePackageCreateV2AndroidOsvFromValue returns a pointer to a valid AudiencePackageCreateV2AndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2AndroidOsvFromValue(v string) (*AudiencePackageCreateV2AndroidOsv, error) {
	ev := AudiencePackageCreateV2AndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2AndroidOsv: valid values are %v", v, AllowedAudiencePackageCreateV2AndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2AndroidOsv) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2AndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_android_osv value
func (v AudiencePackageCreateV2AndroidOsv) Ptr() *AudiencePackageCreateV2AndroidOsv {
	return &v
}
