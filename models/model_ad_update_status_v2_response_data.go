/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdUpdateStatusV2ResponseData
type AdUpdateStatusV2ResponseData struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	Errors []*AdUpdateStatusV2ResponseDataErrorsInner `json:"errors,omitempty"`
}
