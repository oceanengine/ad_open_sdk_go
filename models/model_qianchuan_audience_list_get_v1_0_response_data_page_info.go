/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceListGetV10ResponseDataPageInfo
type QianchuanAudienceListGetV10ResponseDataPageInfo struct {
	//
	Page *int32 `json:"page,omitempty"`
	//
	PageSize *int32 `json:"page_size,omitempty"`
	//
	TotalNumber *int64 `json:"total_number,omitempty"`
	//
	TotalPage *int32 `json:"total_page,omitempty"`
}
