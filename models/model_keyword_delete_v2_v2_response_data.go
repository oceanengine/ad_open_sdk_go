/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordDeleteV2V2ResponseData
type KeywordDeleteV2V2ResponseData struct {
	//
	ErrorList []*KeywordDeleteV2V2ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []int64 `json:"success_list,omitempty"`
}
