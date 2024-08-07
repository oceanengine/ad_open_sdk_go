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

// PromotionListV30DataListNativeSettingAnchorRelatedType
type PromotionListV30DataListNativeSettingAnchorRelatedType string

// List of promotion_list_v3.0_data_list_native_setting_anchor_related_type
const (
	AUTO_PromotionListV30DataListNativeSettingAnchorRelatedType   PromotionListV30DataListNativeSettingAnchorRelatedType = "AUTO"
	OFF_PromotionListV30DataListNativeSettingAnchorRelatedType    PromotionListV30DataListNativeSettingAnchorRelatedType = "OFF"
	SELECT_PromotionListV30DataListNativeSettingAnchorRelatedType PromotionListV30DataListNativeSettingAnchorRelatedType = "SELECT"
)

// All allowed values of PromotionListV30DataListNativeSettingAnchorRelatedType enum
var AllowedPromotionListV30DataListNativeSettingAnchorRelatedTypeEnumValues = []PromotionListV30DataListNativeSettingAnchorRelatedType{
	"AUTO",
	"OFF",
	"SELECT",
}

// NewPromotionListV30DataListNativeSettingAnchorRelatedTypeFromValue returns a pointer to a valid PromotionListV30DataListNativeSettingAnchorRelatedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListNativeSettingAnchorRelatedTypeFromValue(v string) (*PromotionListV30DataListNativeSettingAnchorRelatedType, error) {
	ev := PromotionListV30DataListNativeSettingAnchorRelatedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListNativeSettingAnchorRelatedType: valid values are %v", v, AllowedPromotionListV30DataListNativeSettingAnchorRelatedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListNativeSettingAnchorRelatedType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListNativeSettingAnchorRelatedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_native_setting_anchor_related_type value
func (v PromotionListV30DataListNativeSettingAnchorRelatedType) Ptr() *PromotionListV30DataListNativeSettingAnchorRelatedType {
	return &v
}
