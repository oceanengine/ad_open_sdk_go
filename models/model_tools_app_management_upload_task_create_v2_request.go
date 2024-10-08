/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementUploadTaskCreateV2Request struct for ToolsAppManagementUploadTaskCreateV2Request
type ToolsAppManagementUploadTaskCreateV2Request struct {
	// 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountId   int64                                           `json:"account_id"`
	AccountType ToolsAppManagementUploadTaskCreateV2AccountType `json:"account_type"`
	// 上传文件下载链接，长度最长255
	DownloadUrl string                                       `json:"download_url"`
	FileType    ToolsAppManagementUploadTaskCreateV2FileType `json:"file_type"`
	// 文件md5，用于文件校验
	Md5 string `json:"md5"`
}
