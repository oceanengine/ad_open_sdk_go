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

// EnterpriseCommentListGetV10DataCommentListLevel
type EnterpriseCommentListGetV10DataCommentListLevel string

// List of enterprise_comment_list_get_v1.0_data_comment_list_level
const (
	LEVEL_ALL_EnterpriseCommentListGetV10DataCommentListLevel EnterpriseCommentListGetV10DataCommentListLevel = "LEVEL_ALL"
	LEVEL_TWO_EnterpriseCommentListGetV10DataCommentListLevel EnterpriseCommentListGetV10DataCommentListLevel = "LEVEL_TWO"
	LEVEL_ONE_EnterpriseCommentListGetV10DataCommentListLevel EnterpriseCommentListGetV10DataCommentListLevel = "LEVEL_ONE"
)

// All allowed values of EnterpriseCommentListGetV10DataCommentListLevel enum
var AllowedEnterpriseCommentListGetV10DataCommentListLevelEnumValues = []EnterpriseCommentListGetV10DataCommentListLevel{
	"LEVEL_ALL",
	"LEVEL_TWO",
	"LEVEL_ONE",
}

// NewEnterpriseCommentListGetV10DataCommentListLevelFromValue returns a pointer to a valid EnterpriseCommentListGetV10DataCommentListLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10DataCommentListLevelFromValue(v string) (*EnterpriseCommentListGetV10DataCommentListLevel, error) {
	ev := EnterpriseCommentListGetV10DataCommentListLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10DataCommentListLevel: valid values are %v", v, AllowedEnterpriseCommentListGetV10DataCommentListLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10DataCommentListLevel) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10DataCommentListLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_data_comment_list_level value
func (v EnterpriseCommentListGetV10DataCommentListLevel) Ptr() *EnterpriseCommentListGetV10DataCommentListLevel {
	return &v
}
