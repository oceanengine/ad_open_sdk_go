/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2DeviceBrand
type ToolsBidSuggestV2DeviceBrand string

// List of tools_bid_suggest_v2_device_brand
const (
	NOKIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NOKIA"
	TCL_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "TCL"
	HTC_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "HTC"
	VIVO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "VIVO"
	HISENSE_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "HISENSE"
	APPLE_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "APPLE"
	OPPO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "OPPO"
	LENOVO_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "LENOVO"
	XIAOMI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "XIAOMI"
	MOTO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "MOTO"
	PEPPER_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "PEPPER"
	MEIZU_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "MEIZU"
	ONEPLUS_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "ONEPLUS"
	HONOR_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "HONOR"
	GIONEE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GIONEE"
	CHINAMOBILE_ToolsBidSuggestV2DeviceBrand ToolsBidSuggestV2DeviceBrand = "CHINAMOBILE"
	SAMSUNG_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "SAMSUNG"
	Enum_360_ToolsBidSuggestV2DeviceBrand    ToolsBidSuggestV2DeviceBrand = "360"
	COOLPAD_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "COOLPAD"
	GOOGLE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GOOGLE"
	LETV_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "LETV"
	ZTE_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "ZTE"
	QIKU_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "QIKU"
	HUAWEI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "HUAWEI"
	LG_ToolsBidSuggestV2DeviceBrand          ToolsBidSuggestV2DeviceBrand = "LG"
	SONY_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "SONY"
	NUBIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NUBIA"
	SMARTISAN_ToolsBidSuggestV2DeviceBrand   ToolsBidSuggestV2DeviceBrand = "SMARTISAN"
)

// All allowed values of ToolsBidSuggestV2DeviceBrand enum
var AllowedToolsBidSuggestV2DeviceBrandEnumValues = []ToolsBidSuggestV2DeviceBrand{
	"NOKIA",
	"TCL",
	"HTC",
	"VIVO",
	"HISENSE",
	"APPLE",
	"OPPO",
	"LENOVO",
	"XIAOMI",
	"MOTO",
	"PEPPER",
	"MEIZU",
	"ONEPLUS",
	"HONOR",
	"GIONEE",
	"CHINAMOBILE",
	"SAMSUNG",
	"360",
	"COOLPAD",
	"GOOGLE",
	"LETV",
	"ZTE",
	"QIKU",
	"HUAWEI",
	"LG",
	"SONY",
	"NUBIA",
	"SMARTISAN",
}

// NewToolsBidSuggestV2DeviceBrandFromValue returns a pointer to a valid ToolsBidSuggestV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2DeviceBrandFromValue(v string) (*ToolsBidSuggestV2DeviceBrand, error) {
	ev := ToolsBidSuggestV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2DeviceBrand: valid values are %v", v, AllowedToolsBidSuggestV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_device_brand value
func (v ToolsBidSuggestV2DeviceBrand) Ptr() *ToolsBidSuggestV2DeviceBrand {
	return &v
}
