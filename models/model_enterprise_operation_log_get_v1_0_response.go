/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseOperationLogGetV10Response struct for EnterpriseOperationLogGetV10Response
type EnterpriseOperationLogGetV10Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *EnterpriseOperationLogGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}