/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoMaterialClearTaskGetV2ResponseDataListInner struct for FileVideoMaterialClearTaskGetV2ResponseDataListInner
type FileVideoMaterialClearTaskGetV2ResponseDataListInner struct {
	// 清理任务id
	ClearId         *int64                                                               `json:"clear_id,omitempty"`
	ClearTaskParams *FileVideoMaterialClearTaskGetV2ResponseDataListInnerClearTaskParams `json:"clear_task_params,omitempty"`
	// 任务创建时间
	CreateTime *string                                            `json:"create_time,omitempty"`
	TaskStatus *FileVideoMaterialClearTaskGetV2DataListTaskStatus `json:"task_status,omitempty"`
}
