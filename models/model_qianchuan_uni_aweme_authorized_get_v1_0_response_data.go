/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAuthorizedGetV10ResponseData
type QianchuanUniAwemeAuthorizedGetV10ResponseData struct {
	//
	AwemeIdList []*QianchuanUniAwemeAuthorizedGetV10ResponseDataAwemeIdListInner `json:"aweme_id_list,omitempty"`
	PageInfo    QianchuanUniAwemeAuthorizedGetV10ResponseDataPageInfo            `json:"page_info"`
}
