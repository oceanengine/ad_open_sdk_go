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

// DouplusOrderListV30DataOrderListAdListAdStatus
type DouplusOrderListV30DataOrderListAdListAdStatus string

// List of douplus_order_list_v3.0_data_order_list_ad_list_ad_status
const (
	AUDITING_DouplusOrderListV30DataOrderListAdListAdStatus      DouplusOrderListV30DataOrderListAdListAdStatus = "AUDITING"
	AUDIT_PAUSE_DouplusOrderListV30DataOrderListAdListAdStatus   DouplusOrderListV30DataOrderListAdListAdStatus = "AUDIT_PAUSE"
	DELIVERIED_DouplusOrderListV30DataOrderListAdListAdStatus    DouplusOrderListV30DataOrderListAdListAdStatus = "DELIVERIED"
	DELIVERING_DouplusOrderListV30DataOrderListAdListAdStatus    DouplusOrderListV30DataOrderListAdListAdStatus = "DELIVERING"
	OFFLINE_AUDIT_DouplusOrderListV30DataOrderListAdListAdStatus DouplusOrderListV30DataOrderListAdListAdStatus = "OFFLINE_AUDIT"
	TIME_NO_REACH_DouplusOrderListV30DataOrderListAdListAdStatus DouplusOrderListV30DataOrderListAdListAdStatus = "TIME_NO_REACH"
	UNDELIVERIED_DouplusOrderListV30DataOrderListAdListAdStatus  DouplusOrderListV30DataOrderListAdListAdStatus = "UNDELIVERIED"
	UNPAID_DouplusOrderListV30DataOrderListAdListAdStatus        DouplusOrderListV30DataOrderListAdListAdStatus = "UNPAID"
	WAIT_TO_START_DouplusOrderListV30DataOrderListAdListAdStatus DouplusOrderListV30DataOrderListAdListAdStatus = "WAIT_TO_START"
)

// All allowed values of DouplusOrderListV30DataOrderListAdListAdStatus enum
var AllowedDouplusOrderListV30DataOrderListAdListAdStatusEnumValues = []DouplusOrderListV30DataOrderListAdListAdStatus{
	"AUDITING",
	"AUDIT_PAUSE",
	"DELIVERIED",
	"DELIVERING",
	"OFFLINE_AUDIT",
	"TIME_NO_REACH",
	"UNDELIVERIED",
	"UNPAID",
	"WAIT_TO_START",
}

// NewDouplusOrderListV30DataOrderListAdListAdStatusFromValue returns a pointer to a valid DouplusOrderListV30DataOrderListAdListAdStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30DataOrderListAdListAdStatusFromValue(v string) (*DouplusOrderListV30DataOrderListAdListAdStatus, error) {
	ev := DouplusOrderListV30DataOrderListAdListAdStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30DataOrderListAdListAdStatus: valid values are %v", v, AllowedDouplusOrderListV30DataOrderListAdListAdStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30DataOrderListAdListAdStatus) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30DataOrderListAdListAdStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_data_order_list_ad_list_ad_status value
func (v DouplusOrderListV30DataOrderListAdListAdStatus) Ptr() *DouplusOrderListV30DataOrderListAdListAdStatus {
	return &v
}
