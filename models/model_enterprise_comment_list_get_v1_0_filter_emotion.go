/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseCommentListGetV10FilterEmotion
type EnterpriseCommentListGetV10FilterEmotion string

// List of enterprise_comment_list_get_v1.0_filter_emotion
const (
	ANY_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "ANY_EMOTION"
	NEU_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "NEU_EMOTION"
	POS_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "POS_EMOTION"
	NEG_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "NEG_EMOTION"
)

// All allowed values of EnterpriseCommentListGetV10FilterEmotion enum
var AllowedEnterpriseCommentListGetV10FilterEmotionEnumValues = []EnterpriseCommentListGetV10FilterEmotion{
	"ANY_EMOTION",
	"NEU_EMOTION",
	"POS_EMOTION",
	"NEG_EMOTION",
}

// NewEnterpriseCommentListGetV10FilterEmotionFromValue returns a pointer to a valid EnterpriseCommentListGetV10FilterEmotion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10FilterEmotionFromValue(v string) (*EnterpriseCommentListGetV10FilterEmotion, error) {
	ev := EnterpriseCommentListGetV10FilterEmotion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10FilterEmotion: valid values are %v", v, AllowedEnterpriseCommentListGetV10FilterEmotionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10FilterEmotion) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10FilterEmotionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_emotion value
func (v EnterpriseCommentListGetV10FilterEmotion) Ptr() *EnterpriseCommentListGetV10FilterEmotion {
	return &v
}
