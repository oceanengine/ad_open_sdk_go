/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListKeywordsMatchType
type ProjectListV30DataListKeywordsMatchType string

// List of project_list_v3.0_data_list_keywords_match_type
const (
	EXTENSIVE_ProjectListV30DataListKeywordsMatchType ProjectListV30DataListKeywordsMatchType = "EXTENSIVE"
	PHRASE_ProjectListV30DataListKeywordsMatchType    ProjectListV30DataListKeywordsMatchType = "PHRASE"
	PRECISION_ProjectListV30DataListKeywordsMatchType ProjectListV30DataListKeywordsMatchType = "PRECISION"
)

// All allowed values of ProjectListV30DataListKeywordsMatchType enum
var AllowedProjectListV30DataListKeywordsMatchTypeEnumValues = []ProjectListV30DataListKeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewProjectListV30DataListKeywordsMatchTypeFromValue returns a pointer to a valid ProjectListV30DataListKeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListKeywordsMatchTypeFromValue(v string) (*ProjectListV30DataListKeywordsMatchType, error) {
	ev := ProjectListV30DataListKeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListKeywordsMatchType: valid values are %v", v, AllowedProjectListV30DataListKeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListKeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListKeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_keywords_match_type value
func (v ProjectListV30DataListKeywordsMatchType) Ptr() *ProjectListV30DataListKeywordsMatchType {
	return &v
}
