/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30TrackUrlSettingSendType
type ProjectUpdateV30TrackUrlSettingSendType string

// List of project_update_v3.0_track_url_setting_send_type
const (
	CLIENT_SEND_ProjectUpdateV30TrackUrlSettingSendType ProjectUpdateV30TrackUrlSettingSendType = "CLIENT_SEND"
	SERVER_SEND_ProjectUpdateV30TrackUrlSettingSendType ProjectUpdateV30TrackUrlSettingSendType = "SERVER_SEND"
)

// All allowed values of ProjectUpdateV30TrackUrlSettingSendType enum
var AllowedProjectUpdateV30TrackUrlSettingSendTypeEnumValues = []ProjectUpdateV30TrackUrlSettingSendType{
	"CLIENT_SEND",
	"SERVER_SEND",
}

// NewProjectUpdateV30TrackUrlSettingSendTypeFromValue returns a pointer to a valid ProjectUpdateV30TrackUrlSettingSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30TrackUrlSettingSendTypeFromValue(v string) (*ProjectUpdateV30TrackUrlSettingSendType, error) {
	ev := ProjectUpdateV30TrackUrlSettingSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30TrackUrlSettingSendType: valid values are %v", v, AllowedProjectUpdateV30TrackUrlSettingSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30TrackUrlSettingSendType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30TrackUrlSettingSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_track_url_setting_send_type value
func (v ProjectUpdateV30TrackUrlSettingSendType) Ptr() *ProjectUpdateV30TrackUrlSettingSendType {
	return &v
}
