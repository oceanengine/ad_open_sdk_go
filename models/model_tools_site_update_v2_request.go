/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteUpdateV2Request struct for ToolsSiteUpdateV2Request
type ToolsSiteUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Bricks []*ToolsSiteUpdateV2RequestBricksInner `json:"bricks"`
	//
	Name *string `json:"name,omitempty"`
	//
	SiteId string `json:"site_id"`
}
