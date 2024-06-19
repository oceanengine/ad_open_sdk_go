/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType
type CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType string

// List of creative_procedural_creative_update_v2_ad_data_mini_program_info_type
const (
	TEMPLATE_APP_CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType = "TEMPLATE_APP"
	BYTE_APP_CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType     CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType = "BYTE_APP"
	SHELL_APP_CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType    CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType = "SHELL_APP"
	BYTE_GAME_CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType    CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType = "BYTE_GAME"
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoTypeEnumValues = []CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType{
	"TEMPLATE_APP",
	"BYTE_APP",
	"SHELL_APP",
	"BYTE_GAME",
}

// NewCreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoTypeFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoTypeFromValue(v string) (*CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_mini_program_info_type value
func (v CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType) Ptr() *CreativeProceduralCreativeUpdateV2AdDataMiniProgramInfoType {
	return &v
}
