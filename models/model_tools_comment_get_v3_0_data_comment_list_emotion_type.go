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

// ToolsCommentGetV30DataCommentListEmotionType
type ToolsCommentGetV30DataCommentListEmotionType string

// List of tools_comment_get_v3.0_data_comment_list_emotion_type
const (
	NEGATIVE_ToolsCommentGetV30DataCommentListEmotionType ToolsCommentGetV30DataCommentListEmotionType = "NEGATIVE"
	NEUTRAL_ToolsCommentGetV30DataCommentListEmotionType  ToolsCommentGetV30DataCommentListEmotionType = "NEUTRAL"
	POSITIVE_ToolsCommentGetV30DataCommentListEmotionType ToolsCommentGetV30DataCommentListEmotionType = "POSITIVE"
)

// All allowed values of ToolsCommentGetV30DataCommentListEmotionType enum
var AllowedToolsCommentGetV30DataCommentListEmotionTypeEnumValues = []ToolsCommentGetV30DataCommentListEmotionType{
	"NEGATIVE",
	"NEUTRAL",
	"POSITIVE",
}

// NewToolsCommentGetV30DataCommentListEmotionTypeFromValue returns a pointer to a valid ToolsCommentGetV30DataCommentListEmotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentGetV30DataCommentListEmotionTypeFromValue(v string) (*ToolsCommentGetV30DataCommentListEmotionType, error) {
	ev := ToolsCommentGetV30DataCommentListEmotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentGetV30DataCommentListEmotionType: valid values are %v", v, AllowedToolsCommentGetV30DataCommentListEmotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentGetV30DataCommentListEmotionType) IsValid() bool {
	for _, existing := range AllowedToolsCommentGetV30DataCommentListEmotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_get_v3.0_data_comment_list_emotion_type value
func (v ToolsCommentGetV30DataCommentListEmotionType) Ptr() *ToolsCommentGetV30DataCommentListEmotionType {
	return &v
}
