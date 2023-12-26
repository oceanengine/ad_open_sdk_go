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

// AudiencePackageCreateV2DeviceBrand
type AudiencePackageCreateV2DeviceBrand string

// List of audience_package_create_v2_device_brand
const (
	APPLE_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "APPLE"
	MEIZU_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "MEIZU"
	NOKIA_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "NOKIA"
	SONY_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "SONY"
	SMARTISAN_AudiencePackageCreateV2DeviceBrand   AudiencePackageCreateV2DeviceBrand = "SMARTISAN"
	GOOGLE_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "GOOGLE"
	XIAOMI_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "XIAOMI"
	LETV_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "LETV"
	ZTE_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "ZTE"
	PEPPER_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "PEPPER"
	Enum_360_AudiencePackageCreateV2DeviceBrand    AudiencePackageCreateV2DeviceBrand = "360"
	HUAWEI_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "HUAWEI"
	QIKU_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "QIKU"
	GIONEE_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "GIONEE"
	MOTO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "MOTO"
	LG_AudiencePackageCreateV2DeviceBrand          AudiencePackageCreateV2DeviceBrand = "LG"
	OPPO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "OPPO"
	CHINAMOBILE_AudiencePackageCreateV2DeviceBrand AudiencePackageCreateV2DeviceBrand = "CHINAMOBILE"
	SAMSUNG_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "SAMSUNG"
	HISENSE_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "HISENSE"
	NUBIA_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "NUBIA"
	TCL_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "TCL"
	HTC_AudiencePackageCreateV2DeviceBrand         AudiencePackageCreateV2DeviceBrand = "HTC"
	ONEPLUS_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "ONEPLUS"
	HONOR_AudiencePackageCreateV2DeviceBrand       AudiencePackageCreateV2DeviceBrand = "HONOR"
	LENOVO_AudiencePackageCreateV2DeviceBrand      AudiencePackageCreateV2DeviceBrand = "LENOVO"
	COOLPAD_AudiencePackageCreateV2DeviceBrand     AudiencePackageCreateV2DeviceBrand = "COOLPAD"
	VIVO_AudiencePackageCreateV2DeviceBrand        AudiencePackageCreateV2DeviceBrand = "VIVO"
)

// All allowed values of AudiencePackageCreateV2DeviceBrand enum
var AllowedAudiencePackageCreateV2DeviceBrandEnumValues = []AudiencePackageCreateV2DeviceBrand{
	"APPLE",
	"MEIZU",
	"NOKIA",
	"SONY",
	"SMARTISAN",
	"GOOGLE",
	"XIAOMI",
	"LETV",
	"ZTE",
	"PEPPER",
	"360",
	"HUAWEI",
	"QIKU",
	"GIONEE",
	"MOTO",
	"LG",
	"OPPO",
	"CHINAMOBILE",
	"SAMSUNG",
	"HISENSE",
	"NUBIA",
	"TCL",
	"HTC",
	"ONEPLUS",
	"HONOR",
	"LENOVO",
	"COOLPAD",
	"VIVO",
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
