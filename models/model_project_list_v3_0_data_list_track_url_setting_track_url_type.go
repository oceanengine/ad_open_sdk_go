/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListTrackUrlSettingTrackUrlType
type ProjectListV30DataListTrackUrlSettingTrackUrlType string

// List of project_list_v3.0_data_list_track_url_setting_track_url_type
const (
	CUSTOM_ProjectListV30DataListTrackUrlSettingTrackUrlType   ProjectListV30DataListTrackUrlSettingTrackUrlType = "CUSTOM"
	GROUP_ID_ProjectListV30DataListTrackUrlSettingTrackUrlType ProjectListV30DataListTrackUrlSettingTrackUrlType = "GROUP_ID"
)

// All allowed values of ProjectListV30DataListTrackUrlSettingTrackUrlType enum
var AllowedProjectListV30DataListTrackUrlSettingTrackUrlTypeEnumValues = []ProjectListV30DataListTrackUrlSettingTrackUrlType{
	"CUSTOM",
	"GROUP_ID",
}

// NewProjectListV30DataListTrackUrlSettingTrackUrlTypeFromValue returns a pointer to a valid ProjectListV30DataListTrackUrlSettingTrackUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListTrackUrlSettingTrackUrlTypeFromValue(v string) (*ProjectListV30DataListTrackUrlSettingTrackUrlType, error) {
	ev := ProjectListV30DataListTrackUrlSettingTrackUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListTrackUrlSettingTrackUrlType: valid values are %v", v, AllowedProjectListV30DataListTrackUrlSettingTrackUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListTrackUrlSettingTrackUrlType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListTrackUrlSettingTrackUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_track_url_setting_track_url_type value
func (v ProjectListV30DataListTrackUrlSettingTrackUrlType) Ptr() *ProjectListV30DataListTrackUrlSettingTrackUrlType {
	return &v
}
