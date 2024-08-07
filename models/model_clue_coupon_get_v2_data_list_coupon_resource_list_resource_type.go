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

// ClueCouponGetV2DataListCouponResourceListResourceType
type ClueCouponGetV2DataListCouponResourceListResourceType string

// List of clue_coupon_get_v2_data_list_coupon_resource_list_resource_type
const (
	PHYSICAL_ClueCouponGetV2DataListCouponResourceListResourceType ClueCouponGetV2DataListCouponResourceListResourceType = "PHYSICAL"
	GAME_ClueCouponGetV2DataListCouponResourceListResourceType     ClueCouponGetV2DataListCouponResourceListResourceType = "GAME"
	DISCOUNT_ClueCouponGetV2DataListCouponResourceListResourceType ClueCouponGetV2DataListCouponResourceListResourceType = "DISCOUNT"
	FULL_ClueCouponGetV2DataListCouponResourceListResourceType     ClueCouponGetV2DataListCouponResourceListResourceType = "FULL"
	COMMON_ClueCouponGetV2DataListCouponResourceListResourceType   ClueCouponGetV2DataListCouponResourceListResourceType = "COMMON"
)

// All allowed values of ClueCouponGetV2DataListCouponResourceListResourceType enum
var AllowedClueCouponGetV2DataListCouponResourceListResourceTypeEnumValues = []ClueCouponGetV2DataListCouponResourceListResourceType{
	"PHYSICAL",
	"GAME",
	"DISCOUNT",
	"FULL",
	"COMMON",
}

// NewClueCouponGetV2DataListCouponResourceListResourceTypeFromValue returns a pointer to a valid ClueCouponGetV2DataListCouponResourceListResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2DataListCouponResourceListResourceTypeFromValue(v string) (*ClueCouponGetV2DataListCouponResourceListResourceType, error) {
	ev := ClueCouponGetV2DataListCouponResourceListResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2DataListCouponResourceListResourceType: valid values are %v", v, AllowedClueCouponGetV2DataListCouponResourceListResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2DataListCouponResourceListResourceType) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2DataListCouponResourceListResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_resource_list_resource_type value
func (v ClueCouponGetV2DataListCouponResourceListResourceType) Ptr() *ClueCouponGetV2DataListCouponResourceListResourceType {
	return &v
}
