/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataFormExtendInfoCountConfigCountType
type ClueCouponDetailV2DataFormExtendInfoCountConfigCountType string

// List of clue_coupon_detail_v2_data_form_extend_info_count_config_count_type
const (
	COUNT_TYPE_INCREMENT_ClueCouponDetailV2DataFormExtendInfoCountConfigCountType ClueCouponDetailV2DataFormExtendInfoCountConfigCountType = "COUNT_TYPE_INCREMENT"
	COUNT_TYPE_DECREMENT_ClueCouponDetailV2DataFormExtendInfoCountConfigCountType ClueCouponDetailV2DataFormExtendInfoCountConfigCountType = "COUNT_TYPE_DECREMENT"
)

// All allowed values of ClueCouponDetailV2DataFormExtendInfoCountConfigCountType enum
var AllowedClueCouponDetailV2DataFormExtendInfoCountConfigCountTypeEnumValues = []ClueCouponDetailV2DataFormExtendInfoCountConfigCountType{
	"COUNT_TYPE_INCREMENT",
	"COUNT_TYPE_DECREMENT",
}

// NewClueCouponDetailV2DataFormExtendInfoCountConfigCountTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataFormExtendInfoCountConfigCountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormExtendInfoCountConfigCountTypeFromValue(v string) (*ClueCouponDetailV2DataFormExtendInfoCountConfigCountType, error) {
	ev := ClueCouponDetailV2DataFormExtendInfoCountConfigCountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormExtendInfoCountConfigCountType: valid values are %v", v, AllowedClueCouponDetailV2DataFormExtendInfoCountConfigCountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormExtendInfoCountConfigCountType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormExtendInfoCountConfigCountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_extend_info_count_config_count_type value
func (v ClueCouponDetailV2DataFormExtendInfoCountConfigCountType) Ptr() *ClueCouponDetailV2DataFormExtendInfoCountConfigCountType {
	return &v
}
