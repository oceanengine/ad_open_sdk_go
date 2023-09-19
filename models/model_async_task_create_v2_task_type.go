/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AsyncTaskCreateV2TaskType
type AsyncTaskCreateV2TaskType string

// List of async_task_create_v2_task_type
const (
	REPORT_AsyncTaskCreateV2TaskType         AsyncTaskCreateV2TaskType = "REPORT"
	REPORT_DPA_AsyncTaskCreateV2TaskType     AsyncTaskCreateV2TaskType = "REPORT_DPA"
	REPORT_BIDWORD_AsyncTaskCreateV2TaskType AsyncTaskCreateV2TaskType = "REPORT_BIDWORD"
)

// All allowed values of AsyncTaskCreateV2TaskType enum
var AllowedAsyncTaskCreateV2TaskTypeEnumValues = []AsyncTaskCreateV2TaskType{
	"REPORT",
	"REPORT_DPA",
	"REPORT_BIDWORD",
}

// NewAsyncTaskCreateV2TaskTypeFromValue returns a pointer to a valid AsyncTaskCreateV2TaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAsyncTaskCreateV2TaskTypeFromValue(v string) (*AsyncTaskCreateV2TaskType, error) {
	ev := AsyncTaskCreateV2TaskType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AsyncTaskCreateV2TaskType: valid values are %v", v, AllowedAsyncTaskCreateV2TaskTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AsyncTaskCreateV2TaskType) IsValid() bool {
	for _, existing := range AllowedAsyncTaskCreateV2TaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to async_task_create_v2_task_type value
func (v AsyncTaskCreateV2TaskType) Ptr() *AsyncTaskCreateV2TaskType {
	return &v
}
