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

// ProjectListV30DataListAudiencePlatform
type ProjectListV30DataListAudiencePlatform string

// List of project_list_v3.0_data_list_audience_platform
const (
	ANDROID_ProjectListV30DataListAudiencePlatform ProjectListV30DataListAudiencePlatform = "ANDROID"
	IOS_ProjectListV30DataListAudiencePlatform     ProjectListV30DataListAudiencePlatform = "IOS"
	IPAD_ProjectListV30DataListAudiencePlatform    ProjectListV30DataListAudiencePlatform = "IPAD"
	NONE_ProjectListV30DataListAudiencePlatform    ProjectListV30DataListAudiencePlatform = "NONE"
	PC_ProjectListV30DataListAudiencePlatform      ProjectListV30DataListAudiencePlatform = "PC"
	WAP_ProjectListV30DataListAudiencePlatform     ProjectListV30DataListAudiencePlatform = "WAP"
)

// All allowed values of ProjectListV30DataListAudiencePlatform enum
var AllowedProjectListV30DataListAudiencePlatformEnumValues = []ProjectListV30DataListAudiencePlatform{
	"ANDROID",
	"IOS",
	"IPAD",
	"NONE",
	"PC",
	"WAP",
}

// NewProjectListV30DataListAudiencePlatformFromValue returns a pointer to a valid ProjectListV30DataListAudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudiencePlatformFromValue(v string) (*ProjectListV30DataListAudiencePlatform, error) {
	ev := ProjectListV30DataListAudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudiencePlatform: valid values are %v", v, AllowedProjectListV30DataListAudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudiencePlatform) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_platform value
func (v ProjectListV30DataListAudiencePlatform) Ptr() *ProjectListV30DataListAudiencePlatform {
	return &v
}
