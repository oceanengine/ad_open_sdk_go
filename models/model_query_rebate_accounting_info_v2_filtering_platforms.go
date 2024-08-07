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

// QueryRebateAccountingInfoV2FilteringPlatforms
type QueryRebateAccountingInfoV2FilteringPlatforms string

// List of query_rebate_accounting_info_v2_filtering_platforms
const (
	AD_QueryRebateAccountingInfoV2FilteringPlatforms        QueryRebateAccountingInfoV2FilteringPlatforms = "AD"
	ECOMMERCE_QueryRebateAccountingInfoV2FilteringPlatforms QueryRebateAccountingInfoV2FilteringPlatforms = "ECOMMERCE"
	FC_QueryRebateAccountingInfoV2FilteringPlatforms        QueryRebateAccountingInfoV2FilteringPlatforms = "FC"
	NATIVE_QueryRebateAccountingInfoV2FilteringPlatforms    QueryRebateAccountingInfoV2FilteringPlatforms = "NATIVE"
	QC_QueryRebateAccountingInfoV2FilteringPlatforms        QueryRebateAccountingInfoV2FilteringPlatforms = "QC"
	STAR_QueryRebateAccountingInfoV2FilteringPlatforms      QueryRebateAccountingInfoV2FilteringPlatforms = "STAR"
)

// All allowed values of QueryRebateAccountingInfoV2FilteringPlatforms enum
var AllowedQueryRebateAccountingInfoV2FilteringPlatformsEnumValues = []QueryRebateAccountingInfoV2FilteringPlatforms{
	"AD",
	"ECOMMERCE",
	"FC",
	"NATIVE",
	"QC",
	"STAR",
}

// NewQueryRebateAccountingInfoV2FilteringPlatformsFromValue returns a pointer to a valid QueryRebateAccountingInfoV2FilteringPlatforms
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateAccountingInfoV2FilteringPlatformsFromValue(v string) (*QueryRebateAccountingInfoV2FilteringPlatforms, error) {
	ev := QueryRebateAccountingInfoV2FilteringPlatforms(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateAccountingInfoV2FilteringPlatforms: valid values are %v", v, AllowedQueryRebateAccountingInfoV2FilteringPlatformsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateAccountingInfoV2FilteringPlatforms) IsValid() bool {
	for _, existing := range AllowedQueryRebateAccountingInfoV2FilteringPlatformsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_accounting_info_v2_filtering_platforms value
func (v QueryRebateAccountingInfoV2FilteringPlatforms) Ptr() *QueryRebateAccountingInfoV2FilteringPlatforms {
	return &v
}
