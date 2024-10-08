/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAndroidBasicPackageGetV2ResponseDataCurrentVersion 当前线上版本信息
type ToolsAppManagementAndroidBasicPackageGetV2ResponseDataCurrentVersion struct {
	// 应用名
	AppName *string `json:"app_name,omitempty"`
	// 审核id，审核通过时id显示为0
	AuditId *int64 `json:"audit_id,omitempty"`
	// 审核失败信息，审核通过显示空
	AuditMessage *string                                                             `json:"audit_message,omitempty"`
	Status       *ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus `json:"status,omitempty"`
	// 版本编码
	VersionCode *string `json:"version_code,omitempty"`
	// 版本名
	VersionName *string `json:"version_name,omitempty"`
}
