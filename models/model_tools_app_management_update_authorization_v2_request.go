/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementUpdateAuthorizationV2Request struct for ToolsAppManagementUpdateAuthorizationV2Request
type ToolsAppManagementUpdateAuthorizationV2Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 共享关系变更的广告主对象id，不允许为空，且数量不允许大于20个。
	AdvertiserIds []int64 `json:"advertiser_ids"`
	// 应用资产id
	BasicPackageId string                                               `json:"basic_package_id"`
	OperationType  ToolsAppManagementUpdateAuthorizationV2OperationType `json:"operation_type"`
}
