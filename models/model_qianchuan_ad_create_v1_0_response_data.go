/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10ResponseData
type QianchuanAdCreateV10ResponseData struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	NoticeInfos []*QianchuanAdCreateV10ResponseDataNoticeInfosInner `json:"notice_infos,omitempty"`
}
