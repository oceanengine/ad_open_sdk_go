/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EbpAdvertiserTaskListV2ResponseDataListInner struct for EbpAdvertiserTaskListV2ResponseDataListInner
type EbpAdvertiserTaskListV2ResponseDataListInner struct {
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	ErrMsg *string `json:"err_msg,omitempty"`
	//
	TaskId     *int64                                     `json:"task_id,omitempty"`
	TaskStatus *EbpAdvertiserTaskListV2DataListTaskStatus `json:"task_status,omitempty"`
}
