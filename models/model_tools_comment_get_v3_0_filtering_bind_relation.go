/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCommentGetV30FilteringBindRelation
type ToolsCommentGetV30FilteringBindRelation string

// List of tools_comment_get_v3.0_filtering_bind_relation
const (
	ALL_ToolsCommentGetV30FilteringBindRelation           ToolsCommentGetV30FilteringBindRelation = "ALL"
	BIND_AWEME_ToolsCommentGetV30FilteringBindRelation    ToolsCommentGetV30FilteringBindRelation = "BIND_AWEME"
	VIRTUAL_AWEME_ToolsCommentGetV30FilteringBindRelation ToolsCommentGetV30FilteringBindRelation = "VIRTUAL_AWEME"
)

// All allowed values of ToolsCommentGetV30FilteringBindRelation enum
var AllowedToolsCommentGetV30FilteringBindRelationEnumValues = []ToolsCommentGetV30FilteringBindRelation{
	"ALL",
	"BIND_AWEME",
	"VIRTUAL_AWEME",
}

// NewToolsCommentGetV30FilteringBindRelationFromValue returns a pointer to a valid ToolsCommentGetV30FilteringBindRelation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30FilteringBindRelationFromValue(v string) (*ToolsCommentGetV30FilteringBindRelation, error) {
	ev := ToolsCommentGetV30FilteringBindRelation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30FilteringBindRelation: valid values are %v", v, AllowedToolsCommentGetV30FilteringBindRelationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30FilteringBindRelation) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30FilteringBindRelationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_filtering_bind_relation value
func (v ToolsCommentGetV30FilteringBindRelation) Ptr() *ToolsCommentGetV30FilteringBindRelation {
	return &v
}
