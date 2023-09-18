/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAssetsCreateV2Request struct for EventManagerAssetsCreateV2Request
type EventManagerAssetsCreateV2Request struct {
	// 广告主ID
	AdvertiserId   int64                                            `json:"advertiser_id"`
	AppAsset       *EventManagerAssetsCreateV2RequestAppAsset       `json:"app_asset,omitempty"`
	AssetType      EventManagerAssetsCreateV2AssetType              `json:"asset_type"`
	QuickAppAsset  *EventManagerAssetsCreateV2RequestQuickAppAsset  `json:"quick_app_asset,omitempty"`
	SiteAsset      *EventManagerAssetsCreateV2RequestSiteAsset      `json:"site_asset,omitempty"`
	ThirdPartAsset *EventManagerAssetsCreateV2RequestThirdPartAsset `json:"third_part_asset,omitempty"`
}