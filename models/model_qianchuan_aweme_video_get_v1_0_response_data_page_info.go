/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeVideoGetV10ResponseDataPageInfo
type QianchuanAwemeVideoGetV10ResponseDataPageInfo struct {
	//  下一次分页拉取的游标值
	Cursor  *int64                                        `json:"cursor,omitempty"`
	HasMore *QianchuanAwemeVideoGetV10DataPageInfoHasMore `json:"has_more,omitempty"`
}
