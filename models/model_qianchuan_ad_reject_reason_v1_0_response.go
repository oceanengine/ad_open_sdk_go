/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRejectReasonV10Response struct for QianchuanAdRejectReasonV10Response
type QianchuanAdRejectReasonV10Response struct {
	//
	Code *int64                                  `json:"code,omitempty"`
	Data *QianchuanAdRejectReasonV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
