/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMixcutTaskResultGetV30ResponseDataVideoInfosInner struct for AicMixcutTaskResultGetV30ResponseDataVideoInfosInner
type AicMixcutTaskResultGetV30ResponseDataVideoInfosInner struct {
	// 视频ID，可以调用推送的 MAPI 推送到对应的广告账号。
	VideoId *string `json:"video_id,omitempty"`
	// 视频名称
	VideoName *string `json:"video_name,omitempty"`
	// 视频预览链接，可利用此链接批量下载视频
	VideoUrl *string `json:"video_url,omitempty"`
}
