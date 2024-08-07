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

// StarStarAdUniteTaskItemListV2DataStatInfoItemStatus
type StarStarAdUniteTaskItemListV2DataStatInfoItemStatus string

// List of star_star_ad_unite_task_item_list_v2_data_stat_info_item_status
const (
	AUDIT_PASS_StarStarAdUniteTaskItemListV2DataStatInfoItemStatus   StarStarAdUniteTaskItemListV2DataStatInfoItemStatus = "AUDIT_PASS"
	CREATED_StarStarAdUniteTaskItemListV2DataStatInfoItemStatus      StarStarAdUniteTaskItemListV2DataStatInfoItemStatus = "CREATED"
	PRIVATE_AREA_StarStarAdUniteTaskItemListV2DataStatInfoItemStatus StarStarAdUniteTaskItemListV2DataStatInfoItemStatus = "PRIVATE_AREA"
	USER_DELETED_StarStarAdUniteTaskItemListV2DataStatInfoItemStatus StarStarAdUniteTaskItemListV2DataStatInfoItemStatus = "USER_DELETED"
)

// All allowed values of StarStarAdUniteTaskItemListV2DataStatInfoItemStatus enum
var AllowedStarStarAdUniteTaskItemListV2DataStatInfoItemStatusEnumValues = []StarStarAdUniteTaskItemListV2DataStatInfoItemStatus{
	"AUDIT_PASS",
	"CREATED",
	"PRIVATE_AREA",
	"USER_DELETED",
}

// NewStarStarAdUniteTaskItemListV2DataStatInfoItemStatusFromValue returns a pointer to a valid StarStarAdUniteTaskItemListV2DataStatInfoItemStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarStarAdUniteTaskItemListV2DataStatInfoItemStatusFromValue(v string) (*StarStarAdUniteTaskItemListV2DataStatInfoItemStatus, error) {
	ev := StarStarAdUniteTaskItemListV2DataStatInfoItemStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarStarAdUniteTaskItemListV2DataStatInfoItemStatus: valid values are %v", v, AllowedStarStarAdUniteTaskItemListV2DataStatInfoItemStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarStarAdUniteTaskItemListV2DataStatInfoItemStatus) IsValid() bool {
	for _, existing := range AllowedStarStarAdUniteTaskItemListV2DataStatInfoItemStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_star_ad_unite_task_item_list_v2_data_stat_info_item_status value
func (v StarStarAdUniteTaskItemListV2DataStatInfoItemStatus) Ptr() *StarStarAdUniteTaskItemListV2DataStatInfoItemStatus {
	return &v
}
