/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderQuotaGetV10Response struct for QianchuanAwemeOrderQuotaGetV10Response
type QianchuanAwemeOrderQuotaGetV10Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *QianchuanAwemeOrderQuotaGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
