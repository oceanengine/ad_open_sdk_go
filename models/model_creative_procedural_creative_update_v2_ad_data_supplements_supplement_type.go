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

// CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType
type CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType string

// List of creative_procedural_creative_update_v2_ad_data_supplements_supplement_type
const (
	CLOUD_GAME_CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType = "CLOUD_GAME"
	NORMAL_CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType     CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType = "NORMAL"
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementTypeEnumValues = []CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType{
	"CLOUD_GAME",
	"NORMAL",
}

// NewCreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementTypeFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementTypeFromValue(v string) (*CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_supplements_supplement_type value
func (v CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType) Ptr() *CreativeProceduralCreativeUpdateV2AdDataSupplementsSupplementType {
	return &v
}