/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandListV2DataListUniversalSettlementType
type StarDemandListV2DataListUniversalSettlementType string

// List of star_demand_list_v2_data_list_universal_settlement_type
const (
	DOU_PLUS_StarDemandListV2DataListUniversalSettlementType     StarDemandListV2DataListUniversalSettlementType = "DOU_PLUS"
	GIFT_StarDemandListV2DataListUniversalSettlementType         StarDemandListV2DataListUniversalSettlementType = "GIFT"
	ALL_StarDemandListV2DataListUniversalSettlementType          StarDemandListV2DataListUniversalSettlementType = "ALL"
	CUSTOMIZE_StarDemandListV2DataListUniversalSettlementType    StarDemandListV2DataListUniversalSettlementType = "CUSTOMIZE"
	EFFECT_StarDemandListV2DataListUniversalSettlementType       StarDemandListV2DataListUniversalSettlementType = "EFFECT"
	STAR_SUPPORT_StarDemandListV2DataListUniversalSettlementType StarDemandListV2DataListUniversalSettlementType = "STAR_SUPPORT"
	EXCHANGE_StarDemandListV2DataListUniversalSettlementType     StarDemandListV2DataListUniversalSettlementType = "EXCHANGE"
	FIXED_PRICE_StarDemandListV2DataListUniversalSettlementType  StarDemandListV2DataListUniversalSettlementType = "FIXED_PRICE"
	MONEY_SHARE_StarDemandListV2DataListUniversalSettlementType  StarDemandListV2DataListUniversalSettlementType = "MONEY_SHARE"
	FLOW_SHARE_StarDemandListV2DataListUniversalSettlementType   StarDemandListV2DataListUniversalSettlementType = "FLOW_SHARE"
	CPM_StarDemandListV2DataListUniversalSettlementType          StarDemandListV2DataListUniversalSettlementType = "CPM"
	CPA_StarDemandListV2DataListUniversalSettlementType          StarDemandListV2DataListUniversalSettlementType = "CPA"
	RANK_StarDemandListV2DataListUniversalSettlementType         StarDemandListV2DataListUniversalSettlementType = "RANK"
)

// All allowed values of StarDemandListV2DataListUniversalSettlementType enum
var AllowedStarDemandListV2DataListUniversalSettlementTypeEnumValues = []StarDemandListV2DataListUniversalSettlementType{
	"DOU_PLUS",
	"GIFT",
	"ALL",
	"CUSTOMIZE",
	"EFFECT",
	"STAR_SUPPORT",
	"EXCHANGE",
	"FIXED_PRICE",
	"MONEY_SHARE",
	"FLOW_SHARE",
	"CPM",
	"CPA",
	"RANK",
}

// NewStarDemandListV2DataListUniversalSettlementTypeFromValue returns a pointer to a valid StarDemandListV2DataListUniversalSettlementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandListV2DataListUniversalSettlementTypeFromValue(v string) (*StarDemandListV2DataListUniversalSettlementType, error) {
	ev := StarDemandListV2DataListUniversalSettlementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandListV2DataListUniversalSettlementType: valid values are %v", v, AllowedStarDemandListV2DataListUniversalSettlementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandListV2DataListUniversalSettlementType) IsValid() bool {
	for _, existing := range AllowedStarDemandListV2DataListUniversalSettlementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_list_v2_data_list_universal_settlement_type value
func (v StarDemandListV2DataListUniversalSettlementType) Ptr() *StarDemandListV2DataListUniversalSettlementType {
	return &v
}
