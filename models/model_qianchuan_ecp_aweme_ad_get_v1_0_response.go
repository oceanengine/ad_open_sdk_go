/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanEcpAwemeAdGetV10Response struct for QianchuanEcpAwemeAdGetV10Response
type QianchuanEcpAwemeAdGetV10Response struct {
	//
	Code *int64                                 `json:"code,omitempty"`
	Data *QianchuanEcpAwemeAdGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
