/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentGetV30FilteringEmotionType
type ToolsCommentGetV30FilteringEmotionType string

// List of tools_comment_get_v3.0_filtering_emotion_type
const (
	NEGATIVE_ToolsCommentGetV30FilteringEmotionType ToolsCommentGetV30FilteringEmotionType = "NEGATIVE"
	NEUTRAL_ToolsCommentGetV30FilteringEmotionType  ToolsCommentGetV30FilteringEmotionType = "NEUTRAL"
	POSITIVE_ToolsCommentGetV30FilteringEmotionType ToolsCommentGetV30FilteringEmotionType = "POSITIVE"
)

// Ptr returns reference to tools_comment_get_v3.0_filtering_emotion_type value
func (v ToolsCommentGetV30FilteringEmotionType) Ptr() *ToolsCommentGetV30FilteringEmotionType {
	return &v
}
