/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2FilteringInventoryTypes
type ReportCampaignGetV2FilteringInventoryTypes string

// List of report_campaign_get_v2_filtering_inventory_types
const (
	UNION_BOUTIQUE_GAME_ReportCampaignGetV2FilteringInventoryTypes    ReportCampaignGetV2FilteringInventoryTypes = "UNION_BOUTIQUE_GAME"
	INVENTORY_STUDY_ReportCampaignGetV2FilteringInventoryTypes        ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_STUDY"
	INVENTORY_PIPIXIA_ReportCampaignGetV2FilteringInventoryTypes      ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_PIPIXIA"
	INVENTORY_FACE_U_ReportCampaignGetV2FilteringInventoryTypes       ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_FACE_U"
	INVENTORY_FURNISH_ReportCampaignGetV2FilteringInventoryTypes      ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_FURNISH"
	INVENTORY_TOMATO_NOVEL_ReportCampaignGetV2FilteringInventoryTypes ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_AWEME_FEED_ReportCampaignGetV2FilteringInventoryTypes   ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_AWEME_FEED"
	INVENTORY_FEED_ReportCampaignGetV2FilteringInventoryTypes         ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_FEED"
	INVENTORY_SEARCH_ReportCampaignGetV2FilteringInventoryTypes       ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_SEARCH"
	INVENTORY_UNION_SLOT_ReportCampaignGetV2FilteringInventoryTypes   ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_UNION_SLOT"
	INVENTORY_BEAUTY_ReportCampaignGetV2FilteringInventoryTypes       ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_BEAUTY"
	INVENTORY_UNIVERSAL_ReportCampaignGetV2FilteringInventoryTypes    ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_UNIVERSAL"
	INVENTORY_HOTSOON_FEED_ReportCampaignGetV2FilteringInventoryTypes ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_HOTSOON_FEED"
	INVENTORY_VIDEO_FEED_ReportCampaignGetV2FilteringInventoryTypes   ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_VIDEO_FEED"
	INVENTORY_AUTOMOBILE_ReportCampaignGetV2FilteringInventoryTypes   ReportCampaignGetV2FilteringInventoryTypes = "INVENTORY_AUTOMOBILE"
)

// All allowed values of ReportCampaignGetV2FilteringInventoryTypes enum
var AllowedReportCampaignGetV2FilteringInventoryTypesEnumValues = []ReportCampaignGetV2FilteringInventoryTypes{
	"UNION_BOUTIQUE_GAME",
	"INVENTORY_STUDY",
	"INVENTORY_PIPIXIA",
	"INVENTORY_FACE_U",
	"INVENTORY_FURNISH",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_AWEME_FEED",
	"INVENTORY_FEED",
	"INVENTORY_SEARCH",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_BEAUTY",
	"INVENTORY_UNIVERSAL",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_VIDEO_FEED",
	"INVENTORY_AUTOMOBILE",
}

// NewReportCampaignGetV2FilteringInventoryTypesFromValue returns a pointer to a valid ReportCampaignGetV2FilteringInventoryTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringInventoryTypesFromValue(v string) (*ReportCampaignGetV2FilteringInventoryTypes, error) {
	ev := ReportCampaignGetV2FilteringInventoryTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringInventoryTypes: valid values are %v", v, AllowedReportCampaignGetV2FilteringInventoryTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringInventoryTypes) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringInventoryTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_inventory_types value
func (v ReportCampaignGetV2FilteringInventoryTypes) Ptr() *ReportCampaignGetV2FilteringInventoryTypes {
	return &v
}
