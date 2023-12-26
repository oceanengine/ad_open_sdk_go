/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeStrategyListV2StrategyTypes
type CreativeStrategyListV2StrategyTypes string

// List of creative_strategy_list_v2_strategy_types
const (
	HORIZONTAL2_HORIZONTAL_CreativeStrategyListV2StrategyTypes CreativeStrategyListV2StrategyTypes = "Horizontal2Horizontal"
	HORIZONTAL2_VERTICAL_CreativeStrategyListV2StrategyTypes   CreativeStrategyListV2StrategyTypes = "Horizontal2Vertical"
	VERTICAL2_HORIZONTAL_CreativeStrategyListV2StrategyTypes   CreativeStrategyListV2StrategyTypes = "Vertical2Horizontal"
	VERTICAL2_VERTICAL_CreativeStrategyListV2StrategyTypes     CreativeStrategyListV2StrategyTypes = "Vertical2Vertical"
)

// All allowed values of CreativeStrategyListV2StrategyTypes enum
var AllowedCreativeStrategyListV2StrategyTypesEnumValues = []CreativeStrategyListV2StrategyTypes{
	"Horizontal2Horizontal",
	"Horizontal2Vertical",
	"Vertical2Horizontal",
	"Vertical2Vertical",
}

// NewCreativeStrategyListV2StrategyTypesFromValue returns a pointer to a valid CreativeStrategyListV2StrategyTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeStrategyListV2StrategyTypesFromValue(v string) (*CreativeStrategyListV2StrategyTypes, error) {
	ev := CreativeStrategyListV2StrategyTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeStrategyListV2StrategyTypes: valid values are %v", v, AllowedCreativeStrategyListV2StrategyTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeStrategyListV2StrategyTypes) IsValid() bool {
	for _, existing := range AllowedCreativeStrategyListV2StrategyTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_strategy_list_v2_strategy_types value
func (v CreativeStrategyListV2StrategyTypes) Ptr() *CreativeStrategyListV2StrategyTypes {
	return &v
}
