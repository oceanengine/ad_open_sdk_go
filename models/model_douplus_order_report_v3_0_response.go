/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderReportV30Response struct for DouplusOrderReportV30Response
type DouplusOrderReportV30Response struct {
	//
	Code *int64                             `json:"code,omitempty"`
	Data *DouplusOrderReportV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
