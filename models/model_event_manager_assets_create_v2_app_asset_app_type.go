/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAssetsCreateV2AppAssetAppType
type EventManagerAssetsCreateV2AppAssetAppType string

// List of event_manager_assets_create_v2_app_asset_app_type
const (
	ANDROID_EventManagerAssetsCreateV2AppAssetAppType EventManagerAssetsCreateV2AppAssetAppType = "Android"
	IOS_EventManagerAssetsCreateV2AppAssetAppType     EventManagerAssetsCreateV2AppAssetAppType = "IOS"
)

// Ptr returns reference to event_manager_assets_create_v2_app_asset_app_type value
func (v EventManagerAssetsCreateV2AppAssetAppType) Ptr() *EventManagerAssetsCreateV2AppAssetAppType {
	return &v
}
