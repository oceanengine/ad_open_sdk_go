/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceDpaRtaSwitch
type ProjectCreateV30AudienceDpaRtaSwitch string

// List of project_create_v3.0_audience_dpa_rta_switch
const (
	OFF_ProjectCreateV30AudienceDpaRtaSwitch ProjectCreateV30AudienceDpaRtaSwitch = "OFF"
	ON_ProjectCreateV30AudienceDpaRtaSwitch  ProjectCreateV30AudienceDpaRtaSwitch = "ON"
)

// All allowed values of ProjectCreateV30AudienceDpaRtaSwitch enum
var AllowedProjectCreateV30AudienceDpaRtaSwitchEnumValues = []ProjectCreateV30AudienceDpaRtaSwitch{
	"OFF",
	"ON",
}

// NewProjectCreateV30AudienceDpaRtaSwitchFromValue returns a pointer to a valid ProjectCreateV30AudienceDpaRtaSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceDpaRtaSwitchFromValue(v string) (*ProjectCreateV30AudienceDpaRtaSwitch, error) {
	ev := ProjectCreateV30AudienceDpaRtaSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceDpaRtaSwitch: valid values are %v", v, AllowedProjectCreateV30AudienceDpaRtaSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceDpaRtaSwitch) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceDpaRtaSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_dpa_rta_switch value
func (v ProjectCreateV30AudienceDpaRtaSwitch) Ptr() *ProjectCreateV30AudienceDpaRtaSwitch {
	return &v
}
