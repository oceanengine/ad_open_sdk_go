/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInnerVideoOnlineVideo 在线视频内容
type ToolsSiteTemplateSiteCreateV2RequestBricksInnerVideoOnlineVideo struct {
	// 视频原始地址，例如https://v.youku.com/v_show/id_xxx.html，当前仅支持优酷视频，当`online_video`不为空时，必填
	OriginUrl *string `json:"origin_url,omitempty"`
	// 用户自行上传的图片url，选填，不填充该字段时，会读取线上视频原封面
	PosterUrl *string `json:"poster_url,omitempty"`
}
