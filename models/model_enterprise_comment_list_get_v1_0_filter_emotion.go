/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentListGetV10FilterEmotion
type EnterpriseCommentListGetV10FilterEmotion string

// List of enterprise_comment_list_get_v1.0_filter_emotion
const (
	NEU_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "NEU_EMOTION"
	ANY_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "ANY_EMOTION"
	NEG_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "NEG_EMOTION"
	POS_EMOTION_EnterpriseCommentListGetV10FilterEmotion EnterpriseCommentListGetV10FilterEmotion = "POS_EMOTION"
)

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_emotion value
func (v EnterpriseCommentListGetV10FilterEmotion) Ptr() *EnterpriseCommentListGetV10FilterEmotion {
	return &v
}
