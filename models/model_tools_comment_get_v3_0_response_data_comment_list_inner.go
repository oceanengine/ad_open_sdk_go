/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentGetV30ResponseDataCommentListInner struct for ToolsCommentGetV30ResponseDataCommentListInner
type ToolsCommentGetV30ResponseDataCommentListInner struct {
	// 计划id，仅1.0平台适用，引流流量会返回，自然流量不会返回
	AdId *int64 `json:"ad_id,omitempty"`
	// 计划名称
	AdName *string `json:"ad_name,omitempty"`
	// 评论用户抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 评论用户昵称
	AwemeName *string `json:"aweme_name,omitempty"`
	// 评论ID
	CommentId *int64 `json:"comment_id,omitempty"`
	// 评论创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 广告创意id，1.0：引流流量会返回，自然流量不会返回； 2.0：可能返回，返回值不具参考意义
	CreativeId  *int64                                        `json:"creative_id,omitempty"`
	EmotionType *ToolsCommentGetV30DataCommentListEmotionType `json:"emotion_type,omitempty"`
	HideStatus  *ToolsCommentGetV30DataCommentListHideStatus  `json:"hide_status,omitempty"`
	// 该评论是否置顶，0：表示不置顶，1：表示置顶
	IsStick *int32 `json:"is_stick,omitempty"`
	// 广告视频ID
	ItemId *int64 `json:"item_id,omitempty"`
	// 视频标题
	ItemTitle *string                                     `json:"item_title,omitempty"`
	LevelType *ToolsCommentGetV30DataCommentListLevelType `json:"level_type,omitempty"`
	// 点赞数
	LikeCount *int64 `json:"like_count,omitempty"`
	// 广告id，仅2.0平台适用，仅引流流量返回
	PromotionId *int64 `json:"promotion_id,omitempty"`
	// 评论的回复数量
	ReplyCount  *int64                                        `json:"reply_count,omitempty"`
	ReplyStatus *ToolsCommentGetV30DataCommentListReplyStatus `json:"reply_status,omitempty"`
	// 评论内容
	Text *string `json:"text,omitempty"`
}
