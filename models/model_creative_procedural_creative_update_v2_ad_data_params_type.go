/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeUpdateV2AdDataParamsType
type CreativeProceduralCreativeUpdateV2AdDataParamsType string

// List of creative_procedural_creative_update_v2_ad_data_params_type
const (
	CUSTOM_CreativeProceduralCreativeUpdateV2AdDataParamsType CreativeProceduralCreativeUpdateV2AdDataParamsType = "CUSTOM"
	DPA_CreativeProceduralCreativeUpdateV2AdDataParamsType    CreativeProceduralCreativeUpdateV2AdDataParamsType = "DPA"
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataParamsType enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataParamsTypeEnumValues = []CreativeProceduralCreativeUpdateV2AdDataParamsType{
	"CUSTOM",
	"DPA",
}

// NewCreativeProceduralCreativeUpdateV2AdDataParamsTypeFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataParamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataParamsTypeFromValue(v string) (*CreativeProceduralCreativeUpdateV2AdDataParamsType, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataParamsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataParamsType: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataParamsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataParamsType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataParamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_params_type value
func (v CreativeProceduralCreativeUpdateV2AdDataParamsType) Ptr() *CreativeProceduralCreativeUpdateV2AdDataParamsType {
	return &v
}