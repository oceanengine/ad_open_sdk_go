/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsGrayV10ResponseData
type QianchuanToolsGrayV10ResponseData struct {
	//
	ErrorList []string `json:"error_list,omitempty"`
	//
	SuccessList []*QianchuanToolsGrayV10ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
