/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackageUpdateV2Request struct for ToolsUnionFlowPackageUpdateV2Request
type ToolsUnionFlowPackageUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	FlowPackageId int64 `json:"flow_package_id"`
	//
	Name string `json:"name"`
	//
	Rit []int64 `json:"rit"`
}
