/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataCouponResourceListResourceType
type ClueCouponDetailV2DataCouponResourceListResourceType string

// List of clue_coupon_detail_v2_data_coupon_resource_list_resource_type
const (
	DISCOUNT_ClueCouponDetailV2DataCouponResourceListResourceType ClueCouponDetailV2DataCouponResourceListResourceType = "DISCOUNT"
	COMMON_ClueCouponDetailV2DataCouponResourceListResourceType   ClueCouponDetailV2DataCouponResourceListResourceType = "COMMON"
	PHYSICAL_ClueCouponDetailV2DataCouponResourceListResourceType ClueCouponDetailV2DataCouponResourceListResourceType = "PHYSICAL"
	GAME_ClueCouponDetailV2DataCouponResourceListResourceType     ClueCouponDetailV2DataCouponResourceListResourceType = "GAME"
	FULL_ClueCouponDetailV2DataCouponResourceListResourceType     ClueCouponDetailV2DataCouponResourceListResourceType = "FULL"
)

// All allowed values of ClueCouponDetailV2DataCouponResourceListResourceType enum
var AllowedClueCouponDetailV2DataCouponResourceListResourceTypeEnumValues = []ClueCouponDetailV2DataCouponResourceListResourceType{
	"DISCOUNT",
	"COMMON",
	"PHYSICAL",
	"GAME",
	"FULL",
}

// NewClueCouponDetailV2DataCouponResourceListResourceTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataCouponResourceListResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataCouponResourceListResourceTypeFromValue(v string) (*ClueCouponDetailV2DataCouponResourceListResourceType, error) {
	ev := ClueCouponDetailV2DataCouponResourceListResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataCouponResourceListResourceType: valid values are %v", v, AllowedClueCouponDetailV2DataCouponResourceListResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataCouponResourceListResourceType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataCouponResourceListResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_resource_list_resource_type value
func (v ClueCouponDetailV2DataCouponResourceListResourceType) Ptr() *ClueCouponDetailV2DataCouponResourceListResourceType {
	return &v
}
