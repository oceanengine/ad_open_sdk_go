/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponGetV2DataListCouponStatus
type ClueCouponGetV2DataListCouponStatus string

// List of clue_coupon_get_v2_data_list_coupon_status
const (
	NORMAL_ClueCouponGetV2DataListCouponStatus      ClueCouponGetV2DataListCouponStatus = "NORMAL"
	PAUSE_ClueCouponGetV2DataListCouponStatus       ClueCouponGetV2DataListCouponStatus = "PAUSE"
	OFFLINE_ClueCouponGetV2DataListCouponStatus     ClueCouponGetV2DataListCouponStatus = "OFFLINE"
	UNAUDITED_ClueCouponGetV2DataListCouponStatus   ClueCouponGetV2DataListCouponStatus = "UNAUDITED"
	AUDIT_DOING_ClueCouponGetV2DataListCouponStatus ClueCouponGetV2DataListCouponStatus = "AUDIT_DOING"
	DELETED_ClueCouponGetV2DataListCouponStatus     ClueCouponGetV2DataListCouponStatus = "DELETED"
	AUDIT_FAIL_ClueCouponGetV2DataListCouponStatus  ClueCouponGetV2DataListCouponStatus = "AUDIT_FAIL"
)

// All allowed values of ClueCouponGetV2DataListCouponStatus enum
var AllowedClueCouponGetV2DataListCouponStatusEnumValues = []ClueCouponGetV2DataListCouponStatus{
	"NORMAL",
	"PAUSE",
	"OFFLINE",
	"UNAUDITED",
	"AUDIT_DOING",
	"DELETED",
	"AUDIT_FAIL",
}

// NewClueCouponGetV2DataListCouponStatusFromValue returns a pointer to a valid ClueCouponGetV2DataListCouponStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2DataListCouponStatusFromValue(v string) (*ClueCouponGetV2DataListCouponStatus, error) {
	ev := ClueCouponGetV2DataListCouponStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2DataListCouponStatus: valid values are %v", v, AllowedClueCouponGetV2DataListCouponStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2DataListCouponStatus) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2DataListCouponStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_status value
func (v ClueCouponGetV2DataListCouponStatus) Ptr() *ClueCouponGetV2DataListCouponStatus {
	return &v
}
