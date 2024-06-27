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

// ClueCouponGetV2IsDel
type ClueCouponGetV2IsDel string

// List of clue_coupon_get_v2_is_del
const (
	Enum_1_ClueCouponGetV2IsDel ClueCouponGetV2IsDel = "1"
	Enum_0_ClueCouponGetV2IsDel ClueCouponGetV2IsDel = "0"
)

// All allowed values of ClueCouponGetV2IsDel enum
var AllowedClueCouponGetV2IsDelEnumValues = []ClueCouponGetV2IsDel{
	"1",
	"0",
}

// NewClueCouponGetV2IsDelFromValue returns a pointer to a valid ClueCouponGetV2IsDel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2IsDelFromValue(v string) (*ClueCouponGetV2IsDel, error) {
	ev := ClueCouponGetV2IsDel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2IsDel: valid values are %v", v, AllowedClueCouponGetV2IsDelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2IsDel) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2IsDelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_is_del value
func (v ClueCouponGetV2IsDel) Ptr() *ClueCouponGetV2IsDel {
	return &v
}
