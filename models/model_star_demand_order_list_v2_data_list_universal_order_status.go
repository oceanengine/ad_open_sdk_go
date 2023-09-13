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

// StarDemandOrderListV2DataListUniversalOrderStatus
type StarDemandOrderListV2DataListUniversalOrderStatus string

// List of star_demand_order_list_v2_data_list_universal_order_status
const (
	ALL_StarDemandOrderListV2DataListUniversalOrderStatus           StarDemandOrderListV2DataListUniversalOrderStatus = "ALL"
	CANCELED_StarDemandOrderListV2DataListUniversalOrderStatus      StarDemandOrderListV2DataListUniversalOrderStatus = "CANCELED"
	FINISHED_StarDemandOrderListV2DataListUniversalOrderStatus      StarDemandOrderListV2DataListUniversalOrderStatus = "FINISHED"
	ONGOING_StarDemandOrderListV2DataListUniversalOrderStatus       StarDemandOrderListV2DataListUniversalOrderStatus = "ONGOING"
	RECEIVEING_StarDemandOrderListV2DataListUniversalOrderStatus    StarDemandOrderListV2DataListUniversalOrderStatus = "RECEIVEING"
	WAIT_EVALUATE_StarDemandOrderListV2DataListUniversalOrderStatus StarDemandOrderListV2DataListUniversalOrderStatus = "WAIT_EVALUATE"
	WAIT_PAYMENT_StarDemandOrderListV2DataListUniversalOrderStatus  StarDemandOrderListV2DataListUniversalOrderStatus = "WAIT_PAYMENT"
)

// All allowed values of StarDemandOrderListV2DataListUniversalOrderStatus enum
var AllowedStarDemandOrderListV2DataListUniversalOrderStatusEnumValues = []StarDemandOrderListV2DataListUniversalOrderStatus{
	"ALL",
	"CANCELED",
	"FINISHED",
	"ONGOING",
	"RECEIVEING",
	"WAIT_EVALUATE",
	"WAIT_PAYMENT",
}

// NewStarDemandOrderListV2DataListUniversalOrderStatusFromValue returns a pointer to a valid StarDemandOrderListV2DataListUniversalOrderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandOrderListV2DataListUniversalOrderStatusFromValue(v string) (*StarDemandOrderListV2DataListUniversalOrderStatus, error) {
	ev := StarDemandOrderListV2DataListUniversalOrderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandOrderListV2DataListUniversalOrderStatus: valid values are %v", v, AllowedStarDemandOrderListV2DataListUniversalOrderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandOrderListV2DataListUniversalOrderStatus) IsValid() bool {
	for _, existing := range AllowedStarDemandOrderListV2DataListUniversalOrderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_order_list_v2_data_list_universal_order_status value
func (v StarDemandOrderListV2DataListUniversalOrderStatus) Ptr() *StarDemandOrderListV2DataListUniversalOrderStatus {
	return &v
}
