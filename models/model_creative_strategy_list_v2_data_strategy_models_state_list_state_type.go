/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeStrategyListV2DataStrategyModelsStateListStateType
type CreativeStrategyListV2DataStrategyModelsStateListStateType string

// List of creative_strategy_list_v2_data_strategy_models_state_list_state_type
const (
	IMAGE_CreativeStrategyListV2DataStrategyModelsStateListStateType CreativeStrategyListV2DataStrategyModelsStateListStateType = "Image"
	TEXT_CreativeStrategyListV2DataStrategyModelsStateListStateType  CreativeStrategyListV2DataStrategyModelsStateListStateType = "Text"
)

// All allowed values of CreativeStrategyListV2DataStrategyModelsStateListStateType enum
var AllowedCreativeStrategyListV2DataStrategyModelsStateListStateTypeEnumValues = []CreativeStrategyListV2DataStrategyModelsStateListStateType{
	"Image",
	"Text",
}

// NewCreativeStrategyListV2DataStrategyModelsStateListStateTypeFromValue returns a pointer to a valid CreativeStrategyListV2DataStrategyModelsStateListStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeStrategyListV2DataStrategyModelsStateListStateTypeFromValue(v string) (*CreativeStrategyListV2DataStrategyModelsStateListStateType, error) {
	ev := CreativeStrategyListV2DataStrategyModelsStateListStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeStrategyListV2DataStrategyModelsStateListStateType: valid values are %v", v, AllowedCreativeStrategyListV2DataStrategyModelsStateListStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeStrategyListV2DataStrategyModelsStateListStateType) IsValid() bool {
	for _, existing := range AllowedCreativeStrategyListV2DataStrategyModelsStateListStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_strategy_list_v2_data_strategy_models_state_list_state_type value
func (v CreativeStrategyListV2DataStrategyModelsStateListStateType) Ptr() *CreativeStrategyListV2DataStrategyModelsStateListStateType {
	return &v
}
