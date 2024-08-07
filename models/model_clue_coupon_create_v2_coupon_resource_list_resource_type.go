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

// ClueCouponCreateV2CouponResourceListResourceType
type ClueCouponCreateV2CouponResourceListResourceType string

// List of clue_coupon_create_v2_coupon_resource_list_resource_type
const (
	PHYSICAL_ClueCouponCreateV2CouponResourceListResourceType ClueCouponCreateV2CouponResourceListResourceType = "PHYSICAL"
	GAME_ClueCouponCreateV2CouponResourceListResourceType     ClueCouponCreateV2CouponResourceListResourceType = "GAME"
	DISCOUNT_ClueCouponCreateV2CouponResourceListResourceType ClueCouponCreateV2CouponResourceListResourceType = "DISCOUNT"
	FULL_ClueCouponCreateV2CouponResourceListResourceType     ClueCouponCreateV2CouponResourceListResourceType = "FULL"
	COMMON_ClueCouponCreateV2CouponResourceListResourceType   ClueCouponCreateV2CouponResourceListResourceType = "COMMON"
)

// All allowed values of ClueCouponCreateV2CouponResourceListResourceType enum
var AllowedClueCouponCreateV2CouponResourceListResourceTypeEnumValues = []ClueCouponCreateV2CouponResourceListResourceType{
	"PHYSICAL",
	"GAME",
	"DISCOUNT",
	"FULL",
	"COMMON",
}

// NewClueCouponCreateV2CouponResourceListResourceTypeFromValue returns a pointer to a valid ClueCouponCreateV2CouponResourceListResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCreateV2CouponResourceListResourceTypeFromValue(v string) (*ClueCouponCreateV2CouponResourceListResourceType, error) {
	ev := ClueCouponCreateV2CouponResourceListResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCreateV2CouponResourceListResourceType: valid values are %v", v, AllowedClueCouponCreateV2CouponResourceListResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCreateV2CouponResourceListResourceType) IsValid() bool {
	for _, existing := range AllowedClueCouponCreateV2CouponResourceListResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_create_v2_coupon_resource_list_resource_type value
func (v ClueCouponCreateV2CouponResourceListResourceType) Ptr() *ClueCouponCreateV2CouponResourceListResourceType {
	return &v
}
