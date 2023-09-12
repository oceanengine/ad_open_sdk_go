/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponCodeGetV2DataListStatus
type ClueCouponCodeGetV2DataListStatus string

// List of clue_coupon_code_get_v2_data_list_status
const (
	USED_ClueCouponCodeGetV2DataListStatus      ClueCouponCodeGetV2DataListStatus = "USED"
	EXPIRED_ClueCouponCodeGetV2DataListStatus   ClueCouponCodeGetV2DataListStatus = "EXPIRED"
	VALID_ClueCouponCodeGetV2DataListStatus     ClueCouponCodeGetV2DataListStatus = "VALID"
	INVALID_ClueCouponCodeGetV2DataListStatus   ClueCouponCodeGetV2DataListStatus = "INVALID"
	ABANDONED_ClueCouponCodeGetV2DataListStatus ClueCouponCodeGetV2DataListStatus = "ABANDONED"
)

// All allowed values of ClueCouponCodeGetV2DataListStatus enum
var AllowedClueCouponCodeGetV2DataListStatusEnumValues = []ClueCouponCodeGetV2DataListStatus{
	"USED",
	"EXPIRED",
	"VALID",
	"INVALID",
	"ABANDONED",
}

// NewClueCouponCodeGetV2DataListStatusFromValue returns a pointer to a valid ClueCouponCodeGetV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCodeGetV2DataListStatusFromValue(v string) (*ClueCouponCodeGetV2DataListStatus, error) {
	ev := ClueCouponCodeGetV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCodeGetV2DataListStatus: valid values are %v", v, AllowedClueCouponCodeGetV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCodeGetV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedClueCouponCodeGetV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_code_get_v2_data_list_status value
func (v ClueCouponCodeGetV2DataListStatus) Ptr() *ClueCouponCodeGetV2DataListStatus {
	return &v
}
