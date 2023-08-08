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

// PromotionCreateV30NativeSettingAnchorRelatedType
type PromotionCreateV30NativeSettingAnchorRelatedType string

// List of promotion_create_v3.0_native_setting_anchor_related_type
const (
	AUTO_PromotionCreateV30NativeSettingAnchorRelatedType   PromotionCreateV30NativeSettingAnchorRelatedType = "AUTO"
	OFF_PromotionCreateV30NativeSettingAnchorRelatedType    PromotionCreateV30NativeSettingAnchorRelatedType = "OFF"
	SELECT_PromotionCreateV30NativeSettingAnchorRelatedType PromotionCreateV30NativeSettingAnchorRelatedType = "SELECT"
)

// All allowed values of PromotionCreateV30NativeSettingAnchorRelatedType enum
var AllowedPromotionCreateV30NativeSettingAnchorRelatedTypeEnumValues = []PromotionCreateV30NativeSettingAnchorRelatedType{
	"AUTO",
	"OFF",
	"SELECT",
}

// NewPromotionCreateV30NativeSettingAnchorRelatedTypeFromValue returns a pointer to a valid PromotionCreateV30NativeSettingAnchorRelatedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30NativeSettingAnchorRelatedTypeFromValue(v string) (*PromotionCreateV30NativeSettingAnchorRelatedType, error) {
	ev := PromotionCreateV30NativeSettingAnchorRelatedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30NativeSettingAnchorRelatedType: valid values are %v", v, AllowedPromotionCreateV30NativeSettingAnchorRelatedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30NativeSettingAnchorRelatedType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30NativeSettingAnchorRelatedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_native_setting_anchor_related_type value
func (v PromotionCreateV30NativeSettingAnchorRelatedType) Ptr() *PromotionCreateV30NativeSettingAnchorRelatedType {
	return &v
}
