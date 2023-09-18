/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAwemeGetV2ResponseData
type FileVideoAwemeGetV2ResponseData struct {
	CursorInfo *FileVideoAwemeGetV2ResponseDataCursorInfo `json:"cursor_info,omitempty"`
	// 抖音视频列表
	List     []*FileVideoAwemeGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *FileVideoAwemeGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}