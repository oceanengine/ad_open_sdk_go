/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetVideoV2ResponseData
type StarOrderGetVideoV2ResponseData struct {
	// 任务视频信息
	OrderVideoList []*StarOrderGetVideoV2ResponseDataOrderVideoListInner `json:"order_video_list,omitempty"`
}
