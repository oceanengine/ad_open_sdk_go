/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerVideoLocalVideo 本地视频内容
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerVideoLocalVideo struct {
	// 用户自行上传的图片url，当`local_video`不为空时必填
	PosterUrl *string `json:"poster_url,omitempty"`
	// 视频ID，本地视频上传后得到的视频云id，对应的是[【获取视频素材】]（https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710601820172）接口获取的视频ID，当`local_video`不为空时必填
	VideoId *string `json:"video_id,omitempty"`
}
