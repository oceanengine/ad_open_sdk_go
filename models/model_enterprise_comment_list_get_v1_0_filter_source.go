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

// EnterpriseCommentListGetV10FilterSource
type EnterpriseCommentListGetV10FilterSource string

// List of enterprise_comment_list_get_v1.0_filter_source
const (
	FROM_NATURAL_EnterpriseCommentListGetV10FilterSource EnterpriseCommentListGetV10FilterSource = "FROM_NATURAL"
	FROM_BRAND_EnterpriseCommentListGetV10FilterSource   EnterpriseCommentListGetV10FilterSource = "FROM_BRAND"
	FROM_PERFORM_EnterpriseCommentListGetV10FilterSource EnterpriseCommentListGetV10FilterSource = "FROM_PERFORM"
	FROM_OTHER_EnterpriseCommentListGetV10FilterSource   EnterpriseCommentListGetV10FilterSource = "FROM_OTHER"
	FROM_DOUPLUS_EnterpriseCommentListGetV10FilterSource EnterpriseCommentListGetV10FilterSource = "FROM_DOUPLUS"
)

// All allowed values of EnterpriseCommentListGetV10FilterSource enum
var AllowedEnterpriseCommentListGetV10FilterSourceEnumValues = []EnterpriseCommentListGetV10FilterSource{
	"FROM_NATURAL",
	"FROM_BRAND",
	"FROM_PERFORM",
	"FROM_OTHER",
	"FROM_DOUPLUS",
}

// NewEnterpriseCommentListGetV10FilterSourceFromValue returns a pointer to a valid EnterpriseCommentListGetV10FilterSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10FilterSourceFromValue(v string) (*EnterpriseCommentListGetV10FilterSource, error) {
	ev := EnterpriseCommentListGetV10FilterSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10FilterSource: valid values are %v", v, AllowedEnterpriseCommentListGetV10FilterSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10FilterSource) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10FilterSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_source value
func (v EnterpriseCommentListGetV10FilterSource) Ptr() *EnterpriseCommentListGetV10FilterSource {
	return &v
}
