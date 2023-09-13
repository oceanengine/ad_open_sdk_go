/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringPlatform
type ProjectListV30FilteringPlatform string

// List of project_list_v3.0_filtering_platform
const (
	ANDROID_ProjectListV30FilteringPlatform ProjectListV30FilteringPlatform = "ANDROID"
	IOS_ProjectListV30FilteringPlatform     ProjectListV30FilteringPlatform = "IOS"
)

// All allowed values of ProjectListV30FilteringPlatform enum
var AllowedProjectListV30FilteringPlatformEnumValues = []ProjectListV30FilteringPlatform{
	"ANDROID",
	"IOS",
}

// NewProjectListV30FilteringPlatformFromValue returns a pointer to a valid ProjectListV30FilteringPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringPlatformFromValue(v string) (*ProjectListV30FilteringPlatform, error) {
	ev := ProjectListV30FilteringPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringPlatform: valid values are %v", v, AllowedProjectListV30FilteringPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringPlatform) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_platform value
func (v ProjectListV30FilteringPlatform) Ptr() *ProjectListV30FilteringPlatform {
	return &v
}
