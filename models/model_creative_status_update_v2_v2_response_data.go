/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeStatusUpdateV2V2ResponseData
type CreativeStatusUpdateV2V2ResponseData struct {
	//
	Errors []*CreativeStatusUpdateV2V2ResponseDataErrorsInner `json:"errors,omitempty"`
	//
	Success []int64 `json:"success,omitempty"`
}
