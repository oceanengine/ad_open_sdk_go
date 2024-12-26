/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetInfoV2ResponseDataOrderListInnerOrderItemInfoAuthorListInner struct for StarOrderGetInfoV2ResponseDataOrderListInnerOrderItemInfoAuthorListInner
type StarOrderGetInfoV2ResponseDataOrderListInnerOrderItemInfoAuthorListInner struct {
	// 达人ID（星图中的达人Star ID）
	AuthorId *int64 `json:"author_id,omitempty"`
	// 任务数量
	Cnt *int64 `json:"cnt,omitempty"`
	// 任务类型枚举 (1)：1-20s视频； (2)：21-60s视频； (71)：60s+视频
	VideoType *int64 `json:"video_type,omitempty"`
}