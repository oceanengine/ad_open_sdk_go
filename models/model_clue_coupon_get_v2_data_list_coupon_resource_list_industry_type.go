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

// ClueCouponGetV2DataListCouponResourceListIndustryType
type ClueCouponGetV2DataListCouponResourceListIndustryType string

// List of clue_coupon_get_v2_data_list_coupon_resource_list_industry_type
const (
	FOOD_ClueCouponGetV2DataListCouponResourceListIndustryType          ClueCouponGetV2DataListCouponResourceListIndustryType = "FOOD"
	OTHER_ClueCouponGetV2DataListCouponResourceListIndustryType         ClueCouponGetV2DataListCouponResourceListIndustryType = "OTHER"
	ENTERTAINMENT_ClueCouponGetV2DataListCouponResourceListIndustryType ClueCouponGetV2DataListCouponResourceListIndustryType = "ENTERTAINMENT"
	GAME_ClueCouponGetV2DataListCouponResourceListIndustryType          ClueCouponGetV2DataListCouponResourceListIndustryType = "GAME"
	FINANCIAL_ClueCouponGetV2DataListCouponResourceListIndustryType     ClueCouponGetV2DataListCouponResourceListIndustryType = "FINANCIAL"
	TICKET_ClueCouponGetV2DataListCouponResourceListIndustryType        ClueCouponGetV2DataListCouponResourceListIndustryType = "TICKET"
)

// All allowed values of ClueCouponGetV2DataListCouponResourceListIndustryType enum
var AllowedClueCouponGetV2DataListCouponResourceListIndustryTypeEnumValues = []ClueCouponGetV2DataListCouponResourceListIndustryType{
	"FOOD",
	"OTHER",
	"ENTERTAINMENT",
	"GAME",
	"FINANCIAL",
	"TICKET",
}

// NewClueCouponGetV2DataListCouponResourceListIndustryTypeFromValue returns a pointer to a valid ClueCouponGetV2DataListCouponResourceListIndustryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2DataListCouponResourceListIndustryTypeFromValue(v string) (*ClueCouponGetV2DataListCouponResourceListIndustryType, error) {
	ev := ClueCouponGetV2DataListCouponResourceListIndustryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2DataListCouponResourceListIndustryType: valid values are %v", v, AllowedClueCouponGetV2DataListCouponResourceListIndustryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2DataListCouponResourceListIndustryType) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2DataListCouponResourceListIndustryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_resource_list_industry_type value
func (v ClueCouponGetV2DataListCouponResourceListIndustryType) Ptr() *ClueCouponGetV2DataListCouponResourceListIndustryType {
	return &v
}
