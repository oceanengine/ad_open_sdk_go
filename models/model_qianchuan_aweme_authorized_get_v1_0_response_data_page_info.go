/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeAuthorizedGetV10ResponseDataPageInfo
type QianchuanAwemeAuthorizedGetV10ResponseDataPageInfo struct {
	//
	Page int64 `json:"page"`
	//
	PageSize int64 `json:"page_size"`
	//
	TotalNumber *int64 `json:"total_number,omitempty"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
}
