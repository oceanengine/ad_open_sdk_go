/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderGetV10ResponseDataPageInfo
type QianchuanAwemeOrderGetV10ResponseDataPageInfo struct {
	// 拉取的视频个数
	Count *int64 `json:"count,omitempty"`
	// 下一次分页拉取的游标值
	Cursor  *int64                                        `json:"cursor,omitempty"`
	HasMore *QianchuanAwemeOrderGetV10DataPageInfoHasMore `json:"has_more,omitempty"`
}
