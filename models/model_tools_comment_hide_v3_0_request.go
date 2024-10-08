/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentHideV30Request struct for ToolsCommentHideV30Request
type ToolsCommentHideV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 评论id列表，只允许传入小于等于20个评论ID。
	CommentIds []int64 `json:"comment_ids"`
}
