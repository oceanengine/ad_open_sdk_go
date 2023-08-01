/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30DownloadMode
type ProjectUpdateV30DownloadMode string

// List of project_update_v3.0_download_mode
const (
	APP_STORE_DELIVERY_ProjectUpdateV30DownloadMode ProjectUpdateV30DownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_ProjectUpdateV30DownloadMode            ProjectUpdateV30DownloadMode = "DEFAULT"
)

// All allowed values of ProjectUpdateV30DownloadMode enum
var AllowedProjectUpdateV30DownloadModeEnumValues = []ProjectUpdateV30DownloadMode{
	"APP_STORE_DELIVERY",
	"DEFAULT",
}

// NewProjectUpdateV30DownloadModeFromValue returns a pointer to a valid ProjectUpdateV30DownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30DownloadModeFromValue(v string) (*ProjectUpdateV30DownloadMode, error) {
	ev := ProjectUpdateV30DownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30DownloadMode: valid values are %v", v, AllowedProjectUpdateV30DownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30DownloadMode) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30DownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_download_mode value
func (v ProjectUpdateV30DownloadMode) Ptr() *ProjectUpdateV30DownloadMode {
	return &v
}
