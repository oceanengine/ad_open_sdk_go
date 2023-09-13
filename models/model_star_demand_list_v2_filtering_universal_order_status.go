/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandListV2FilteringUniversalOrderStatus
type StarDemandListV2FilteringUniversalOrderStatus string

// List of star_demand_list_v2_filtering_universal_order_status
const (
	FINISHED_StarDemandListV2FilteringUniversalOrderStatus      StarDemandListV2FilteringUniversalOrderStatus = "FINISHED"
	CANCELED_StarDemandListV2FilteringUniversalOrderStatus      StarDemandListV2FilteringUniversalOrderStatus = "CANCELED"
	ALL_StarDemandListV2FilteringUniversalOrderStatus           StarDemandListV2FilteringUniversalOrderStatus = "ALL"
	WAIT_PAYMENT_StarDemandListV2FilteringUniversalOrderStatus  StarDemandListV2FilteringUniversalOrderStatus = "WAIT_PAYMENT"
	ONGOING_StarDemandListV2FilteringUniversalOrderStatus       StarDemandListV2FilteringUniversalOrderStatus = "ONGOING"
	RECEIVEING_StarDemandListV2FilteringUniversalOrderStatus    StarDemandListV2FilteringUniversalOrderStatus = "RECEIVEING"
	WAIT_EVALUATE_StarDemandListV2FilteringUniversalOrderStatus StarDemandListV2FilteringUniversalOrderStatus = "WAIT_EVALUATE"
)

// All allowed values of StarDemandListV2FilteringUniversalOrderStatus enum
var AllowedStarDemandListV2FilteringUniversalOrderStatusEnumValues = []StarDemandListV2FilteringUniversalOrderStatus{
	"FINISHED",
	"CANCELED",
	"ALL",
	"WAIT_PAYMENT",
	"ONGOING",
	"RECEIVEING",
	"WAIT_EVALUATE",
}

// NewStarDemandListV2FilteringUniversalOrderStatusFromValue returns a pointer to a valid StarDemandListV2FilteringUniversalOrderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandListV2FilteringUniversalOrderStatusFromValue(v string) (*StarDemandListV2FilteringUniversalOrderStatus, error) {
	ev := StarDemandListV2FilteringUniversalOrderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandListV2FilteringUniversalOrderStatus: valid values are %v", v, AllowedStarDemandListV2FilteringUniversalOrderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandListV2FilteringUniversalOrderStatus) IsValid() bool {
	for _, existing := range AllowedStarDemandListV2FilteringUniversalOrderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_list_v2_filtering_universal_order_status value
func (v StarDemandListV2FilteringUniversalOrderStatus) Ptr() *StarDemandListV2FilteringUniversalOrderStatus {
	return &v
}
