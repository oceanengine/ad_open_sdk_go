/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30Response struct for DouplusOrderListV30Response
type DouplusOrderListV30Response struct {
	//
	Code *int64                           `json:"code,omitempty"`
	Data *DouplusOrderListV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
