/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10ResponseDataNoticeInfosInner struct for QianchuanAdUpdateV10ResponseDataNoticeInfosInner
type QianchuanAdUpdateV10ResponseDataNoticeInfosInner struct {
	//
	Code *int64 `json:"code,omitempty"`
	//
	Message            *string                                                             `json:"message,omitempty"`
	SearchKeywordError *QianchuanAdUpdateV10ResponseDataNoticeInfosInnerSearchKeywordError `json:"search_keyword_error,omitempty"`
}
