/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageCreateV2Request struct for ToolsAppManagementExtendPackageCreateV2Request
type ToolsAppManagementExtendPackageCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ChannelCount *int64 `json:"channel_count,omitempty"`
	//
	ChannelList []*ToolsAppManagementExtendPackageCreateV2RequestChannelListInner `json:"channel_list,omitempty"`
	Mode        *ToolsAppManagementExtendPackageCreateV2Mode                      `json:"mode,omitempty"`
	//
	PackageId string `json:"package_id"`
}
