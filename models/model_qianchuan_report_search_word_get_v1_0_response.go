/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportSearchWordGetV10Response struct for QianchuanReportSearchWordGetV10Response
type QianchuanReportSearchWordGetV10Response struct {
	//
	Code *int64                                       `json:"code,omitempty"`
	Data *QianchuanReportSearchWordGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
