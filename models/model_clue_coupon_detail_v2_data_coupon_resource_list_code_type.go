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

// ClueCouponDetailV2DataCouponResourceListCodeType
type ClueCouponDetailV2DataCouponResourceListCodeType string

// List of clue_coupon_detail_v2_data_coupon_resource_list_code_type
const (
	API_ClueCouponDetailV2DataCouponResourceListCodeType      ClueCouponDetailV2DataCouponResourceListCodeType = "API"
	MERCHANT_ClueCouponDetailV2DataCouponResourceListCodeType ClueCouponDetailV2DataCouponResourceListCodeType = "MERCHANT"
	PLATFORM_ClueCouponDetailV2DataCouponResourceListCodeType ClueCouponDetailV2DataCouponResourceListCodeType = "PLATFORM"
	COMMON_ClueCouponDetailV2DataCouponResourceListCodeType   ClueCouponDetailV2DataCouponResourceListCodeType = "COMMON"
)

// All allowed values of ClueCouponDetailV2DataCouponResourceListCodeType enum
var AllowedClueCouponDetailV2DataCouponResourceListCodeTypeEnumValues = []ClueCouponDetailV2DataCouponResourceListCodeType{
	"API",
	"MERCHANT",
	"PLATFORM",
	"COMMON",
}

// NewClueCouponDetailV2DataCouponResourceListCodeTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataCouponResourceListCodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataCouponResourceListCodeTypeFromValue(v string) (*ClueCouponDetailV2DataCouponResourceListCodeType, error) {
	ev := ClueCouponDetailV2DataCouponResourceListCodeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataCouponResourceListCodeType: valid values are %v", v, AllowedClueCouponDetailV2DataCouponResourceListCodeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataCouponResourceListCodeType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataCouponResourceListCodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_resource_list_code_type value
func (v ClueCouponDetailV2DataCouponResourceListCodeType) Ptr() *ClueCouponDetailV2DataCouponResourceListCodeType {
	return &v
}
