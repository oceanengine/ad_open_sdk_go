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

// AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType
type AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType string

// List of ad_convert_optimized_target_get_v3.0_data_list_converts_external_actions_convert_data_type
const (
	AD_CONVERT_DATA_TYPE_EVERY_ONE_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType         AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_EVERY_ONE"
	AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType     AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW"
	AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED"
	AD_CONVERT_DATA_TYPE_LESS_THAN_TEN_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType     AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_LESS_THAN_TEN"
	AD_CONVERT_DATA_TYPE_ONLY_ONE_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType          AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_ONLY_ONE"
	AD_CONVERT_DATA_TYPE_PER_TIMES_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType         AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType = "AD_CONVERT_DATA_TYPE_PER_TIMES"
)

// All allowed values of AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType enum
var AllowedAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataTypeEnumValues = []AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType{
	"AD_CONVERT_DATA_TYPE_EVERY_ONE",
	"AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW",
	"AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED",
	"AD_CONVERT_DATA_TYPE_LESS_THAN_TEN",
	"AD_CONVERT_DATA_TYPE_ONLY_ONE",
	"AD_CONVERT_DATA_TYPE_PER_TIMES",
}

// NewAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataTypeFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataTypeFromValue(v string) (*AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType, error) {
	ev := AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_data_list_converts_external_actions_convert_data_type value
func (v AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType) Ptr() *AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsConvertDataType {
	return &v
}
