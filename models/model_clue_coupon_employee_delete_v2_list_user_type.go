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

// ClueCouponEmployeeDeleteV2ListUserType
type ClueCouponEmployeeDeleteV2ListUserType string

// List of clue_coupon_employee_delete_v2_list_user_type
const (
	TOUTIAO_ClueCouponEmployeeDeleteV2ListUserType ClueCouponEmployeeDeleteV2ListUserType = "TOUTIAO"
	DOUYIN_ClueCouponEmployeeDeleteV2ListUserType  ClueCouponEmployeeDeleteV2ListUserType = "DOUYIN"
)

// All allowed values of ClueCouponEmployeeDeleteV2ListUserType enum
var AllowedClueCouponEmployeeDeleteV2ListUserTypeEnumValues = []ClueCouponEmployeeDeleteV2ListUserType{
	"TOUTIAO",
	"DOUYIN",
}

// NewClueCouponEmployeeDeleteV2ListUserTypeFromValue returns a pointer to a valid ClueCouponEmployeeDeleteV2ListUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponEmployeeDeleteV2ListUserTypeFromValue(v string) (*ClueCouponEmployeeDeleteV2ListUserType, error) {
	ev := ClueCouponEmployeeDeleteV2ListUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponEmployeeDeleteV2ListUserType: valid values are %v", v, AllowedClueCouponEmployeeDeleteV2ListUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponEmployeeDeleteV2ListUserType) IsValid() bool {
	for _, existing := range AllowedClueCouponEmployeeDeleteV2ListUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_employee_delete_v2_list_user_type value
func (v ClueCouponEmployeeDeleteV2ListUserType) Ptr() *ClueCouponEmployeeDeleteV2ListUserType {
	return &v
}
