/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoUploadTaskListV2ResponseDataListInner struct for FileVideoUploadTaskListV2ResponseDataListInner
type FileVideoUploadTaskListV2ResponseDataListInner struct {
	//
	CreateTime *string `json:"create_time,omitempty"`
	// 当任务失败后，会返回失败信息
	ErrorMsg *string                                  `json:"error_msg,omitempty"`
	Status   *FileVideoUploadTaskListV2DataListStatus `json:"status,omitempty"`
	//
	TaskId    *int64                                                   `json:"task_id,omitempty"`
	VideoInfo *FileVideoUploadTaskListV2ResponseDataListInnerVideoInfo `json:"video_info,omitempty"`
}
