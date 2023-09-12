/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionAutoGenerateConfigGetV30DataStrategyDataType
type PromotionAutoGenerateConfigGetV30DataStrategyDataType string

// List of promotion_auto_generate_config_get_v3.0_data_strategy_data_type
const (
	HORIZONTAL2_HORIZONTAL_PromotionAutoGenerateConfigGetV30DataStrategyDataType PromotionAutoGenerateConfigGetV30DataStrategyDataType = "Horizontal2Horizontal"
	HORIZONTAL2_VERTICAL_PromotionAutoGenerateConfigGetV30DataStrategyDataType   PromotionAutoGenerateConfigGetV30DataStrategyDataType = "Horizontal2Vertical"
	VERTICAL2_HORIZONTAL_PromotionAutoGenerateConfigGetV30DataStrategyDataType   PromotionAutoGenerateConfigGetV30DataStrategyDataType = "Vertical2Horizontal"
	VERTICAL2_VERTICAL_PromotionAutoGenerateConfigGetV30DataStrategyDataType     PromotionAutoGenerateConfigGetV30DataStrategyDataType = "Vertical2Vertical"
)

// All allowed values of PromotionAutoGenerateConfigGetV30DataStrategyDataType enum
var AllowedPromotionAutoGenerateConfigGetV30DataStrategyDataTypeEnumValues = []PromotionAutoGenerateConfigGetV30DataStrategyDataType{
	"Horizontal2Horizontal",
	"Horizontal2Vertical",
	"Vertical2Horizontal",
	"Vertical2Vertical",
}

// NewPromotionAutoGenerateConfigGetV30DataStrategyDataTypeFromValue returns a pointer to a valid PromotionAutoGenerateConfigGetV30DataStrategyDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionAutoGenerateConfigGetV30DataStrategyDataTypeFromValue(v string) (*PromotionAutoGenerateConfigGetV30DataStrategyDataType, error) {
	ev := PromotionAutoGenerateConfigGetV30DataStrategyDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionAutoGenerateConfigGetV30DataStrategyDataType: valid values are %v", v, AllowedPromotionAutoGenerateConfigGetV30DataStrategyDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionAutoGenerateConfigGetV30DataStrategyDataType) IsValid() bool {
	for _, existing := range AllowedPromotionAutoGenerateConfigGetV30DataStrategyDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_auto_generate_config_get_v3.0_data_strategy_data_type value
func (v PromotionAutoGenerateConfigGetV30DataStrategyDataType) Ptr() *PromotionAutoGenerateConfigGetV30DataStrategyDataType {
	return &v
}
