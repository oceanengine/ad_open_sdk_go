/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus
type PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus string

// List of promotion_cost_protect_status_get_v3.0_data_compensate_status_info_list_compensate_status
const (
	DEFAULT_STATUS_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "DEFAULT_STATUS"
	IN_EFFECT_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus      PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "IN_EFFECT"
	INVALID_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus        PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "INVALID"
	CONFIRMING_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus     PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "CONFIRMING"
	PAID_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus           PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "PAID"
	ENDED_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus          PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus = "ENDED"
)

// All allowed values of PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus enum
var AllowedPromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusEnumValues = []PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus{
	"DEFAULT_STATUS",
	"IN_EFFECT",
	"INVALID",
	"CONFIRMING",
	"PAID",
	"ENDED",
}

// NewPromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusFromValue returns a pointer to a valid PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusFromValue(v string) (*PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus, error) {
	ev := PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus: valid values are %v", v, AllowedPromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus) IsValid() bool {
	for _, existing := range AllowedPromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_cost_protect_status_get_v3.0_data_compensate_status_info_list_compensate_status value
func (v PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus) Ptr() *PromotionCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus {
	return &v
}
