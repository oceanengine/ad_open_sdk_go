/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponEmployeeCreateV2DataExistedListUserType
type ClueCouponEmployeeCreateV2DataExistedListUserType string

// List of clue_coupon_employee_create_v2_data_existed_list_user_type
const (
	DOUYIN_ClueCouponEmployeeCreateV2DataExistedListUserType  ClueCouponEmployeeCreateV2DataExistedListUserType = "DOUYIN"
	TOUTIAO_ClueCouponEmployeeCreateV2DataExistedListUserType ClueCouponEmployeeCreateV2DataExistedListUserType = "TOUTIAO"
)

// All allowed values of ClueCouponEmployeeCreateV2DataExistedListUserType enum
var AllowedClueCouponEmployeeCreateV2DataExistedListUserTypeEnumValues = []ClueCouponEmployeeCreateV2DataExistedListUserType{
	"DOUYIN",
	"TOUTIAO",
}

// NewClueCouponEmployeeCreateV2DataExistedListUserTypeFromValue returns a pointer to a valid ClueCouponEmployeeCreateV2DataExistedListUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponEmployeeCreateV2DataExistedListUserTypeFromValue(v string) (*ClueCouponEmployeeCreateV2DataExistedListUserType, error) {
	ev := ClueCouponEmployeeCreateV2DataExistedListUserType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponEmployeeCreateV2DataExistedListUserType: valid values are %v", v, AllowedClueCouponEmployeeCreateV2DataExistedListUserTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponEmployeeCreateV2DataExistedListUserType) IsValid() bool {
	for _, existing := range AllowedClueCouponEmployeeCreateV2DataExistedListUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_employee_create_v2_data_existed_list_user_type value
func (v ClueCouponEmployeeCreateV2DataExistedListUserType) Ptr() *ClueCouponEmployeeCreateV2DataExistedListUserType {
	return &v
}
