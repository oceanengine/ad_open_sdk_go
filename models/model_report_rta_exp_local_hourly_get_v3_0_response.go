/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRtaExpLocalHourlyGetV30Response struct for ReportRtaExpLocalHourlyGetV30Response
type ReportRtaExpLocalHourlyGetV30Response struct {
	//
	Code *int64                                     `json:"code,omitempty"`
	Data *ReportRtaExpLocalHourlyGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
