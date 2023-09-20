/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponEmployeeCreateV2ListUserType
type ClueCouponEmployeeCreateV2ListUserType string

// List of clue_coupon_employee_create_v2_list_user_type
const (
	DOUYIN_ClueCouponEmployeeCreateV2ListUserType  ClueCouponEmployeeCreateV2ListUserType = "DOUYIN"
	TOUTIAO_ClueCouponEmployeeCreateV2ListUserType ClueCouponEmployeeCreateV2ListUserType = "TOUTIAO"
)

// All allowed values of ClueCouponEmployeeCreateV2ListUserType enum
var AllowedClueCouponEmployeeCreateV2ListUserTypeEnumValues = []ClueCouponEmployeeCreateV2ListUserType{
	"DOUYIN",
	"TOUTIAO",
}

// NewClueCouponEmployeeCreateV2ListUserTypeFromValue returns a pointer to a valid ClueCouponEmployeeCreateV2ListUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponEmployeeCreateV2ListUserTypeFromValue(v string) (*ClueCouponEmployeeCreateV2ListUserType, error) {
	ev := ClueCouponEmployeeCreateV2ListUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponEmployeeCreateV2ListUserType: valid values are %v", v, AllowedClueCouponEmployeeCreateV2ListUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponEmployeeCreateV2ListUserType) IsValid() bool {
	for _, existing := range AllowedClueCouponEmployeeCreateV2ListUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_employee_create_v2_list_user_type value
func (v ClueCouponEmployeeCreateV2ListUserType) Ptr() *ClueCouponEmployeeCreateV2ListUserType {
	return &v
}