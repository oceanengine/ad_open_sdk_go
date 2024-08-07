/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceDeviceBrand
type ProjectListV30DataListAudienceDeviceBrand string

// List of project_list_v3.0_data_list_audience_device_brand
const (
	Enum_360_ProjectListV30DataListAudienceDeviceBrand    ProjectListV30DataListAudienceDeviceBrand = "360"
	APPLE_ProjectListV30DataListAudienceDeviceBrand       ProjectListV30DataListAudienceDeviceBrand = "APPLE"
	CHINAMOBILE_ProjectListV30DataListAudienceDeviceBrand ProjectListV30DataListAudienceDeviceBrand = "CHINAMOBILE"
	COOLPAD_ProjectListV30DataListAudienceDeviceBrand     ProjectListV30DataListAudienceDeviceBrand = "COOLPAD"
	GIONEE_ProjectListV30DataListAudienceDeviceBrand      ProjectListV30DataListAudienceDeviceBrand = "GIONEE"
	GOOGLE_ProjectListV30DataListAudienceDeviceBrand      ProjectListV30DataListAudienceDeviceBrand = "GOOGLE"
	HISENSE_ProjectListV30DataListAudienceDeviceBrand     ProjectListV30DataListAudienceDeviceBrand = "HISENSE"
	HONOR_ProjectListV30DataListAudienceDeviceBrand       ProjectListV30DataListAudienceDeviceBrand = "HONOR"
	HTC_ProjectListV30DataListAudienceDeviceBrand         ProjectListV30DataListAudienceDeviceBrand = "HTC"
	HUAWEI_ProjectListV30DataListAudienceDeviceBrand      ProjectListV30DataListAudienceDeviceBrand = "HUAWEI"
	LENOVO_ProjectListV30DataListAudienceDeviceBrand      ProjectListV30DataListAudienceDeviceBrand = "LENOVO"
	LETV_ProjectListV30DataListAudienceDeviceBrand        ProjectListV30DataListAudienceDeviceBrand = "LETV"
	LG_ProjectListV30DataListAudienceDeviceBrand          ProjectListV30DataListAudienceDeviceBrand = "LG"
	MEIZU_ProjectListV30DataListAudienceDeviceBrand       ProjectListV30DataListAudienceDeviceBrand = "MEIZU"
	MOTO_ProjectListV30DataListAudienceDeviceBrand        ProjectListV30DataListAudienceDeviceBrand = "MOTO"
	NOKIA_ProjectListV30DataListAudienceDeviceBrand       ProjectListV30DataListAudienceDeviceBrand = "NOKIA"
	NUBIA_ProjectListV30DataListAudienceDeviceBrand       ProjectListV30DataListAudienceDeviceBrand = "NUBIA"
	ONEPLUS_ProjectListV30DataListAudienceDeviceBrand     ProjectListV30DataListAudienceDeviceBrand = "ONEPLUS"
	OPPO_ProjectListV30DataListAudienceDeviceBrand        ProjectListV30DataListAudienceDeviceBrand = "OPPO"
	PEPPER_ProjectListV30DataListAudienceDeviceBrand      ProjectListV30DataListAudienceDeviceBrand = "PEPPER"
	QIKU_ProjectListV30DataListAudienceDeviceBrand        ProjectListV30DataListAudienceDeviceBrand = "QIKU"
	SAMSUNG_ProjectListV30DataListAudienceDeviceBrand     ProjectListV30DataListAudienceDeviceBrand = "SAMSUNG"
	SMARTISAN_ProjectListV30DataListAudienceDeviceBrand   ProjectListV30DataListAudienceDeviceBrand = "SMARTISAN"
	SONY_ProjectListV30DataListAudienceDeviceBrand        ProjectListV30DataListAudienceDeviceBrand = "SONY"
	TCL_ProjectListV30DataListAudienceDeviceBrand         ProjectListV30DataListAudienceDeviceBrand = "TCL"
	VIVO_ProjectListV30DataListAudienceDeviceBrand        ProjectListV30DataListAudienceDeviceBrand = "VIVO"
	XIAOMI_ProjectListV30DataListAudienceDeviceBrand      ProjectListV30DataListAudienceDeviceBrand = "XIAOMI"
	ZTE_ProjectListV30DataListAudienceDeviceBrand         ProjectListV30DataListAudienceDeviceBrand = "ZTE"
)

// All allowed values of ProjectListV30DataListAudienceDeviceBrand enum
var AllowedProjectListV30DataListAudienceDeviceBrandEnumValues = []ProjectListV30DataListAudienceDeviceBrand{
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

// NewProjectListV30DataListAudienceDeviceBrandFromValue returns a pointer to a valid ProjectListV30DataListAudienceDeviceBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceDeviceBrandFromValue(v string) (*ProjectListV30DataListAudienceDeviceBrand, error) {
	ev := ProjectListV30DataListAudienceDeviceBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceDeviceBrand: valid values are %v", v, AllowedProjectListV30DataListAudienceDeviceBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceDeviceBrand) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceDeviceBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_device_brand value
func (v ProjectListV30DataListAudienceDeviceBrand) Ptr() *ProjectListV30DataListAudienceDeviceBrand {
	return &v
}
