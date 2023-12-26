/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBookingGetV2ResponseDataListInner struct for ToolsAppManagementBookingGetV2ResponseDataListInner
type ToolsAppManagementBookingGetV2ResponseDataListInner struct {
	//
	AppCloudId *int64 `json:"app_cloud_id,omitempty"`
	//
	AppName *string `json:"app_name,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	DownloadUrl *string `json:"download_url,omitempty"`
	//
	IconUrl *string `json:"icon_url,omitempty"`
	//
	PackageId *string `json:"package_id,omitempty"`
	//
	PackageName *string `json:"package_name,omitempty"`
	//
	PublishTime *string `json:"publish_time,omitempty"`
	//
	ScheduledPublishTime *string `json:"scheduled_publish_time,omitempty"`
	//
	UpdateTime *string `json:"update_time,omitempty"`
	//
	Version *string `json:"version,omitempty"`
}
