/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponEmployeeCreateV2DataSuccessListUserType
type ClueCouponEmployeeCreateV2DataSuccessListUserType string

// List of clue_coupon_employee_create_v2_data_success_list_user_type
const (
	DOUYIN_ClueCouponEmployeeCreateV2DataSuccessListUserType  ClueCouponEmployeeCreateV2DataSuccessListUserType = "DOUYIN"
	TOUTIAO_ClueCouponEmployeeCreateV2DataSuccessListUserType ClueCouponEmployeeCreateV2DataSuccessListUserType = "TOUTIAO"
)

// All allowed values of ClueCouponEmployeeCreateV2DataSuccessListUserType enum
var AllowedClueCouponEmployeeCreateV2DataSuccessListUserTypeEnumValues = []ClueCouponEmployeeCreateV2DataSuccessListUserType{
	"DOUYIN",
	"TOUTIAO",
}

// NewClueCouponEmployeeCreateV2DataSuccessListUserTypeFromValue returns a pointer to a valid ClueCouponEmployeeCreateV2DataSuccessListUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponEmployeeCreateV2DataSuccessListUserTypeFromValue(v string) (*ClueCouponEmployeeCreateV2DataSuccessListUserType, error) {
	ev := ClueCouponEmployeeCreateV2DataSuccessListUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponEmployeeCreateV2DataSuccessListUserType: valid values are %v", v, AllowedClueCouponEmployeeCreateV2DataSuccessListUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponEmployeeCreateV2DataSuccessListUserType) IsValid() bool {
	for _, existing := range AllowedClueCouponEmployeeCreateV2DataSuccessListUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_employee_create_v2_data_success_list_user_type value
func (v ClueCouponEmployeeCreateV2DataSuccessListUserType) Ptr() *ClueCouponEmployeeCreateV2DataSuccessListUserType {
	return &v
}
