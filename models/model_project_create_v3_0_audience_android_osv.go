/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceAndroidOsv
type ProjectCreateV30AudienceAndroidOsv string

// List of project_create_v3.0_audience_android_osv
const (
	Enum_0_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "0.0"
	Enum_10_0_ProjectCreateV30AudienceAndroidOsv ProjectCreateV30AudienceAndroidOsv = "10.0"
	Enum_10_1_ProjectCreateV30AudienceAndroidOsv ProjectCreateV30AudienceAndroidOsv = "10.1"
	Enum_11_0_ProjectCreateV30AudienceAndroidOsv ProjectCreateV30AudienceAndroidOsv = "11.0"
	Enum_2_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "2.0"
	Enum_2_1_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "2.1"
	Enum_2_2_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "2.2"
	Enum_2_3_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "2.3"
	Enum_3_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "3.0"
	Enum_3_1_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "3.1"
	Enum_3_2_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "3.2"
	Enum_4_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "4.0"
	Enum_4_1_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "4.1"
	Enum_4_2_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "4.2"
	Enum_4_3_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "4.3"
	Enum_4_4_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "4.4"
	Enum_4_5_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "4.5"
	Enum_5_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "5.0"
	Enum_5_1_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "5.1"
	Enum_6_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "6.0"
	Enum_7_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "7.0"
	Enum_7_1_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "7.1"
	Enum_8_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "8.0"
	Enum_8_1_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "8.1"
	Enum_9_0_ProjectCreateV30AudienceAndroidOsv  ProjectCreateV30AudienceAndroidOsv = "9.0"
	NONE_ProjectCreateV30AudienceAndroidOsv      ProjectCreateV30AudienceAndroidOsv = "NONE"
)

// All allowed values of ProjectCreateV30AudienceAndroidOsv enum
var AllowedProjectCreateV30AudienceAndroidOsvEnumValues = []ProjectCreateV30AudienceAndroidOsv{
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

// NewProjectCreateV30AudienceAndroidOsvFromValue returns a pointer to a valid ProjectCreateV30AudienceAndroidOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceAndroidOsvFromValue(v string) (*ProjectCreateV30AudienceAndroidOsv, error) {
	ev := ProjectCreateV30AudienceAndroidOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceAndroidOsv: valid values are %v", v, AllowedProjectCreateV30AudienceAndroidOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceAndroidOsv) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceAndroidOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_android_osv value
func (v ProjectCreateV30AudienceAndroidOsv) Ptr() *ProjectCreateV30AudienceAndroidOsv {
	return &v
}
