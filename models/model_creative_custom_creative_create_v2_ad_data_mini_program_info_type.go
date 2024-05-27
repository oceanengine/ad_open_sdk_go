/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType
type CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType string

// List of creative_custom_creative_create_v2_ad_data_mini_program_info_type
const (
	TEMPLATE_APP_CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType = "TEMPLATE_APP"
	BYTE_GAME_CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType    CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType = "BYTE_GAME"
	SHELL_APP_CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType    CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType = "SHELL_APP"
	BYTE_APP_CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType     CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType = "BYTE_APP"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType enum
var AllowedCreativeCustomCreativeCreateV2AdDataMiniProgramInfoTypeEnumValues = []CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType{
	"TEMPLATE_APP",
	"BYTE_GAME",
	"SHELL_APP",
	"BYTE_APP",
}

// NewCreativeCustomCreativeCreateV2AdDataMiniProgramInfoTypeFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataMiniProgramInfoTypeFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType, error) {
	ev := CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataMiniProgramInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataMiniProgramInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_mini_program_info_type value
func (v CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType) Ptr() *CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType {
	return &v
}
