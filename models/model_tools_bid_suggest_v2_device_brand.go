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

// ToolsBidSuggestV2DeviceBrand
type ToolsBidSuggestV2DeviceBrand string

// List of tools_bid_suggest_v2_device_brand
const (
	HUAWEI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "HUAWEI"
	QIKU_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "QIKU"
	HISENSE_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "HISENSE"
	LENOVO_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "LENOVO"
	SONY_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "SONY"
	MEIZU_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "MEIZU"
	CHINAMOBILE_ToolsBidSuggestV2DeviceBrand ToolsBidSuggestV2DeviceBrand = "CHINAMOBILE"
	HTC_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "HTC"
	ONEPLUS_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "ONEPLUS"
	LG_ToolsBidSuggestV2DeviceBrand          ToolsBidSuggestV2DeviceBrand = "LG"
	COOLPAD_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "COOLPAD"
	PEPPER_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "PEPPER"
	SMARTISAN_ToolsBidSuggestV2DeviceBrand   ToolsBidSuggestV2DeviceBrand = "SMARTISAN"
	LETV_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "LETV"
	SAMSUNG_ToolsBidSuggestV2DeviceBrand     ToolsBidSuggestV2DeviceBrand = "SAMSUNG"
	ZTE_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "ZTE"
	GOOGLE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GOOGLE"
	TCL_ToolsBidSuggestV2DeviceBrand         ToolsBidSuggestV2DeviceBrand = "TCL"
	HONOR_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "HONOR"
	GIONEE_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "GIONEE"
	Enum_360_ToolsBidSuggestV2DeviceBrand    ToolsBidSuggestV2DeviceBrand = "360"
	OPPO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "OPPO"
	NOKIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NOKIA"
	NUBIA_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "NUBIA"
	XIAOMI_ToolsBidSuggestV2DeviceBrand      ToolsBidSuggestV2DeviceBrand = "XIAOMI"
	APPLE_ToolsBidSuggestV2DeviceBrand       ToolsBidSuggestV2DeviceBrand = "APPLE"
	MOTO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "MOTO"
	VIVO_ToolsBidSuggestV2DeviceBrand        ToolsBidSuggestV2DeviceBrand = "VIVO"
)

// All allowed values of ToolsBidSuggestV2DeviceBrand enum
var AllowedToolsBidSuggestV2DeviceBrandEnumValues = []ToolsBidSuggestV2DeviceBrand{
	"HUAWEI",
	"QIKU",
	"HISENSE",
	"LENOVO",
	"SONY",
	"MEIZU",
	"CHINAMOBILE",
	"HTC",
	"ONEPLUS",
	"LG",
	"COOLPAD",
	"PEPPER",
	"SMARTISAN",
	"LETV",
	"SAMSUNG",
	"ZTE",
	"GOOGLE",
	"TCL",
	"HONOR",
	"GIONEE",
	"360",
	"OPPO",
	"NOKIA",
	"NUBIA",
	"XIAOMI",
	"APPLE",
	"MOTO",
	"VIVO",
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
