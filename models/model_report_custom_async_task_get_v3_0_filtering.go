/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomAsyncTaskGetV30Filtering
type ReportCustomAsyncTaskGetV30Filtering struct {
	//
	DataTopics []*ReportCustomAsyncTaskGetV30FilteringDataTopics `json:"data_topics,omitempty"`
	//
	TaskIds []int64 `json:"task_ids,omitempty"`
	//
	TaskName *string `json:"task_name,omitempty"`
}
