/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAccountBudgetGetV10Response struct for QianchuanAccountBudgetGetV10Response
type QianchuanAccountBudgetGetV10Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *QianchuanAccountBudgetGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
