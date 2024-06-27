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

// ToolsEstimateAudienceV2DeviceBrand
type ToolsEstimateAudienceV2DeviceBrand string

// List of tools_estimate_audience_v2_device_brand
const (
	HUAWEI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "HUAWEI"
	QIKU_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "QIKU"
	HISENSE_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "HISENSE"
	LENOVO_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "LENOVO"
	SONY_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "SONY"
	MEIZU_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "MEIZU"
	CHINAMOBILE_ToolsEstimateAudienceV2DeviceBrand ToolsEstimateAudienceV2DeviceBrand = "CHINAMOBILE"
	HTC_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "HTC"
	ONEPLUS_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "ONEPLUS"
	LG_ToolsEstimateAudienceV2DeviceBrand          ToolsEstimateAudienceV2DeviceBrand = "LG"
	COOLPAD_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "COOLPAD"
	PEPPER_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "PEPPER"
	SMARTISAN_ToolsEstimateAudienceV2DeviceBrand   ToolsEstimateAudienceV2DeviceBrand = "SMARTISAN"
	LETV_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "LETV"
	SAMSUNG_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "SAMSUNG"
	ZTE_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "ZTE"
	GOOGLE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GOOGLE"
	TCL_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "TCL"
	HONOR_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "HONOR"
	GIONEE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GIONEE"
	Enum_360_ToolsEstimateAudienceV2DeviceBrand    ToolsEstimateAudienceV2DeviceBrand = "360"
	OPPO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "OPPO"
	NOKIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NOKIA"
	NUBIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NUBIA"
	XIAOMI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "XIAOMI"
	APPLE_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "APPLE"
	MOTO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "MOTO"
	VIVO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "VIVO"
)

// All allowed values of ToolsEstimateAudienceV2DeviceBrand enum
var AllowedToolsEstimateAudienceV2DeviceBrandEnumValues = []ToolsEstimateAudienceV2DeviceBrand{
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

// NewToolsEstimateAudienceV2DeviceBrandFromValue returns a pointer to a valid ToolsEstimateAudienceV2DeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2DeviceBrandFromValue(v string) (*ToolsEstimateAudienceV2DeviceBrand, error) {
	ev := ToolsEstimateAudienceV2DeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2DeviceBrand: valid values are %v", v, AllowedToolsEstimateAudienceV2DeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2DeviceBrand) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2DeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_device_brand value
func (v ToolsEstimateAudienceV2DeviceBrand) Ptr() *ToolsEstimateAudienceV2DeviceBrand {
	return &v
}
