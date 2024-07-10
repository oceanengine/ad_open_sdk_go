/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarAttachmentUploadV2Request struct for StarAttachmentUploadV2Request
type StarAttachmentUploadV2Request struct {
	// 上传的文件
	File *FormFileInfo `json:"file"`
	// 文件名
	FileName string `json:"file_name"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
