/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCarouselAwemeGetV10ResponseDataPageInfo
type QianchuanCarouselAwemeGetV10ResponseDataPageInfo struct {
	//
	Count *int64 `json:"count,omitempty"`
	//
	Cursor  *int64                                           `json:"cursor,omitempty"`
	HasMore *QianchuanCarouselAwemeGetV10DataPageInfoHasMore `json:"has_more,omitempty"`
}
