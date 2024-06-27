/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand
type AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand string

// List of audience_package_get_v2_data_audience_packages_audience_device_brand
const (
	Enum_360_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand    AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "360"
	APPLE_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand       AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "APPLE"
	CHINAMOBILE_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "CHINAMOBILE"
	COOLPAD_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand     AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "COOLPAD"
	GIONEE_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand      AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "GIONEE"
	GOOGLE_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand      AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "GOOGLE"
	HISENSE_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand     AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "HISENSE"
	HONOR_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand       AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "HONOR"
	HTC_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand         AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "HTC"
	HUAWEI_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand      AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "HUAWEI"
	LENOVO_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand      AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "LENOVO"
	LETV_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand        AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "LETV"
	LG_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand          AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "LG"
	MEIZU_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand       AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "MEIZU"
	MOTO_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand        AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "MOTO"
	NOKIA_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand       AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "NOKIA"
	NUBIA_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand       AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "NUBIA"
	ONEPLUS_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand     AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "ONEPLUS"
	OPPO_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand        AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "OPPO"
	PEPPER_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand      AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "PEPPER"
	QIKU_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand        AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "QIKU"
	SAMSUNG_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand     AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "SAMSUNG"
	SMARTISAN_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand   AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "SMARTISAN"
	SONY_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand        AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "SONY"
	TCL_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand         AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "TCL"
	VIVO_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand        AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "VIVO"
	XIAOMI_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand      AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "XIAOMI"
	ZTE_AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand         AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand = "ZTE"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrandEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand{
	"360",
	"APPLE",
	"CHINAMOBILE",
	"COOLPAD",
	"GIONEE",
	"GOOGLE",
	"HISENSE",
	"HONOR",
	"HTC",
	"HUAWEI",
	"LENOVO",
	"LETV",
	"LG",
	"MEIZU",
	"MOTO",
	"NOKIA",
	"NUBIA",
	"ONEPLUS",
	"OPPO",
	"PEPPER",
	"QIKU",
	"SAMSUNG",
	"SMARTISAN",
	"SONY",
	"TCL",
	"VIVO",
	"XIAOMI",
	"ZTE",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrandFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrandFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_device_brand value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceDeviceBrand {
	return &v
}
