/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAndroidAppListV2Filtering
type ToolsAppManagementAndroidAppListV2Filtering struct {
	CreateTime  *ToolsAppManagementAndroidAppListV2FilteringCreateTime  `json:"create_time,omitempty"`
	PublishTime *ToolsAppManagementAndroidAppListV2FilteringPublishTime `json:"publish_time,omitempty"`
	//
	SearchKey  *string                                                `json:"search_key,omitempty"`
	SearchType *ToolsAppManagementAndroidAppListV2FilteringSearchType `json:"search_type,omitempty"`
	Status     *ToolsAppManagementAndroidAppListV2FilteringStatus     `json:"status,omitempty"`
}
