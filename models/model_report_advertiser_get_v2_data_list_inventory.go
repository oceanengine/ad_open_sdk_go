/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2DataListInventory
type ReportAdvertiserGetV2DataListInventory string

// List of report_advertiser_get_v2_data_list_inventory
const (
	INVENTORY_AWEME_FEED_ReportAdvertiserGetV2DataListInventory   ReportAdvertiserGetV2DataListInventory = "INVENTORY_AWEME_FEED"
	INVENTORY_VIDEO_FEED_ReportAdvertiserGetV2DataListInventory   ReportAdvertiserGetV2DataListInventory = "INVENTORY_VIDEO_FEED"
	INVENTORY_BEAUTY_ReportAdvertiserGetV2DataListInventory       ReportAdvertiserGetV2DataListInventory = "INVENTORY_BEAUTY"
	INVENTORY_SEARCH_ReportAdvertiserGetV2DataListInventory       ReportAdvertiserGetV2DataListInventory = "INVENTORY_SEARCH"
	INVENTORY_FURNISH_ReportAdvertiserGetV2DataListInventory      ReportAdvertiserGetV2DataListInventory = "INVENTORY_FURNISH"
	INVENTORY_UNION_SLOT_ReportAdvertiserGetV2DataListInventory   ReportAdvertiserGetV2DataListInventory = "INVENTORY_UNION_SLOT"
	INVENTORY_FACE_U_ReportAdvertiserGetV2DataListInventory       ReportAdvertiserGetV2DataListInventory = "INVENTORY_FACE_U"
	INVENTORY_UNIVERSAL_ReportAdvertiserGetV2DataListInventory    ReportAdvertiserGetV2DataListInventory = "INVENTORY_UNIVERSAL"
	INVENTORY_HOTSOON_FEED_ReportAdvertiserGetV2DataListInventory ReportAdvertiserGetV2DataListInventory = "INVENTORY_HOTSOON_FEED"
	UNION_BOUTIQUE_GAME_ReportAdvertiserGetV2DataListInventory    ReportAdvertiserGetV2DataListInventory = "UNION_BOUTIQUE_GAME"
	INVENTORY_PIPIXIA_ReportAdvertiserGetV2DataListInventory      ReportAdvertiserGetV2DataListInventory = "INVENTORY_PIPIXIA"
	INVENTORY_AUTOMOBILE_ReportAdvertiserGetV2DataListInventory   ReportAdvertiserGetV2DataListInventory = "INVENTORY_AUTOMOBILE"
	INVENTORY_TOMATO_NOVEL_ReportAdvertiserGetV2DataListInventory ReportAdvertiserGetV2DataListInventory = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_FEED_ReportAdvertiserGetV2DataListInventory         ReportAdvertiserGetV2DataListInventory = "INVENTORY_FEED"
	INVENTORY_STUDY_ReportAdvertiserGetV2DataListInventory        ReportAdvertiserGetV2DataListInventory = "INVENTORY_STUDY"
)

// All allowed values of ReportAdvertiserGetV2DataListInventory enum
var AllowedReportAdvertiserGetV2DataListInventoryEnumValues = []ReportAdvertiserGetV2DataListInventory{
	"INVENTORY_AWEME_FEED",
	"INVENTORY_VIDEO_FEED",
	"INVENTORY_BEAUTY",
	"INVENTORY_SEARCH",
	"INVENTORY_FURNISH",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_FACE_U",
	"INVENTORY_UNIVERSAL",
	"INVENTORY_HOTSOON_FEED",
	"UNION_BOUTIQUE_GAME",
	"INVENTORY_PIPIXIA",
	"INVENTORY_AUTOMOBILE",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_FEED",
	"INVENTORY_STUDY",
}

// NewReportAdvertiserGetV2DataListInventoryFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListInventory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListInventoryFromValue(v string) (*ReportAdvertiserGetV2DataListInventory, error) {
	ev := ReportAdvertiserGetV2DataListInventory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListInventory: valid values are %v", v, AllowedReportAdvertiserGetV2DataListInventoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListInventory) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListInventoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_inventory value
func (v ReportAdvertiserGetV2DataListInventory) Ptr() *ReportAdvertiserGetV2DataListInventory {
	return &v
}
