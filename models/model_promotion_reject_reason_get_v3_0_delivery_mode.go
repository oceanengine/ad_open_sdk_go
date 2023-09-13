/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionRejectReasonGetV30DeliveryMode
type PromotionRejectReasonGetV30DeliveryMode string

// List of promotion_reject_reason_get_v3.0_delivery_mode
const (
	MANUAL_PromotionRejectReasonGetV30DeliveryMode     PromotionRejectReasonGetV30DeliveryMode = "MANUAL"
	PROCEDURAL_PromotionRejectReasonGetV30DeliveryMode PromotionRejectReasonGetV30DeliveryMode = "PROCEDURAL"
)

// All allowed values of PromotionRejectReasonGetV30DeliveryMode enum
var AllowedPromotionRejectReasonGetV30DeliveryModeEnumValues = []PromotionRejectReasonGetV30DeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewPromotionRejectReasonGetV30DeliveryModeFromValue returns a pointer to a valid PromotionRejectReasonGetV30DeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionRejectReasonGetV30DeliveryModeFromValue(v string) (*PromotionRejectReasonGetV30DeliveryMode, error) {
	ev := PromotionRejectReasonGetV30DeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionRejectReasonGetV30DeliveryMode: valid values are %v", v, AllowedPromotionRejectReasonGetV30DeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionRejectReasonGetV30DeliveryMode) IsValid() bool {
	for _, existing := range AllowedPromotionRejectReasonGetV30DeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_reject_reason_get_v3.0_delivery_mode value
func (v PromotionRejectReasonGetV30DeliveryMode) Ptr() *PromotionRejectReasonGetV30DeliveryMode {
	return &v
}
