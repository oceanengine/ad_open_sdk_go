/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30FilteringStarDeliveryType
type PromotionListV30FilteringStarDeliveryType string

// List of promotion_list_v3.0_filtering_star_delivery_type
const (
	NOT_STAR_DELIVERY_PromotionListV30FilteringStarDeliveryType PromotionListV30FilteringStarDeliveryType = "NOT_STAR_DELIVERY"
	STAR_DELIVERY_PromotionListV30FilteringStarDeliveryType     PromotionListV30FilteringStarDeliveryType = "STAR_DELIVERY"
)

// All allowed values of PromotionListV30FilteringStarDeliveryType enum
var AllowedPromotionListV30FilteringStarDeliveryTypeEnumValues = []PromotionListV30FilteringStarDeliveryType{
	"NOT_STAR_DELIVERY",
	"STAR_DELIVERY",
}

// NewPromotionListV30FilteringStarDeliveryTypeFromValue returns a pointer to a valid PromotionListV30FilteringStarDeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringStarDeliveryTypeFromValue(v string) (*PromotionListV30FilteringStarDeliveryType, error) {
	ev := PromotionListV30FilteringStarDeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringStarDeliveryType: valid values are %v", v, AllowedPromotionListV30FilteringStarDeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringStarDeliveryType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringStarDeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_star_delivery_type value
func (v PromotionListV30FilteringStarDeliveryType) Ptr() *PromotionListV30FilteringStarDeliveryType {
	return &v
}
