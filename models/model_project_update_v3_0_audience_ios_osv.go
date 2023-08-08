/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceIosOsv
type ProjectUpdateV30AudienceIosOsv string

// List of project_update_v3.0_audience_ios_osv
const (
	Enum_0_0_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "0.0"
	Enum_10_1_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "10.1"
	Enum_10_2_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "10.2"
	Enum_10_3_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "10.3"
	Enum_11_0_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "11.0"
	Enum_11_1_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "11.1"
	Enum_11_2_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "11.2"
	Enum_11_3_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "11.3"
	Enum_11_4_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "11.4"
	Enum_12_0_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "12.0"
	Enum_12_1_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "12.1"
	Enum_12_2_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "12.2"
	Enum_12_3_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "12.3"
	Enum_12_4_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "12.4"
	Enum_13_0_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "13.0"
	Enum_13_1_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "13.1"
	Enum_13_2_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "13.2"
	Enum_13_3_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "13.3"
	Enum_13_4_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "13.4"
	Enum_14_0_ProjectUpdateV30AudienceIosOsv ProjectUpdateV30AudienceIosOsv = "14.0"
	Enum_4_0_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "4.0"
	Enum_4_1_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "4.1"
	Enum_4_2_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "4.2"
	Enum_4_3_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "4.3"
	Enum_5_0_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "5.0"
	Enum_5_1_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "5.1"
	Enum_6_0_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "6.0"
	Enum_7_0_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "7.0"
	Enum_7_1_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "7.1"
	Enum_8_0_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "8.0"
	Enum_8_1_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "8.1"
	Enum_8_2_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "8.2"
	Enum_9_0_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "9.0"
	Enum_9_1_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "9.1"
	Enum_9_2_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "9.2"
	Enum_9_3_ProjectUpdateV30AudienceIosOsv  ProjectUpdateV30AudienceIosOsv = "9.3"
	NONE_ProjectUpdateV30AudienceIosOsv      ProjectUpdateV30AudienceIosOsv = "NONE"
)

// All allowed values of ProjectUpdateV30AudienceIosOsv enum
var AllowedProjectUpdateV30AudienceIosOsvEnumValues = []ProjectUpdateV30AudienceIosOsv{
	"0.0",
	"10.1",
	"10.2",
	"10.3",
	"11.0",
	"11.1",
	"11.2",
	"11.3",
	"11.4",
	"12.0",
	"12.1",
	"12.2",
	"12.3",
	"12.4",
	"13.0",
	"13.1",
	"13.2",
	"13.3",
	"13.4",
	"14.0",
	"4.0",
	"4.1",
	"4.2",
	"4.3",
	"5.0",
	"5.1",
	"6.0",
	"7.0",
	"7.1",
	"8.0",
	"8.1",
	"8.2",
	"9.0",
	"9.1",
	"9.2",
	"9.3",
	"NONE",
}

// NewProjectUpdateV30AudienceIosOsvFromValue returns a pointer to a valid ProjectUpdateV30AudienceIosOsv
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceIosOsvFromValue(v string) (*ProjectUpdateV30AudienceIosOsv, error) {
	ev := ProjectUpdateV30AudienceIosOsv(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceIosOsv: valid values are %v", v, AllowedProjectUpdateV30AudienceIosOsvEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceIosOsv) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceIosOsvEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_ios_osv value
func (v ProjectUpdateV30AudienceIosOsv) Ptr() *ProjectUpdateV30AudienceIosOsv {
	return &v
}
