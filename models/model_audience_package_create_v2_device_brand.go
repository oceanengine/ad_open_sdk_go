/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2DeviceBrand
type AudiencePackageCreateV2DeviceBrand string

// List of audience_package_create_v2_device_brand
const (
	LG_AudiencePackageCreateV2DeviceBrand          AudiencePackageCreateV2DeviceBrand = "LG"
	HTC_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "HTC"
	NUBIA_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "NUBIA"
	APPLE_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "APPLE"
	MOTO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "MOTO"
	GOOGLE_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "GOOGLE"
	SMARTISAN_AudiencePackageCreateV2DeviceBrand   AudiencePackageCreateV2DeviceBrand = "SMARTISAN"
	Enum_360_AudiencePackageCreateV2DeviceBrand    AudiencePackageCreateV2DeviceBrand = "360"
	GIONEE_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "GIONEE"
	ONEPLUS_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "ONEPLUS"
	LENOVO_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "LENOVO"
	NOKIA_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "NOKIA"
	SAMSUNG_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "SAMSUNG"
	LETV_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "LETV"
	CHINAMOBILE_AudiencePackageCreateV2DeviceBrand AudiencePackageCreateV2DeviceBrand = "CHINAMOBILE"
	HISENSE_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "HISENSE"
	MEIZU_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "MEIZU"
	OPPO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "OPPO"
	QIKU_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "QIKU"
	VIVO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "VIVO"
	PEPPER_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "PEPPER"
	COOLPAD_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "COOLPAD"
	TCL_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "TCL"
	SONY_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "SONY"
	ZTE_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "ZTE"
	HONOR_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "HONOR"
	XIAOMI_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "XIAOMI"
	HUAWEI_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "HUAWEI"
)

// All allowed values of AudiencePackageCreateV2DeviceBrand enum
var AllowedAudiencePackageCreateV2DeviceBrandEnumValues = []AudiencePackageCreateV2DeviceBrand{
	"LG",
	"HTC",
	"NUBIA",
	"APPLE",
	"MOTO",
	"GOOGLE",
	"SMARTISAN",
	"360",
	"GIONEE",
	"ONEPLUS",
	"LENOVO",
	"NOKIA",
	"SAMSUNG",
	"LETV",
	"CHINAMOBILE",
	"HISENSE",
	"MEIZU",
	"OPPO",
	"QIKU",
	"VIVO",
	"PEPPER",
	"COOLPAD",
	"TCL",
	"SONY",
	"ZTE",
	"HONOR",
	"XIAOMI",
	"HUAWEI",
}

// NewAudiencePackageCreateV2DeviceBrandFromValue returns a pointer to a valid AudiencePackageCreateV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2DeviceBrandFromValue(v string) (*AudiencePackageCreateV2DeviceBrand, error) {
	ev := AudiencePackageCreateV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2DeviceBrand: valid values are %v", v, AllowedAudiencePackageCreateV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_device_brand value
func (v AudiencePackageCreateV2DeviceBrand) Ptr() *AudiencePackageCreateV2DeviceBrand {
	return &v
}
