/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementUploadTaskListV2ResponseDataListInner struct for ToolsAppManagementUploadTaskListV2ResponseDataListInner
type ToolsAppManagementUploadTaskListV2ResponseDataListInner struct {
	// 任务信息
	Message *string                                           `json:"message,omitempty"`
	Status  *ToolsAppManagementUploadTaskListV2DataListStatus `json:"status,omitempty"`
	// 上传文件id
	UploadId *int64 `json:"upload_id,omitempty"`
}
