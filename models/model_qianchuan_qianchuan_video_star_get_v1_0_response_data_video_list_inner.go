/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanQianchuanVideoStarGetV10ResponseDataVideoListInner struct for QianchuanQianchuanVideoStarGetV10ResponseDataVideoListInner
type QianchuanQianchuanVideoStarGetV10ResponseDataVideoListInner struct {
	// 抖音号id
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	CommentCnt *int64 `json:"comment_cnt,omitempty"`
	//
	Duration *float64 `json:"duration,omitempty"`
	//
	Height *int64 `json:"height,omitempty"`
	//
	LikeCnt *int64 `json:"like_cnt,omitempty"`
	//
	ShareCnt *int64 `json:"share_cnt,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	Url *string `json:"url,omitempty"`
	//
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	ViewCnt *int64 `json:"view_cnt,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}
