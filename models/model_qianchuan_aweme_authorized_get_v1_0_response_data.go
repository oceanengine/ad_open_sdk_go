/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeAuthorizedGetV10ResponseData
type QianchuanAwemeAuthorizedGetV10ResponseData struct {
	//
	AwemeIdList []*QianchuanAwemeAuthorizedGetV10ResponseDataAwemeIdListInner `json:"aweme_id_list,omitempty"`
	PageInfo    *QianchuanAwemeAuthorizedGetV10ResponseDataPageInfo           `json:"page_info,omitempty"`
}
