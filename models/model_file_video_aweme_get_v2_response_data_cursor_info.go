/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAwemeGetV2ResponseDataCursorInfo 游标查询返回值，游标查询场景下以该返回值为准
type FileVideoAwemeGetV2ResponseDataCursorInfo struct {
	// 下一次游标查询的游标值
	Cursor *string `json:"cursor,omitempty"`
	// 游标查询场景下，是否还有更多数据
	HasMore *bool `json:"has_more,omitempty"`
}
