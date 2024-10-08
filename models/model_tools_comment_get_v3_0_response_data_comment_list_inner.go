/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentGetV30ResponseDataCommentListInner struct for ToolsCommentGetV30ResponseDataCommentListInner
type ToolsCommentGetV30ResponseDataCommentListInner struct {
	// 评论用户抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 评论用户昵称
	AwemeName *string `json:"aweme_name,omitempty"`
	// 评论ID
	CommentId *int64 `json:"comment_id,omitempty"`
	// 评论图片内容，返回图片url列表
	CommentImages     []string                                            `json:"comment_images,omitempty"`
	CommentPermission *ToolsCommentGetV30DataCommentListCommentPermission `json:"comment_permission,omitempty"`
	CommentType       *ToolsCommentGetV30DataCommentListCommentType       `json:"comment_type,omitempty"`
	// 评论创建时间
	CreateTime  *string                                       `json:"create_time,omitempty"`
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
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	// 广告id，仅2.0平台适用，仅引流流量返回
	PromotionId *int64 `json:"promotion_id,omitempty"`
	// 评论的回复数量
	ReplyCount *int64 `json:"reply_count,omitempty"`
	// 评论内容
	Text *string `json:"text,omitempty"`
}
