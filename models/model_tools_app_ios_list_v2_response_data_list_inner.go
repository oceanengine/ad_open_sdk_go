/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppIosListV2ResponseDataListInner struct for ToolsAppIosListV2ResponseDataListInner
type ToolsAppIosListV2ResponseDataListInner struct {
	//
	AppName *string `json:"app_name,omitempty"`
	//
	DownloadUrl *string `json:"download_url,omitempty"`
	//
	IconUrl *string `json:"icon_url,omitempty"`
	//
	ItunesId *int64 `json:"itunes_id,omitempty"`
	//
	PackageId *int64 `json:"package_id,omitempty"`
	//
	PackageName *string                          `json:"package_name,omitempty"`
	Status      *ToolsAppIosListV2DataListStatus `json:"status,omitempty"`
	//
	Version *string `json:"version,omitempty"`
}
