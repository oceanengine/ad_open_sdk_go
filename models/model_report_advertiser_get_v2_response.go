/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdvertiserGetV2Response struct for ReportAdvertiserGetV2Response
type ReportAdvertiserGetV2Response struct {
	//
	Code *int64                             `json:"code,omitempty"`
	Data *ReportAdvertiserGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}