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

// ProjectCreateV30KeywordsMatchType
type ProjectCreateV30KeywordsMatchType string

// List of project_create_v3.0_keywords_match_type
const (
	EXTENSIVE_ProjectCreateV30KeywordsMatchType ProjectCreateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_ProjectCreateV30KeywordsMatchType    ProjectCreateV30KeywordsMatchType = "PHRASE"
	PRECISION_ProjectCreateV30KeywordsMatchType ProjectCreateV30KeywordsMatchType = "PRECISION"
)

// All allowed values of ProjectCreateV30KeywordsMatchType enum
var AllowedProjectCreateV30KeywordsMatchTypeEnumValues = []ProjectCreateV30KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewProjectCreateV30KeywordsMatchTypeFromValue returns a pointer to a valid ProjectCreateV30KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30KeywordsMatchTypeFromValue(v string) (*ProjectCreateV30KeywordsMatchType, error) {
	ev := ProjectCreateV30KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30KeywordsMatchType: valid values are %v", v, AllowedProjectCreateV30KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_keywords_match_type value
func (v ProjectCreateV30KeywordsMatchType) Ptr() *ProjectCreateV30KeywordsMatchType {
	return &v
}
