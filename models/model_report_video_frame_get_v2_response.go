/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportVideoFrameGetV2Response struct for ReportVideoFrameGetV2Response
type ReportVideoFrameGetV2Response struct {
	//
	Code *int64                             `json:"code,omitempty"`
	Data *ReportVideoFrameGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
