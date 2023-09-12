/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30TrackUrlSettingSendType
type ProjectCreateV30TrackUrlSettingSendType string

// List of project_create_v3.0_track_url_setting_send_type
const (
	CLIENT_SEND_ProjectCreateV30TrackUrlSettingSendType ProjectCreateV30TrackUrlSettingSendType = "CLIENT_SEND"
	SERVER_SEND_ProjectCreateV30TrackUrlSettingSendType ProjectCreateV30TrackUrlSettingSendType = "SERVER_SEND"
)

// All allowed values of ProjectCreateV30TrackUrlSettingSendType enum
var AllowedProjectCreateV30TrackUrlSettingSendTypeEnumValues = []ProjectCreateV30TrackUrlSettingSendType{
	"CLIENT_SEND",
	"SERVER_SEND",
}

// NewProjectCreateV30TrackUrlSettingSendTypeFromValue returns a pointer to a valid ProjectCreateV30TrackUrlSettingSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30TrackUrlSettingSendTypeFromValue(v string) (*ProjectCreateV30TrackUrlSettingSendType, error) {
	ev := ProjectCreateV30TrackUrlSettingSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30TrackUrlSettingSendType: valid values are %v", v, AllowedProjectCreateV30TrackUrlSettingSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30TrackUrlSettingSendType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30TrackUrlSettingSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_track_url_setting_send_type value
func (v ProjectCreateV30TrackUrlSettingSendType) Ptr() *ProjectCreateV30TrackUrlSettingSendType {
	return &v
}
