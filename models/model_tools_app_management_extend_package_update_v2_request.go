/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageUpdateV2Request struct for ToolsAppManagementExtendPackageUpdateV2Request
type ToolsAppManagementExtendPackageUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ChannelList []*ToolsAppManagementExtendPackageUpdateV2RequestChannelListInner `json:"channel_list,omitempty"`
	Mode        *ToolsAppManagementExtendPackageUpdateV2Mode                      `json:"mode,omitempty"`
	//
	PackageId string `json:"package_id"`
}
