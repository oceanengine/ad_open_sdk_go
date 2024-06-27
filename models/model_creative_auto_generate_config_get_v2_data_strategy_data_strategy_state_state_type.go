/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType
type CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType string

// List of creative_auto_generate_config_get_v2_data_strategy_data_strategy_state_state_type
const (
	IMAGE_CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType = "Image"
	TEXT_CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType  CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType = "Text"
)

// All allowed values of CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType enum
var AllowedCreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateTypeEnumValues = []CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType{
	"Image",
	"Text",
}

// NewCreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateTypeFromValue returns a pointer to a valid CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateTypeFromValue(v string) (*CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType, error) {
	ev := CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType: valid values are %v", v, AllowedCreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType) IsValid() bool {
	for _, existing := range AllowedCreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_auto_generate_config_get_v2_data_strategy_data_strategy_state_state_type value
func (v CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType) Ptr() *CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType {
	return &v
}
