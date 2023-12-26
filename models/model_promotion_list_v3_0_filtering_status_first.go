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

// PromotionListV30FilteringStatusFirst
type PromotionListV30FilteringStatusFirst string

// List of promotion_list_v3.0_filtering_status_first
const (
	PROMOTION_STATUS_DELETED_PromotionListV30FilteringStatusFirst PromotionListV30FilteringStatusFirst = "PROMOTION_STATUS_DELETED"
	PROMOTION_STATUS_DISABLE_PromotionListV30FilteringStatusFirst PromotionListV30FilteringStatusFirst = "PROMOTION_STATUS_DISABLE"
	PROMOTION_STATUS_DONE_PromotionListV30FilteringStatusFirst    PromotionListV30FilteringStatusFirst = "PROMOTION_STATUS_DONE"
	PROMOTION_STATUS_ENABLE_PromotionListV30FilteringStatusFirst  PromotionListV30FilteringStatusFirst = "PROMOTION_STATUS_ENABLE"
	PROMOTION_STATUS_FROZEN_PromotionListV30FilteringStatusFirst  PromotionListV30FilteringStatusFirst = "PROMOTION_STATUS_FROZEN"
)

// All allowed values of PromotionListV30FilteringStatusFirst enum
var AllowedPromotionListV30FilteringStatusFirstEnumValues = []PromotionListV30FilteringStatusFirst{
	"PROMOTION_STATUS_DELETED",
	"PROMOTION_STATUS_DISABLE",
	"PROMOTION_STATUS_DONE",
	"PROMOTION_STATUS_ENABLE",
	"PROMOTION_STATUS_FROZEN",
}

// NewPromotionListV30FilteringStatusFirstFromValue returns a pointer to a valid PromotionListV30FilteringStatusFirst
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringStatusFirstFromValue(v string) (*PromotionListV30FilteringStatusFirst, error) {
	ev := PromotionListV30FilteringStatusFirst(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringStatusFirst: valid values are %v", v, AllowedPromotionListV30FilteringStatusFirstEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringStatusFirst) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringStatusFirstEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_status_first value
func (v PromotionListV30FilteringStatusFirst) Ptr() *PromotionListV30FilteringStatusFirst {
	return &v
}
