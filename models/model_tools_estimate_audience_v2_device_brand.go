/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
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
	LETV_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "LETV"
	SMARTISAN_ToolsEstimateAudienceV2DeviceBrand   ToolsEstimateAudienceV2DeviceBrand = "SMARTISAN"
	SAMSUNG_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "SAMSUNG"
	LG_ToolsEstimateAudienceV2DeviceBrand          ToolsEstimateAudienceV2DeviceBrand = "LG"
	GOOGLE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GOOGLE"
	GIONEE_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "GIONEE"
	COOLPAD_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "COOLPAD"
	VIVO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "VIVO"
	MEIZU_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "MEIZU"
	HUAWEI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "HUAWEI"
	SONY_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "SONY"
	HISENSE_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "HISENSE"
	LENOVO_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "LENOVO"
	ONEPLUS_ToolsEstimateAudienceV2DeviceBrand     ToolsEstimateAudienceV2DeviceBrand = "ONEPLUS"
	MOTO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "MOTO"
	HTC_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "HTC"
	TCL_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "TCL"
	APPLE_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "APPLE"
	HONOR_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "HONOR"
	NUBIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NUBIA"
	PEPPER_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "PEPPER"
	XIAOMI_ToolsEstimateAudienceV2DeviceBrand      ToolsEstimateAudienceV2DeviceBrand = "XIAOMI"
	OPPO_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "OPPO"
	Enum_360_ToolsEstimateAudienceV2DeviceBrand    ToolsEstimateAudienceV2DeviceBrand = "360"
	NOKIA_ToolsEstimateAudienceV2DeviceBrand       ToolsEstimateAudienceV2DeviceBrand = "NOKIA"
	CHINAMOBILE_ToolsEstimateAudienceV2DeviceBrand ToolsEstimateAudienceV2DeviceBrand = "CHINAMOBILE"
	ZTE_ToolsEstimateAudienceV2DeviceBrand         ToolsEstimateAudienceV2DeviceBrand = "ZTE"
	QIKU_ToolsEstimateAudienceV2DeviceBrand        ToolsEstimateAudienceV2DeviceBrand = "QIKU"
)

// All allowed values of ToolsEstimateAudienceV2DeviceBrand enum
var AllowedToolsEstimateAudienceV2DeviceBrandEnumValues = []ToolsEstimateAudienceV2DeviceBrand{
	"LETV",
	"SMARTISAN",
	"SAMSUNG",
	"LG",
	"GOOGLE",
	"GIONEE",
	"COOLPAD",
	"VIVO",
	"MEIZU",
	"HUAWEI",
	"SONY",
	"HISENSE",
	"LENOVO",
	"ONEPLUS",
	"MOTO",
	"HTC",
	"TCL",
	"APPLE",
	"HONOR",
	"NUBIA",
	"PEPPER",
	"XIAOMI",
	"OPPO",
	"360",
	"NOKIA",
	"CHINAMOBILE",
	"ZTE",
	"QIKU",
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
