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

// ProjectCreateV30LandingType
type ProjectCreateV30LandingType string

// List of project_create_v3.0_landing_type
const (
	APP_ProjectCreateV30LandingType           ProjectCreateV30LandingType = "APP"
	DPA_ProjectCreateV30LandingType           ProjectCreateV30LandingType = "DPA"
	LINK_ProjectCreateV30LandingType          ProjectCreateV30LandingType = "LINK"
	MICRO_GAME_ProjectCreateV30LandingType    ProjectCreateV30LandingType = "MICRO_GAME"
	NATIVE_ACTION_ProjectCreateV30LandingType ProjectCreateV30LandingType = "NATIVE_ACTION"
	QUICK_APP_ProjectCreateV30LandingType     ProjectCreateV30LandingType = "QUICK_APP"
	SHOP_ProjectCreateV30LandingType          ProjectCreateV30LandingType = "SHOP"
)

// All allowed values of ProjectCreateV30LandingType enum
var AllowedProjectCreateV30LandingTypeEnumValues = []ProjectCreateV30LandingType{
	"APP",
	"DPA",
	"LINK",
	"MICRO_GAME",
	"NATIVE_ACTION",
	"QUICK_APP",
	"SHOP",
}

// NewProjectCreateV30LandingTypeFromValue returns a pointer to a valid ProjectCreateV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30LandingTypeFromValue(v string) (*ProjectCreateV30LandingType, error) {
	ev := ProjectCreateV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30LandingType: valid values are %v", v, AllowedProjectCreateV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30LandingType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_landing_type value
func (v ProjectCreateV30LandingType) Ptr() *ProjectCreateV30LandingType {
	return &v
}
