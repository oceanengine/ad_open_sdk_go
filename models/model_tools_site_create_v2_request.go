/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteCreateV2Request struct for ToolsSiteCreateV2Request
type ToolsSiteCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Bricks []*ToolsSiteCreateV2RequestBricksInner `json:"bricks"`
	//
	Name     string                     `json:"name"`
	SiteType *ToolsSiteCreateV2SiteType `json:"site_type,omitempty"`
}
