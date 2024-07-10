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

// AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus
type AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus string

// List of advertiser_delivery_pkg_get_v3.0_data_delivery_pkg_status
const (
	STATUS_CONFIRM_AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus         AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus    AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus      AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus    AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus enum
var AllowedAdvertiserDeliveryPkgGetV30DataDeliveryPkgStatusEnumValues = []AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserDeliveryPkgGetV30DataDeliveryPkgStatusFromValue returns a pointer to a valid AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryPkgGetV30DataDeliveryPkgStatusFromValue(v string) (*AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus, error) {
	ev := AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus: valid values are %v", v, AllowedAdvertiserDeliveryPkgGetV30DataDeliveryPkgStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryPkgGetV30DataDeliveryPkgStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_pkg_get_v3.0_data_delivery_pkg_status value
func (v AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus) Ptr() *AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus {
	return &v
}
