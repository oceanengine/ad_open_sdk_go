/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDeliveryRangeInventoryType
type ProjectListV30DataListDeliveryRangeInventoryType string

// List of project_list_v3.0_data_list_delivery_range_inventory_type
const (
	INVENTORY_AUTOMOBILE_ProjectListV30DataListDeliveryRangeInventoryType      ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_AUTOMOBILE"
	INVENTORY_AWEME_FEED_ProjectListV30DataListDeliveryRangeInventoryType      ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_AWEME_FEED"
	INVENTORY_BEAUTY_ProjectListV30DataListDeliveryRangeInventoryType          ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_BEAUTY"
	INVENTORY_FACE_U_ProjectListV30DataListDeliveryRangeInventoryType          ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_FACE_U"
	INVENTORY_FEED_ProjectListV30DataListDeliveryRangeInventoryType            ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_FEED"
	INVENTORY_HOMED_AGGREGATE_ProjectListV30DataListDeliveryRangeInventoryType ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_HOMED_AGGREGATE"
	INVENTORY_HOTSOON_FEED_ProjectListV30DataListDeliveryRangeInventoryType    ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_HOTSOON_FEED"
	INVENTORY_PIPIXIA_ProjectListV30DataListDeliveryRangeInventoryType         ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_PIPIXIA"
	INVENTORY_SEARCH_ProjectListV30DataListDeliveryRangeInventoryType          ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_SEARCH"
	INVENTORY_STUDY_ProjectListV30DataListDeliveryRangeInventoryType           ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_STUDY"
	INVENTORY_TOMATO_NOVEL_ProjectListV30DataListDeliveryRangeInventoryType    ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_ProjectListV30DataListDeliveryRangeInventoryType      ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_UNION_SLOT"
	INVENTORY_UNIVERSAL_ProjectListV30DataListDeliveryRangeInventoryType       ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_UNIVERSAL"
	INVENTORY_VIDEO_FEED_ProjectListV30DataListDeliveryRangeInventoryType      ProjectListV30DataListDeliveryRangeInventoryType = "INVENTORY_VIDEO_FEED"
	UNION_BOUTIQUE_GAME_ProjectListV30DataListDeliveryRangeInventoryType       ProjectListV30DataListDeliveryRangeInventoryType = "UNION_BOUTIQUE_GAME"
)

// All allowed values of ProjectListV30DataListDeliveryRangeInventoryType enum
var AllowedProjectListV30DataListDeliveryRangeInventoryTypeEnumValues = []ProjectListV30DataListDeliveryRangeInventoryType{
	"INVENTORY_AUTOMOBILE",
	"INVENTORY_AWEME_FEED",
	"INVENTORY_BEAUTY",
	"INVENTORY_FACE_U",
	"INVENTORY_FEED",
	"INVENTORY_HOMED_AGGREGATE",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_PIPIXIA",
	"INVENTORY_SEARCH",
	"INVENTORY_STUDY",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_UNIVERSAL",
	"INVENTORY_VIDEO_FEED",
	"UNION_BOUTIQUE_GAME",
}

// NewProjectListV30DataListDeliveryRangeInventoryTypeFromValue returns a pointer to a valid ProjectListV30DataListDeliveryRangeInventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliveryRangeInventoryTypeFromValue(v string) (*ProjectListV30DataListDeliveryRangeInventoryType, error) {
	ev := ProjectListV30DataListDeliveryRangeInventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliveryRangeInventoryType: valid values are %v", v, AllowedProjectListV30DataListDeliveryRangeInventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliveryRangeInventoryType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliveryRangeInventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_range_inventory_type value
func (v ProjectListV30DataListDeliveryRangeInventoryType) Ptr() *ProjectListV30DataListDeliveryRangeInventoryType {
	return &v
}
