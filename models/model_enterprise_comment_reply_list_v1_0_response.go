/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentReplyListV10Response struct for EnterpriseCommentReplyListV10Response
type EnterpriseCommentReplyListV10Response struct {
	//
	Code *int64                                     `json:"code,omitempty"`
	Data *EnterpriseCommentReplyListV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
