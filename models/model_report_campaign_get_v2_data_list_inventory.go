/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2DataListInventory
type ReportCampaignGetV2DataListInventory string

// List of report_campaign_get_v2_data_list_inventory
const (
	INVENTORY_PIPIXIA_ReportCampaignGetV2DataListInventory      ReportCampaignGetV2DataListInventory = "INVENTORY_PIPIXIA"
	INVENTORY_UNIVERSAL_ReportCampaignGetV2DataListInventory    ReportCampaignGetV2DataListInventory = "INVENTORY_UNIVERSAL"
	INVENTORY_SEARCH_ReportCampaignGetV2DataListInventory       ReportCampaignGetV2DataListInventory = "INVENTORY_SEARCH"
	INVENTORY_AUTOMOBILE_ReportCampaignGetV2DataListInventory   ReportCampaignGetV2DataListInventory = "INVENTORY_AUTOMOBILE"
	INVENTORY_UNION_SLOT_ReportCampaignGetV2DataListInventory   ReportCampaignGetV2DataListInventory = "INVENTORY_UNION_SLOT"
	INVENTORY_AWEME_FEED_ReportCampaignGetV2DataListInventory   ReportCampaignGetV2DataListInventory = "INVENTORY_AWEME_FEED"
	INVENTORY_HOTSOON_FEED_ReportCampaignGetV2DataListInventory ReportCampaignGetV2DataListInventory = "INVENTORY_HOTSOON_FEED"
	INVENTORY_FEED_ReportCampaignGetV2DataListInventory         ReportCampaignGetV2DataListInventory = "INVENTORY_FEED"
	INVENTORY_STUDY_ReportCampaignGetV2DataListInventory        ReportCampaignGetV2DataListInventory = "INVENTORY_STUDY"
	INVENTORY_VIDEO_FEED_ReportCampaignGetV2DataListInventory   ReportCampaignGetV2DataListInventory = "INVENTORY_VIDEO_FEED"
	INVENTORY_TOMATO_NOVEL_ReportCampaignGetV2DataListInventory ReportCampaignGetV2DataListInventory = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_FACE_U_ReportCampaignGetV2DataListInventory       ReportCampaignGetV2DataListInventory = "INVENTORY_FACE_U"
	INVENTORY_BEAUTY_ReportCampaignGetV2DataListInventory       ReportCampaignGetV2DataListInventory = "INVENTORY_BEAUTY"
	INVENTORY_FURNISH_ReportCampaignGetV2DataListInventory      ReportCampaignGetV2DataListInventory = "INVENTORY_FURNISH"
	UNION_BOUTIQUE_GAME_ReportCampaignGetV2DataListInventory    ReportCampaignGetV2DataListInventory = "UNION_BOUTIQUE_GAME"
)

// All allowed values of ReportCampaignGetV2DataListInventory enum
var AllowedReportCampaignGetV2DataListInventoryEnumValues = []ReportCampaignGetV2DataListInventory{
	"INVENTORY_PIPIXIA",
	"INVENTORY_UNIVERSAL",
	"INVENTORY_SEARCH",
	"INVENTORY_AUTOMOBILE",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_AWEME_FEED",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_FEED",
	"INVENTORY_STUDY",
	"INVENTORY_VIDEO_FEED",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_FACE_U",
	"INVENTORY_BEAUTY",
	"INVENTORY_FURNISH",
	"UNION_BOUTIQUE_GAME",
}

// NewReportCampaignGetV2DataListInventoryFromValue returns a pointer to a valid ReportCampaignGetV2DataListInventory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListInventoryFromValue(v string) (*ReportCampaignGetV2DataListInventory, error) {
	ev := ReportCampaignGetV2DataListInventory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListInventory: valid values are %v", v, AllowedReportCampaignGetV2DataListInventoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListInventory) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListInventoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_inventory value
func (v ReportCampaignGetV2DataListInventory) Ptr() *ReportCampaignGetV2DataListInventory {
	return &v
}
