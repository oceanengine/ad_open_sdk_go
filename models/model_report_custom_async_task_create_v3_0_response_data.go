/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomAsyncTaskCreateV30ResponseData
type ReportCustomAsyncTaskCreateV30ResponseData struct {
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	TaskId *int64 `json:"task_id,omitempty"`
	//
	TaskName   *string                                       `json:"task_name,omitempty"`
	TaskStatus *ReportCustomAsyncTaskCreateV30DataTaskStatus `json:"task_status,omitempty"`
}
