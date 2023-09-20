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

// AudiencePackageUpdateV2DeviceBrand
type AudiencePackageUpdateV2DeviceBrand string

// List of audience_package_update_v2_device_brand
const (
	QIKU_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "QIKU"
	OPPO_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "OPPO"
	MOTO_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "MOTO"
	SAMSUNG_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "SAMSUNG"
	HTC_AudiencePackageUpdateV2DeviceBrand         AudiencePackageUpdateV2DeviceBrand = "HTC"
	CHINAMOBILE_AudiencePackageUpdateV2DeviceBrand AudiencePackageUpdateV2DeviceBrand = "CHINAMOBILE"
	XIAOMI_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "XIAOMI"
	ONEPLUS_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "ONEPLUS"
	HUAWEI_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "HUAWEI"
	COOLPAD_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "COOLPAD"
	HONOR_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "HONOR"
	ZTE_AudiencePackageUpdateV2DeviceBrand         AudiencePackageUpdateV2DeviceBrand = "ZTE"
	TCL_AudiencePackageUpdateV2DeviceBrand         AudiencePackageUpdateV2DeviceBrand = "TCL"
	NOKIA_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "NOKIA"
	APPLE_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "APPLE"
	MEIZU_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "MEIZU"
	LENOVO_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "LENOVO"
	LG_AudiencePackageUpdateV2DeviceBrand          AudiencePackageUpdateV2DeviceBrand = "LG"
	VIVO_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "VIVO"
	PEPPER_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "PEPPER"
	Enum_360_AudiencePackageUpdateV2DeviceBrand    AudiencePackageUpdateV2DeviceBrand = "360"
	GOOGLE_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "GOOGLE"
	SONY_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "SONY"
	HISENSE_AudiencePackageUpdateV2DeviceBrand     AudiencePackageUpdateV2DeviceBrand = "HISENSE"
	SMARTISAN_AudiencePackageUpdateV2DeviceBrand   AudiencePackageUpdateV2DeviceBrand = "SMARTISAN"
	LETV_AudiencePackageUpdateV2DeviceBrand        AudiencePackageUpdateV2DeviceBrand = "LETV"
	GIONEE_AudiencePackageUpdateV2DeviceBrand      AudiencePackageUpdateV2DeviceBrand = "GIONEE"
	NUBIA_AudiencePackageUpdateV2DeviceBrand       AudiencePackageUpdateV2DeviceBrand = "NUBIA"
)

// All allowed values of AudiencePackageUpdateV2DeviceBrand enum
var AllowedAudiencePackageUpdateV2DeviceBrandEnumValues = []AudiencePackageUpdateV2DeviceBrand{
	"QIKU",
	"OPPO",
	"MOTO",
	"SAMSUNG",
	"HTC",
	"CHINAMOBILE",
	"XIAOMI",
	"ONEPLUS",
	"HUAWEI",
	"COOLPAD",
	"HONOR",
	"ZTE",
	"TCL",
	"NOKIA",
	"APPLE",
	"MEIZU",
	"LENOVO",
	"LG",
	"VIVO",
	"PEPPER",
	"360",
	"GOOGLE",
	"SONY",
	"HISENSE",
	"SMARTISAN",
	"LETV",
	"GIONEE",
	"NUBIA",
}

// NewAudiencePackageUpdateV2DeviceBrandFromValue returns a pointer to a valid AudiencePackageUpdateV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2DeviceBrandFromValue(v string) (*AudiencePackageUpdateV2DeviceBrand, error) {
	ev := AudiencePackageUpdateV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2DeviceBrand: valid values are %v", v, AllowedAudiencePackageUpdateV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_device_brand value
func (v AudiencePackageUpdateV2DeviceBrand) Ptr() *AudiencePackageUpdateV2DeviceBrand {
	return &v
}
