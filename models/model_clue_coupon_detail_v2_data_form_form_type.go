/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataFormFormType
type ClueCouponDetailV2DataFormFormType string

// List of clue_coupon_detail_v2_data_form_form_type
const (
	ADVANCED_CREATIVE_FORM_ClueCouponDetailV2DataFormFormType ClueCouponDetailV2DataFormFormType = "ADVANCED_CREATIVE_FORM"
	NATIVE_FORM_ClueCouponDetailV2DataFormFormType            ClueCouponDetailV2DataFormFormType = "NATIVE_FORM"
	NORMAL_FORM_ClueCouponDetailV2DataFormFormType            ClueCouponDetailV2DataFormFormType = "NORMAL_FORM"
)

// All allowed values of ClueCouponDetailV2DataFormFormType enum
var AllowedClueCouponDetailV2DataFormFormTypeEnumValues = []ClueCouponDetailV2DataFormFormType{
	"ADVANCED_CREATIVE_FORM",
	"NATIVE_FORM",
	"NORMAL_FORM",
}

// NewClueCouponDetailV2DataFormFormTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataFormFormType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormFormTypeFromValue(v string) (*ClueCouponDetailV2DataFormFormType, error) {
	ev := ClueCouponDetailV2DataFormFormType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormFormType: valid values are %v", v, AllowedClueCouponDetailV2DataFormFormTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormFormType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormFormTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_form_type value
func (v ClueCouponDetailV2DataFormFormType) Ptr() *ClueCouponDetailV2DataFormFormType {
	return &v
}
