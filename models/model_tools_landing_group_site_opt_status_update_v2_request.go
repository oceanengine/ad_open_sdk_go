/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupSiteOptStatusUpdateV2Request struct for ToolsLandingGroupSiteOptStatusUpdateV2Request
type ToolsLandingGroupSiteOptStatusUpdateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*ToolsLandingGroupSiteOptStatusUpdateV2RequestDataInner `json:"data"`
	// 落地页组 ID
	GroupId int64 `json:"group_id"`
}
