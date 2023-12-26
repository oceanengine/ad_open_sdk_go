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

// AdGetV2DataAudienceDeviceBrand
type AdGetV2DataAudienceDeviceBrand string

// List of ad_get_v2_data_audience_device_brand
const (
	APPLE_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "APPLE"
	MEIZU_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "MEIZU"
	NOKIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NOKIA"
	SONY_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "SONY"
	SMARTISAN_AdGetV2DataAudienceDeviceBrand   AdGetV2DataAudienceDeviceBrand = "SMARTISAN"
	GOOGLE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GOOGLE"
	XIAOMI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "XIAOMI"
	LETV_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "LETV"
	ZTE_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "ZTE"
	PEPPER_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "PEPPER"
	Enum_360_AdGetV2DataAudienceDeviceBrand    AdGetV2DataAudienceDeviceBrand = "360"
	HUAWEI_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "HUAWEI"
	QIKU_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "QIKU"
	GIONEE_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "GIONEE"
	MOTO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "MOTO"
	LG_AdGetV2DataAudienceDeviceBrand          AdGetV2DataAudienceDeviceBrand = "LG"
	OPPO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "OPPO"
	CHINAMOBILE_AdGetV2DataAudienceDeviceBrand AdGetV2DataAudienceDeviceBrand = "CHINAMOBILE"
	SAMSUNG_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "SAMSUNG"
	HISENSE_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "HISENSE"
	NUBIA_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "NUBIA"
	TCL_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "TCL"
	HTC_AdGetV2DataAudienceDeviceBrand         AdGetV2DataAudienceDeviceBrand = "HTC"
	ONEPLUS_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "ONEPLUS"
	HONOR_AdGetV2DataAudienceDeviceBrand       AdGetV2DataAudienceDeviceBrand = "HONOR"
	LENOVO_AdGetV2DataAudienceDeviceBrand      AdGetV2DataAudienceDeviceBrand = "LENOVO"
	COOLPAD_AdGetV2DataAudienceDeviceBrand     AdGetV2DataAudienceDeviceBrand = "COOLPAD"
	VIVO_AdGetV2DataAudienceDeviceBrand        AdGetV2DataAudienceDeviceBrand = "VIVO"
)

// All allowed values of AdGetV2DataAudienceDeviceBrand enum
var AllowedAdGetV2DataAudienceDeviceBrandEnumValues = []AdGetV2DataAudienceDeviceBrand{
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

// NewAdGetV2DataAudienceDeviceBrandFromValue returns a pointer to a valid AdGetV2DataAudienceDeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceDeviceBrandFromValue(v string) (*AdGetV2DataAudienceDeviceBrand, error) {
	ev := AdGetV2DataAudienceDeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceDeviceBrand: valid values are %v", v, AllowedAdGetV2DataAudienceDeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceDeviceBrand) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceDeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_device_brand value
func (v AdGetV2DataAudienceDeviceBrand) Ptr() *AdGetV2DataAudienceDeviceBrand {
	return &v
}
