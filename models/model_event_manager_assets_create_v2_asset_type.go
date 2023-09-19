/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerAssetsCreateV2AssetType
type EventManagerAssetsCreateV2AssetType string

// List of event_manager_assets_create_v2_asset_type
const (
	APP_EventManagerAssetsCreateV2AssetType             EventManagerAssetsCreateV2AssetType = "APP"
	TETRIS_EXTERNAL_EventManagerAssetsCreateV2AssetType EventManagerAssetsCreateV2AssetType = "TETRIS_EXTERNAL"
	THIRD_EXTERNAL_EventManagerAssetsCreateV2AssetType  EventManagerAssetsCreateV2AssetType = "THIRD_EXTERNAL"
	MINI_PROGRAME_EventManagerAssetsCreateV2AssetType   EventManagerAssetsCreateV2AssetType = "MINI_PROGRAME"
	OFFLINE_EVENT_EventManagerAssetsCreateV2AssetType   EventManagerAssetsCreateV2AssetType = "OFFLINE_EVENT"
	OTHER_EventManagerAssetsCreateV2AssetType           EventManagerAssetsCreateV2AssetType = "OTHER"
	QUICK_APP_EventManagerAssetsCreateV2AssetType       EventManagerAssetsCreateV2AssetType = "QUICK_APP"
)

// All allowed values of EventManagerAssetsCreateV2AssetType enum
var AllowedEventManagerAssetsCreateV2AssetTypeEnumValues = []EventManagerAssetsCreateV2AssetType{
	"APP",
	"TETRIS_EXTERNAL",
	"THIRD_EXTERNAL",
	"MINI_PROGRAME",
	"OFFLINE_EVENT",
	"OTHER",
	"QUICK_APP",
}

// NewEventManagerAssetsCreateV2AssetTypeFromValue returns a pointer to a valid EventManagerAssetsCreateV2AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerAssetsCreateV2AssetTypeFromValue(v string) (*EventManagerAssetsCreateV2AssetType, error) {
	ev := EventManagerAssetsCreateV2AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerAssetsCreateV2AssetType: valid values are %v", v, AllowedEventManagerAssetsCreateV2AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerAssetsCreateV2AssetType) IsValid() bool {
	for _, existing := range AllowedEventManagerAssetsCreateV2AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_assets_create_v2_asset_type value
func (v EventManagerAssetsCreateV2AssetType) Ptr() *EventManagerAssetsCreateV2AssetType {
	return &v
}
