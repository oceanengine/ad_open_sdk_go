/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseFlowCategoryGetV10Response struct for EnterpriseFlowCategoryGetV10Response
type EnterpriseFlowCategoryGetV10Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *EnterpriseFlowCategoryGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
