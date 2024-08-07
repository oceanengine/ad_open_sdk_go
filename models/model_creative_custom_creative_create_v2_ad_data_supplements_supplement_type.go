/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType
type CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType string

// List of creative_custom_creative_create_v2_ad_data_supplements_supplement_type
const (
	CLOUD_GAME_CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType = "CLOUD_GAME"
	NORMAL_CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType     CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType = "NORMAL"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType enum
var AllowedCreativeCustomCreativeCreateV2AdDataSupplementsSupplementTypeEnumValues = []CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType{
	"CLOUD_GAME",
	"NORMAL",
}

// NewCreativeCustomCreativeCreateV2AdDataSupplementsSupplementTypeFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataSupplementsSupplementTypeFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType, error) {
	ev := CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataSupplementsSupplementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataSupplementsSupplementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_supplements_supplement_type value
func (v CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType) Ptr() *CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType {
	return &v
}
