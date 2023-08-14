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

// ProjectListV30DataListDownloadMode
type ProjectListV30DataListDownloadMode string

// List of project_list_v3.0_data_list_download_mode
const (
	APP_STORE_DELIVERY_ProjectListV30DataListDownloadMode ProjectListV30DataListDownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_ProjectListV30DataListDownloadMode            ProjectListV30DataListDownloadMode = "DEFAULT"
)

// All allowed values of ProjectListV30DataListDownloadMode enum
var AllowedProjectListV30DataListDownloadModeEnumValues = []ProjectListV30DataListDownloadMode{
	"APP_STORE_DELIVERY",
	"DEFAULT",
}

// NewProjectListV30DataListDownloadModeFromValue returns a pointer to a valid ProjectListV30DataListDownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDownloadModeFromValue(v string) (*ProjectListV30DataListDownloadMode, error) {
	ev := ProjectListV30DataListDownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDownloadMode: valid values are %v", v, AllowedProjectListV30DataListDownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDownloadMode) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_download_mode value
func (v ProjectListV30DataListDownloadMode) Ptr() *ProjectListV30DataListDownloadMode {
	return &v
}
