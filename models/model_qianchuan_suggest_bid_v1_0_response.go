/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanSuggestBidV10Response struct for QianchuanSuggestBidV10Response
type QianchuanSuggestBidV10Response struct {
	//
	Code *int64                              `json:"code,omitempty"`
	Data *QianchuanSuggestBidV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
