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

// ProjectCreateV30DownloadMode
type ProjectCreateV30DownloadMode string

// List of project_create_v3.0_download_mode
const (
	APP_STORE_DELIVERY_ProjectCreateV30DownloadMode ProjectCreateV30DownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_ProjectCreateV30DownloadMode            ProjectCreateV30DownloadMode = "DEFAULT"
)

// All allowed values of ProjectCreateV30DownloadMode enum
var AllowedProjectCreateV30DownloadModeEnumValues = []ProjectCreateV30DownloadMode{
	"APP_STORE_DELIVERY",
	"DEFAULT",
}

// NewProjectCreateV30DownloadModeFromValue returns a pointer to a valid ProjectCreateV30DownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DownloadModeFromValue(v string) (*ProjectCreateV30DownloadMode, error) {
	ev := ProjectCreateV30DownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DownloadMode: valid values are %v", v, AllowedProjectCreateV30DownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DownloadMode) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_download_mode value
func (v ProjectCreateV30DownloadMode) Ptr() *ProjectCreateV30DownloadMode {
	return &v
}
