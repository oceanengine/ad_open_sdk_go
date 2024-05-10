/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponCreateV2CouponResourceListCodeType
type ClueCouponCreateV2CouponResourceListCodeType string

// List of clue_coupon_create_v2_coupon_resource_list_code_type
const (
	MERCHANT_ClueCouponCreateV2CouponResourceListCodeType ClueCouponCreateV2CouponResourceListCodeType = "MERCHANT"
	PLATFORM_ClueCouponCreateV2CouponResourceListCodeType ClueCouponCreateV2CouponResourceListCodeType = "PLATFORM"
	API_ClueCouponCreateV2CouponResourceListCodeType      ClueCouponCreateV2CouponResourceListCodeType = "API"
	COMMON_ClueCouponCreateV2CouponResourceListCodeType   ClueCouponCreateV2CouponResourceListCodeType = "COMMON"
)

// All allowed values of ClueCouponCreateV2CouponResourceListCodeType enum
var AllowedClueCouponCreateV2CouponResourceListCodeTypeEnumValues = []ClueCouponCreateV2CouponResourceListCodeType{
	"MERCHANT",
	"PLATFORM",
	"API",
	"COMMON",
}

// NewClueCouponCreateV2CouponResourceListCodeTypeFromValue returns a pointer to a valid ClueCouponCreateV2CouponResourceListCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCreateV2CouponResourceListCodeTypeFromValue(v string) (*ClueCouponCreateV2CouponResourceListCodeType, error) {
	ev := ClueCouponCreateV2CouponResourceListCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCreateV2CouponResourceListCodeType: valid values are %v", v, AllowedClueCouponCreateV2CouponResourceListCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCreateV2CouponResourceListCodeType) IsValid() bool {
	for _, existing := range AllowedClueCouponCreateV2CouponResourceListCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_create_v2_coupon_resource_list_code_type value
func (v ClueCouponCreateV2CouponResourceListCodeType) Ptr() *ClueCouponCreateV2CouponResourceListCodeType {
	return &v
}
