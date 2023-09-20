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

// ProjectListV30FilteringLandingType
type ProjectListV30FilteringLandingType string

// List of project_list_v3.0_filtering_landing_type
const (
	APP_ProjectListV30FilteringLandingType           ProjectListV30FilteringLandingType = "APP"
	DPA_ProjectListV30FilteringLandingType           ProjectListV30FilteringLandingType = "DPA"
	LINK_ProjectListV30FilteringLandingType          ProjectListV30FilteringLandingType = "LINK"
	MICRO_GAME_ProjectListV30FilteringLandingType    ProjectListV30FilteringLandingType = "MICRO_GAME"
	NATIVE_ACTION_ProjectListV30FilteringLandingType ProjectListV30FilteringLandingType = "NATIVE_ACTION"
	QUICK_APP_ProjectListV30FilteringLandingType     ProjectListV30FilteringLandingType = "QUICK_APP"
	SHOP_ProjectListV30FilteringLandingType          ProjectListV30FilteringLandingType = "SHOP"
)

// All allowed values of ProjectListV30FilteringLandingType enum
var AllowedProjectListV30FilteringLandingTypeEnumValues = []ProjectListV30FilteringLandingType{
	"APP",
	"DPA",
	"LINK",
	"MICRO_GAME",
	"NATIVE_ACTION",
	"QUICK_APP",
	"SHOP",
}

// NewProjectListV30FilteringLandingTypeFromValue returns a pointer to a valid ProjectListV30FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringLandingTypeFromValue(v string) (*ProjectListV30FilteringLandingType, error) {
	ev := ProjectListV30FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringLandingType: valid values are %v", v, AllowedProjectListV30FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_landing_type value
func (v ProjectListV30FilteringLandingType) Ptr() *ProjectListV30FilteringLandingType {
	return &v
}
