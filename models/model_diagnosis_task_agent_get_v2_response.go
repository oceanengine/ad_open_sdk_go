/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DiagnosisTaskAgentGetV2Response struct for DiagnosisTaskAgentGetV2Response
type DiagnosisTaskAgentGetV2Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *DiagnosisTaskAgentGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
