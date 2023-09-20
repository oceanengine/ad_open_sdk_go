/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAssetsCreateV2RequestSiteAsset 橙子落地页信息
type EventManagerAssetsCreateV2RequestSiteAsset struct {
	// 橙子建站站点id，橙子落地页必填
	SiteId *int64 `json:"site_id,omitempty"`
	// 橙子建站站点名称，橙子落地页必填
	SiteName *string `json:"site_name,omitempty"`
}
