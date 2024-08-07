/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListKeywordsBidType
type ProjectListV30DataListKeywordsBidType string

// List of project_list_v3.0_data_list_keywords_bid_type
const (
	FEED_TO_SEARCH_ProjectListV30DataListKeywordsBidType ProjectListV30DataListKeywordsBidType = "FEED_TO_SEARCH"
)

// All allowed values of ProjectListV30DataListKeywordsBidType enum
var AllowedProjectListV30DataListKeywordsBidTypeEnumValues = []ProjectListV30DataListKeywordsBidType{
	"FEED_TO_SEARCH",
}

// NewProjectListV30DataListKeywordsBidTypeFromValue returns a pointer to a valid ProjectListV30DataListKeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListKeywordsBidTypeFromValue(v string) (*ProjectListV30DataListKeywordsBidType, error) {
	ev := ProjectListV30DataListKeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListKeywordsBidType: valid values are %v", v, AllowedProjectListV30DataListKeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListKeywordsBidType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListKeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_keywords_bid_type value
func (v ProjectListV30DataListKeywordsBidType) Ptr() *ProjectListV30DataListKeywordsBidType {
	return &v
}
