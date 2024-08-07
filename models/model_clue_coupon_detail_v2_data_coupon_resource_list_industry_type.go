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

// ClueCouponDetailV2DataCouponResourceListIndustryType
type ClueCouponDetailV2DataCouponResourceListIndustryType string

// List of clue_coupon_detail_v2_data_coupon_resource_list_industry_type
const (
	FOOD_ClueCouponDetailV2DataCouponResourceListIndustryType          ClueCouponDetailV2DataCouponResourceListIndustryType = "FOOD"
	OTHER_ClueCouponDetailV2DataCouponResourceListIndustryType         ClueCouponDetailV2DataCouponResourceListIndustryType = "OTHER"
	ENTERTAINMENT_ClueCouponDetailV2DataCouponResourceListIndustryType ClueCouponDetailV2DataCouponResourceListIndustryType = "ENTERTAINMENT"
	GAME_ClueCouponDetailV2DataCouponResourceListIndustryType          ClueCouponDetailV2DataCouponResourceListIndustryType = "GAME"
	FINANCIAL_ClueCouponDetailV2DataCouponResourceListIndustryType     ClueCouponDetailV2DataCouponResourceListIndustryType = "FINANCIAL"
	TICKET_ClueCouponDetailV2DataCouponResourceListIndustryType        ClueCouponDetailV2DataCouponResourceListIndustryType = "TICKET"
)

// All allowed values of ClueCouponDetailV2DataCouponResourceListIndustryType enum
var AllowedClueCouponDetailV2DataCouponResourceListIndustryTypeEnumValues = []ClueCouponDetailV2DataCouponResourceListIndustryType{
	"FOOD",
	"OTHER",
	"ENTERTAINMENT",
	"GAME",
	"FINANCIAL",
	"TICKET",
}

// NewClueCouponDetailV2DataCouponResourceListIndustryTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataCouponResourceListIndustryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataCouponResourceListIndustryTypeFromValue(v string) (*ClueCouponDetailV2DataCouponResourceListIndustryType, error) {
	ev := ClueCouponDetailV2DataCouponResourceListIndustryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataCouponResourceListIndustryType: valid values are %v", v, AllowedClueCouponDetailV2DataCouponResourceListIndustryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataCouponResourceListIndustryType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataCouponResourceListIndustryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_coupon_resource_list_industry_type value
func (v ClueCouponDetailV2DataCouponResourceListIndustryType) Ptr() *ClueCouponDetailV2DataCouponResourceListIndustryType {
	return &v
}
