/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2AdDataAnchorRelatedType
type CreativeCustomCreativeCreateV2AdDataAnchorRelatedType string

// List of creative_custom_creative_create_v2_ad_data_anchor_related_type
const (
	SELECT_CreativeCustomCreativeCreateV2AdDataAnchorRelatedType CreativeCustomCreativeCreateV2AdDataAnchorRelatedType = "SELECT"
	AUTO_CreativeCustomCreativeCreateV2AdDataAnchorRelatedType   CreativeCustomCreativeCreateV2AdDataAnchorRelatedType = "AUTO"
	OFF_CreativeCustomCreativeCreateV2AdDataAnchorRelatedType    CreativeCustomCreativeCreateV2AdDataAnchorRelatedType = "OFF"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataAnchorRelatedType enum
var AllowedCreativeCustomCreativeCreateV2AdDataAnchorRelatedTypeEnumValues = []CreativeCustomCreativeCreateV2AdDataAnchorRelatedType{
	"SELECT",
	"AUTO",
	"OFF",
}

// NewCreativeCustomCreativeCreateV2AdDataAnchorRelatedTypeFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataAnchorRelatedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataAnchorRelatedTypeFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataAnchorRelatedType, error) {
	ev := CreativeCustomCreativeCreateV2AdDataAnchorRelatedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataAnchorRelatedType: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataAnchorRelatedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataAnchorRelatedType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataAnchorRelatedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_anchor_related_type value
func (v CreativeCustomCreativeCreateV2AdDataAnchorRelatedType) Ptr() *CreativeCustomCreativeCreateV2AdDataAnchorRelatedType {
	return &v
}
