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

// ProjectUpdateV30AudienceAndroidOsv
type ProjectUpdateV30AudienceAndroidOsv string

// List of project_update_v3.0_audience_android_osv
const (
	Enum_0_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "0.0"
	Enum_10_0_ProjectUpdateV30AudienceAndroidOsv ProjectUpdateV30AudienceAndroidOsv = "10.0"
	Enum_10_1_ProjectUpdateV30AudienceAndroidOsv ProjectUpdateV30AudienceAndroidOsv = "10.1"
	Enum_11_0_ProjectUpdateV30AudienceAndroidOsv ProjectUpdateV30AudienceAndroidOsv = "11.0"
	Enum_2_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "2.0"
	Enum_2_1_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "2.1"
	Enum_2_2_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "2.2"
	Enum_2_3_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "2.3"
	Enum_3_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "3.0"
	Enum_3_1_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "3.1"
	Enum_3_2_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "3.2"
	Enum_4_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "4.0"
	Enum_4_1_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "4.1"
	Enum_4_2_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "4.2"
	Enum_4_3_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "4.3"
	Enum_4_4_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "4.4"
	Enum_4_5_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "4.5"
	Enum_5_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "5.0"
	Enum_5_1_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "5.1"
	Enum_6_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "6.0"
	Enum_7_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "7.0"
	Enum_7_1_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "7.1"
	Enum_8_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "8.0"
	Enum_8_1_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "8.1"
	Enum_9_0_ProjectUpdateV30AudienceAndroidOsv  ProjectUpdateV30AudienceAndroidOsv = "9.0"
	NONE_ProjectUpdateV30AudienceAndroidOsv      ProjectUpdateV30AudienceAndroidOsv = "NONE"
)

// All allowed values of ProjectUpdateV30AudienceAndroidOsv enum
var AllowedProjectUpdateV30AudienceAndroidOsvEnumValues = []ProjectUpdateV30AudienceAndroidOsv{
	"0.0",
	"10.0",
	"10.1",
	"11.0",
	"2.0",
	"2.1",
	"2.2",
	"2.3",
	"3.0",
	"3.1",
	"3.2",
	"4.0",
	"4.1",
	"4.2",
	"4.3",
	"4.4",
	"4.5",
	"5.0",
	"5.1",
	"6.0",
	"7.0",
	"7.1",
	"8.0",
	"8.1",
	"9.0",
	"NONE",
}

// NewProjectUpdateV30AudienceAndroidOsvFromValue returns a pointer to a valid ProjectUpdateV30AudienceAndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceAndroidOsvFromValue(v string) (*ProjectUpdateV30AudienceAndroidOsv, error) {
	ev := ProjectUpdateV30AudienceAndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceAndroidOsv: valid values are %v", v, AllowedProjectUpdateV30AudienceAndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceAndroidOsv) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceAndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_android_osv value
func (v ProjectUpdateV30AudienceAndroidOsv) Ptr() *ProjectUpdateV30AudienceAndroidOsv {
	return &v
}
