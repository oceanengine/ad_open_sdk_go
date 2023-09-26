/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeAutoGenerateConfigGetV2DataStrategyDataType
type CreativeAutoGenerateConfigGetV2DataStrategyDataType string

// List of creative_auto_generate_config_get_v2_data_strategy_data_type
const (
	HORIZONTAL2_HORIZONTAL_CreativeAutoGenerateConfigGetV2DataStrategyDataType CreativeAutoGenerateConfigGetV2DataStrategyDataType = "Horizontal2Horizontal"
	HORIZONTAL2_VERTICAL_CreativeAutoGenerateConfigGetV2DataStrategyDataType   CreativeAutoGenerateConfigGetV2DataStrategyDataType = "Horizontal2Vertical"
	VERTICAL2_HORIZONTAL_CreativeAutoGenerateConfigGetV2DataStrategyDataType   CreativeAutoGenerateConfigGetV2DataStrategyDataType = "Vertical2Horizontal"
	VERTICAL2_VERTICAL_CreativeAutoGenerateConfigGetV2DataStrategyDataType     CreativeAutoGenerateConfigGetV2DataStrategyDataType = "Vertical2Vertical"
)

// All allowed values of CreativeAutoGenerateConfigGetV2DataStrategyDataType enum
var AllowedCreativeAutoGenerateConfigGetV2DataStrategyDataTypeEnumValues = []CreativeAutoGenerateConfigGetV2DataStrategyDataType{
	"Horizontal2Horizontal",
	"Horizontal2Vertical",
	"Vertical2Horizontal",
	"Vertical2Vertical",
}

// NewCreativeAutoGenerateConfigGetV2DataStrategyDataTypeFromValue returns a pointer to a valid CreativeAutoGenerateConfigGetV2DataStrategyDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeAutoGenerateConfigGetV2DataStrategyDataTypeFromValue(v string) (*CreativeAutoGenerateConfigGetV2DataStrategyDataType, error) {
	ev := CreativeAutoGenerateConfigGetV2DataStrategyDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeAutoGenerateConfigGetV2DataStrategyDataType: valid values are %v", v, AllowedCreativeAutoGenerateConfigGetV2DataStrategyDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeAutoGenerateConfigGetV2DataStrategyDataType) IsValid() bool {
	for _, existing := range AllowedCreativeAutoGenerateConfigGetV2DataStrategyDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_auto_generate_config_get_v2_data_strategy_data_type value
func (v CreativeAutoGenerateConfigGetV2DataStrategyDataType) Ptr() *CreativeAutoGenerateConfigGetV2DataStrategyDataType {
	return &v
}