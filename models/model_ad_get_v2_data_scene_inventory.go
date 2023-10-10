/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataSceneInventory
type AdGetV2DataSceneInventory string

// List of ad_get_v2_data_scene_inventory
const (
	VIDEO_SCENE_AdGetV2DataSceneInventory AdGetV2DataSceneInventory = "VIDEO_SCENE"
	FEED_SCENE_AdGetV2DataSceneInventory  AdGetV2DataSceneInventory = "FEED_SCENE"
	NOT_SELECT_AdGetV2DataSceneInventory  AdGetV2DataSceneInventory = "NOT_SELECT"
	TAIL_SCENE_AdGetV2DataSceneInventory  AdGetV2DataSceneInventory = "TAIL_SCENE"
)

// All allowed values of AdGetV2DataSceneInventory enum
var AllowedAdGetV2DataSceneInventoryEnumValues = []AdGetV2DataSceneInventory{
	"VIDEO_SCENE",
	"FEED_SCENE",
	"NOT_SELECT",
	"TAIL_SCENE",
}

// NewAdGetV2DataSceneInventoryFromValue returns a pointer to a valid AdGetV2DataSceneInventory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataSceneInventoryFromValue(v string) (*AdGetV2DataSceneInventory, error) {
	ev := AdGetV2DataSceneInventory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataSceneInventory: valid values are %v", v, AllowedAdGetV2DataSceneInventoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataSceneInventory) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataSceneInventoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_scene_inventory value
func (v AdGetV2DataSceneInventory) Ptr() *AdGetV2DataSceneInventory {
	return &v
}
