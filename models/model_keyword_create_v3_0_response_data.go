/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV30ResponseData
type KeywordCreateV30ResponseData struct {
	//
	ErrorList []*KeywordCreateV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	//
	SuccessList []*KeywordCreateV30ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
