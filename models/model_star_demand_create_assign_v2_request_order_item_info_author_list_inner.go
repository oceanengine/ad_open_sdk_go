/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestOrderItemInfoAuthorListInner struct for StarDemandCreateAssignV2RequestOrderItemInfoAuthorListInner
type StarDemandCreateAssignV2RequestOrderItemInfoAuthorListInner struct {
	// 达人ID（星图中的达人Star ID）
	AuthorId int64 `json:"author_id"`
	// 任务数量  1到10的整数，所有达人任务累计数量不能超过50
	Cnt int64 `json:"cnt"`
	// 任务类型枚举  [video_type]仅限以下 (1)：1-20s视频； (2)：21-60s视频； (71)：60s+视频
	VideoType int64 `json:"video_type"`
}
