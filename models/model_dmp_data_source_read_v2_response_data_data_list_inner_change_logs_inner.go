/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpDataSourceReadV2ResponseDataDataListInnerChangeLogsInner struct for DmpDataSourceReadV2ResponseDataDataListInnerChangeLogsInner
type DmpDataSourceReadV2ResponseDataDataListInnerChangeLogsInner struct {
	//
	Action *int64 `json:"action,omitempty"`
	//
	ChangeLogId *int64 `json:"change_log_id,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	FilePaths []string `json:"file_paths,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	PublishedTime *string `json:"published_time,omitempty"`
	//
	Status *int64 `json:"status,omitempty"`
}
