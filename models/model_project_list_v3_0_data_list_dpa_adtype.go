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

// ProjectListV30DataListDpaAdtype
type ProjectListV30DataListDpaAdtype string

// List of project_list_v3.0_data_list_dpa_adtype
const (
	DPA_APP_ProjectListV30DataListDpaAdtype  ProjectListV30DataListDpaAdtype = "DPA_APP"
	DPA_LINK_ProjectListV30DataListDpaAdtype ProjectListV30DataListDpaAdtype = "DPA_LINK"
)

// All allowed values of ProjectListV30DataListDpaAdtype enum
var AllowedProjectListV30DataListDpaAdtypeEnumValues = []ProjectListV30DataListDpaAdtype{
	"DPA_APP",
	"DPA_LINK",
}

// NewProjectListV30DataListDpaAdtypeFromValue returns a pointer to a valid ProjectListV30DataListDpaAdtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDpaAdtypeFromValue(v string) (*ProjectListV30DataListDpaAdtype, error) {
	ev := ProjectListV30DataListDpaAdtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDpaAdtype: valid values are %v", v, AllowedProjectListV30DataListDpaAdtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDpaAdtype) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDpaAdtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_dpa_adtype value
func (v ProjectListV30DataListDpaAdtype) Ptr() *ProjectListV30DataListDpaAdtype {
	return &v
}
