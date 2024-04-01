/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAndroidBasicPackageUpdateV2Request struct for ToolsAppManagementAndroidBasicPackageUpdateV2Request
type ToolsAppManagementAndroidBasicPackageUpdateV2Request struct {
	// 账户id指可以接的账号体系如广告主id、巨量纵横组织id等
	AccountId   int64                                                    `json:"account_id"`
	AccountType ToolsAppManagementAndroidBasicPackageUpdateV2AccountType `json:"account_type"`
	// 应用介绍
	AppDescription string `json:"app_description"`
	// 应用开发者名称
	AppDeveloperName string `json:"app_developer_name"`
	// 应用名称
	AppName string `json:"app_name"`
	// 是否自动发布
	AutoPublish bool                                                    `json:"auto_publish"`
	FileOption  ToolsAppManagementAndroidBasicPackageUpdateV2FileOption `json:"file_option"`
	// 文件内容<br>必须包含一个apk；可选：五个图片、一个视频
	Files []*ToolsAppManagementAndroidBasicPackageUpdateV2RequestFilesInner `json:"files"`
	// 分类id
	IndustryId string `json:"industry_id"`
	// 应用包id
	PackageId   string                                                   `json:"package_id"`
	PaymentType ToolsAppManagementAndroidBasicPackageUpdateV2PaymentType `json:"payment_type"`
	// 隐私权限说明
	PermissionsDescription string `json:"permissions_description"`
	// 推荐语
	Recommend string `json:"recommend"`
	// 应用题材标签id
	ThemeTagId *string `json:"theme_tag_id,omitempty"`
	// 版本更新说明
	UpdateDescription string `json:"update_description"`
}
