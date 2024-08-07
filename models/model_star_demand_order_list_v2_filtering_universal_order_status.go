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

// StarDemandOrderListV2FilteringUniversalOrderStatus
type StarDemandOrderListV2FilteringUniversalOrderStatus string

// List of star_demand_order_list_v2_filtering_universal_order_status
const (
	ALL_StarDemandOrderListV2FilteringUniversalOrderStatus           StarDemandOrderListV2FilteringUniversalOrderStatus = "ALL"
	CANCELED_StarDemandOrderListV2FilteringUniversalOrderStatus      StarDemandOrderListV2FilteringUniversalOrderStatus = "CANCELED"
	FINISHED_StarDemandOrderListV2FilteringUniversalOrderStatus      StarDemandOrderListV2FilteringUniversalOrderStatus = "FINISHED"
	ONGOING_StarDemandOrderListV2FilteringUniversalOrderStatus       StarDemandOrderListV2FilteringUniversalOrderStatus = "ONGOING"
	RECEIVEING_StarDemandOrderListV2FilteringUniversalOrderStatus    StarDemandOrderListV2FilteringUniversalOrderStatus = "RECEIVEING"
	WAIT_EVALUATE_StarDemandOrderListV2FilteringUniversalOrderStatus StarDemandOrderListV2FilteringUniversalOrderStatus = "WAIT_EVALUATE"
	WAIT_PAYMENT_StarDemandOrderListV2FilteringUniversalOrderStatus  StarDemandOrderListV2FilteringUniversalOrderStatus = "WAIT_PAYMENT"
)

// All allowed values of StarDemandOrderListV2FilteringUniversalOrderStatus enum
var AllowedStarDemandOrderListV2FilteringUniversalOrderStatusEnumValues = []StarDemandOrderListV2FilteringUniversalOrderStatus{
	"ALL",
	"CANCELED",
	"FINISHED",
	"ONGOING",
	"RECEIVEING",
	"WAIT_EVALUATE",
	"WAIT_PAYMENT",
}

// NewStarDemandOrderListV2FilteringUniversalOrderStatusFromValue returns a pointer to a valid StarDemandOrderListV2FilteringUniversalOrderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandOrderListV2FilteringUniversalOrderStatusFromValue(v string) (*StarDemandOrderListV2FilteringUniversalOrderStatus, error) {
	ev := StarDemandOrderListV2FilteringUniversalOrderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandOrderListV2FilteringUniversalOrderStatus: valid values are %v", v, AllowedStarDemandOrderListV2FilteringUniversalOrderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandOrderListV2FilteringUniversalOrderStatus) IsValid() bool {
	for _, existing := range AllowedStarDemandOrderListV2FilteringUniversalOrderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_order_list_v2_filtering_universal_order_status value
func (v StarDemandOrderListV2FilteringUniversalOrderStatus) Ptr() *StarDemandOrderListV2FilteringUniversalOrderStatus {
	return &v
}
