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

// ClueCouponUpdateV2Status
type ClueCouponUpdateV2Status string

// List of clue_coupon_update_v2_status
const (
	DELETED_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "DELETED"
	UNAUDITED_ClueCouponUpdateV2Status   ClueCouponUpdateV2Status = "UNAUDITED"
	PAUSE_ClueCouponUpdateV2Status       ClueCouponUpdateV2Status = "PAUSE"
	AUDIT_DOING_ClueCouponUpdateV2Status ClueCouponUpdateV2Status = "AUDIT_DOING"
	NORMAL_ClueCouponUpdateV2Status      ClueCouponUpdateV2Status = "NORMAL"
	AUDIT_FAIL_ClueCouponUpdateV2Status  ClueCouponUpdateV2Status = "AUDIT_FAIL"
	OFFLINE_ClueCouponUpdateV2Status     ClueCouponUpdateV2Status = "OFFLINE"
)

// All allowed values of ClueCouponUpdateV2Status enum
var AllowedClueCouponUpdateV2StatusEnumValues = []ClueCouponUpdateV2Status{
	"DELETED",
	"UNAUDITED",
	"PAUSE",
	"AUDIT_DOING",
	"NORMAL",
	"AUDIT_FAIL",
	"OFFLINE",
}

// NewClueCouponUpdateV2StatusFromValue returns a pointer to a valid ClueCouponUpdateV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponUpdateV2StatusFromValue(v string) (*ClueCouponUpdateV2Status, error) {
	ev := ClueCouponUpdateV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponUpdateV2Status: valid values are %v", v, AllowedClueCouponUpdateV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponUpdateV2Status) IsValid() bool {
	for _, existing := range AllowedClueCouponUpdateV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_update_v2_status value
func (v ClueCouponUpdateV2Status) Ptr() *ClueCouponUpdateV2Status {
	return &v
}
