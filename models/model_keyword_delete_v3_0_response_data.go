/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordDeleteV30ResponseData
type KeywordDeleteV30ResponseData struct {
	//
	ErrorList []*KeywordDeleteV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []int64 `json:"success_list,omitempty"`
}
