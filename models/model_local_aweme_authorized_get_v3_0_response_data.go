/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalAwemeAuthorizedGetV30ResponseData
type LocalAwemeAuthorizedGetV30ResponseData struct {
	// 抖音号列表
	AwemeIdList []*LocalAwemeAuthorizedGetV30ResponseDataAwemeIdListInner `json:"aweme_id_list,omitempty"`
	PageInfo    *LocalAwemeAuthorizedGetV30ResponseDataPageInfo           `json:"page_info,omitempty"`
}
