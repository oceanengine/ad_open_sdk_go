/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFinanceDetailGetV10Response struct for QianchuanFinanceDetailGetV10Response
type QianchuanFinanceDetailGetV10Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *QianchuanFinanceDetailGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
