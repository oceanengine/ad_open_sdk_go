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

// ProjectListV30DataListAudienceDpaCity
type ProjectListV30DataListAudienceDpaCity string

// List of project_list_v3.0_data_list_audience_dpa_city
const (
	OFF_ProjectListV30DataListAudienceDpaCity ProjectListV30DataListAudienceDpaCity = "OFF"
	ON_ProjectListV30DataListAudienceDpaCity  ProjectListV30DataListAudienceDpaCity = "ON"
)

// All allowed values of ProjectListV30DataListAudienceDpaCity enum
var AllowedProjectListV30DataListAudienceDpaCityEnumValues = []ProjectListV30DataListAudienceDpaCity{
	"OFF",
	"ON",
}

// NewProjectListV30DataListAudienceDpaCityFromValue returns a pointer to a valid ProjectListV30DataListAudienceDpaCity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceDpaCityFromValue(v string) (*ProjectListV30DataListAudienceDpaCity, error) {
	ev := ProjectListV30DataListAudienceDpaCity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceDpaCity: valid values are %v", v, AllowedProjectListV30DataListAudienceDpaCityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceDpaCity) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceDpaCityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_dpa_city value
func (v ProjectListV30DataListAudienceDpaCity) Ptr() *ProjectListV30DataListAudienceDpaCity {
	return &v
}
