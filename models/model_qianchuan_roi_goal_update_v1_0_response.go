/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanRoiGoalUpdateV10Response struct for QianchuanRoiGoalUpdateV10Response
type QianchuanRoiGoalUpdateV10Response struct {
	//
	Code *int64                                 `json:"code,omitempty"`
	Data *QianchuanRoiGoalUpdateV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
