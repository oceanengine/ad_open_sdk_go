/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AsyncTaskGetV2Filtering
type AsyncTaskGetV2Filtering struct {
	//
	TaskIds []int64 `json:"task_ids,omitempty"`
	//
	TaskName *string `json:"task_name,omitempty"`
}
