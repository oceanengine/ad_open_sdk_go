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

// PromotionListV30FilteringDeliveryMode
type PromotionListV30FilteringDeliveryMode string

// List of promotion_list_v3.0_filtering_delivery_mode
const (
	MANUAL_PromotionListV30FilteringDeliveryMode     PromotionListV30FilteringDeliveryMode = "MANUAL"
	PROCEDURAL_PromotionListV30FilteringDeliveryMode PromotionListV30FilteringDeliveryMode = "PROCEDURAL"
)

// All allowed values of PromotionListV30FilteringDeliveryMode enum
var AllowedPromotionListV30FilteringDeliveryModeEnumValues = []PromotionListV30FilteringDeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewPromotionListV30FilteringDeliveryModeFromValue returns a pointer to a valid PromotionListV30FilteringDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringDeliveryModeFromValue(v string) (*PromotionListV30FilteringDeliveryMode, error) {
	ev := PromotionListV30FilteringDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringDeliveryMode: valid values are %v", v, AllowedPromotionListV30FilteringDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringDeliveryMode) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_delivery_mode value
func (v PromotionListV30FilteringDeliveryMode) Ptr() *PromotionListV30FilteringDeliveryMode {
	return &v
}
