/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupDeleteV30Response struct for BudgetGroupDeleteV30Response
type BudgetGroupDeleteV30Response struct {
	//
	Code *int64                            `json:"code,omitempty"`
	Data *BudgetGroupDeleteV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
