/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectAddV30Request struct for ToolsPrivativeWordProjectAddV30Request
type ToolsPrivativeWordProjectAddV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 项目列表
	ProjectList []*ToolsPrivativeWordProjectAddV30RequestProjectListInner `json:"project_list"`
}
