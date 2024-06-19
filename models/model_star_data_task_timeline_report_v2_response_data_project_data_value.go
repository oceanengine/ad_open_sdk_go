/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDataTaskTimelineReportV2ResponseDataProjectDataValue struct for StarDataTaskTimelineReportV2ResponseDataProjectDataValue
type StarDataTaskTimelineReportV2ResponseDataProjectDataValue struct {
	// 评论次数
	CommentCnt *int64 `json:"comment_cnt,omitempty"`
	// 点赞次数
	LikeCnt *int64 `json:"like_cnt,omitempty"`
	// 播放次数
	PlayCnt *int64 `json:"play_cnt,omitempty"`
	// 分享次数
	ShareCnt *int64 `json:"share_cnt,omitempty"`
}
