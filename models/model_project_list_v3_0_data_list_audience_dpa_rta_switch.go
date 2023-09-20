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

// ProjectListV30DataListAudienceDpaRtaSwitch
type ProjectListV30DataListAudienceDpaRtaSwitch string

// List of project_list_v3.0_data_list_audience_dpa_rta_switch
const (
	OFF_ProjectListV30DataListAudienceDpaRtaSwitch ProjectListV30DataListAudienceDpaRtaSwitch = "OFF"
	ON_ProjectListV30DataListAudienceDpaRtaSwitch  ProjectListV30DataListAudienceDpaRtaSwitch = "ON"
)

// All allowed values of ProjectListV30DataListAudienceDpaRtaSwitch enum
var AllowedProjectListV30DataListAudienceDpaRtaSwitchEnumValues = []ProjectListV30DataListAudienceDpaRtaSwitch{
	"OFF",
	"ON",
}

// NewProjectListV30DataListAudienceDpaRtaSwitchFromValue returns a pointer to a valid ProjectListV30DataListAudienceDpaRtaSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceDpaRtaSwitchFromValue(v string) (*ProjectListV30DataListAudienceDpaRtaSwitch, error) {
	ev := ProjectListV30DataListAudienceDpaRtaSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceDpaRtaSwitch: valid values are %v", v, AllowedProjectListV30DataListAudienceDpaRtaSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceDpaRtaSwitch) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceDpaRtaSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_dpa_rta_switch value
func (v ProjectListV30DataListAudienceDpaRtaSwitch) Ptr() *ProjectListV30DataListAudienceDpaRtaSwitch {
	return &v
}
