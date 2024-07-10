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

// PromotionListV30DataListOptStatus
type PromotionListV30DataListOptStatus string

// List of promotion_list_v3.0_data_list_opt_status
const (
	DISABLE_PromotionListV30DataListOptStatus          PromotionListV30DataListOptStatus = "DISABLE"
	DISABLE_BY_QUOTA_PromotionListV30DataListOptStatus PromotionListV30DataListOptStatus = "DISABLE_BY_QUOTA"
	ENABLE_PromotionListV30DataListOptStatus           PromotionListV30DataListOptStatus = "ENABLE"
	FROZEN_PromotionListV30DataListOptStatus           PromotionListV30DataListOptStatus = "FROZEN"
)

// All allowed values of PromotionListV30DataListOptStatus enum
var AllowedPromotionListV30DataListOptStatusEnumValues = []PromotionListV30DataListOptStatus{
	"DISABLE",
	"DISABLE_BY_QUOTA",
	"ENABLE",
	"FROZEN",
}

// NewPromotionListV30DataListOptStatusFromValue returns a pointer to a valid PromotionListV30DataListOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListOptStatusFromValue(v string) (*PromotionListV30DataListOptStatus, error) {
	ev := PromotionListV30DataListOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListOptStatus: valid values are %v", v, AllowedPromotionListV30DataListOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListOptStatus) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_opt_status value
func (v PromotionListV30DataListOptStatus) Ptr() *PromotionListV30DataListOptStatus {
	return &v
}
