/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAdvancedCreativeType
type AdGetV2DataAdvancedCreativeType string

// List of ad_get_v2_data_advanced_creative_type
const (
	ATTACHED_CREATIVE_FORM_AdGetV2DataAdvancedCreativeType            AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_FORM"
	ATTACHED_CREATIVE_COMMERCE_CARD_AdGetV2DataAdvancedCreativeType   AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_COMMERCE_CARD"
	ATTACHED_CREATIVE_COUPON_AdGetV2DataAdvancedCreativeType          AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_COUPON"
	ATTACHED_CREATIVE_APP_AdGetV2DataAdvancedCreativeType             AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_APP"
	ATTACHED_CREATIVE_CONSULTANT_AdGetV2DataAdvancedCreativeType      AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_CONSULTANT"
	ATTACHED_CREATIVE_INTERACT_AdGetV2DataAdvancedCreativeType        AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_INTERACT"
	ATTACHED_CREATIVE_NONE_AdGetV2DataAdvancedCreativeType            AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_NONE"
	ATTACHED_CREATIVE_GAME_PACKAGE_AdGetV2DataAdvancedCreativeType    AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_GAME_PACKAGE"
	ATTACHED_CREATIVE_COUPON_INTERACT_AdGetV2DataAdvancedCreativeType AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_COUPON_INTERACT"
	ATTACHED_CREATIVE_PHONE_AdGetV2DataAdvancedCreativeType           AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_PHONE"
	ATTACHED_CREATIVE_VOTE_INTERACT_AdGetV2DataAdvancedCreativeType   AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_VOTE_INTERACT"
	ATTACHED_CREATIVE_DOWNLOAD_CARD_AdGetV2DataAdvancedCreativeType   AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_DOWNLOAD_CARD"
	ATTACHED_CREATIVE_CARD_AdGetV2DataAdvancedCreativeType            AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_CARD"
	ATTACHED_CREATIVE_GAME_SUBSCRIBE_AdGetV2DataAdvancedCreativeType  AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_GAME_SUBSCRIBE"
	ATTACHED_CREATIVE_GAME_FORM_AdGetV2DataAdvancedCreativeType       AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_GAME_FORM"
	ATTACHED_CREATIVE_SMART_PHONE_AdGetV2DataAdvancedCreativeType     AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_SMART_PHONE"
	ATTACHED_CREATIVE_LIVE_CARD_AdGetV2DataAdvancedCreativeType       AdGetV2DataAdvancedCreativeType = "ATTACHED_CREATIVE_LIVE_CARD"
)

// All allowed values of AdGetV2DataAdvancedCreativeType enum
var AllowedAdGetV2DataAdvancedCreativeTypeEnumValues = []AdGetV2DataAdvancedCreativeType{
	"ATTACHED_CREATIVE_FORM",
	"ATTACHED_CREATIVE_COMMERCE_CARD",
	"ATTACHED_CREATIVE_COUPON",
	"ATTACHED_CREATIVE_APP",
	"ATTACHED_CREATIVE_CONSULTANT",
	"ATTACHED_CREATIVE_INTERACT",
	"ATTACHED_CREATIVE_NONE",
	"ATTACHED_CREATIVE_GAME_PACKAGE",
	"ATTACHED_CREATIVE_COUPON_INTERACT",
	"ATTACHED_CREATIVE_PHONE",
	"ATTACHED_CREATIVE_VOTE_INTERACT",
	"ATTACHED_CREATIVE_DOWNLOAD_CARD",
	"ATTACHED_CREATIVE_CARD",
	"ATTACHED_CREATIVE_GAME_SUBSCRIBE",
	"ATTACHED_CREATIVE_GAME_FORM",
	"ATTACHED_CREATIVE_SMART_PHONE",
	"ATTACHED_CREATIVE_LIVE_CARD",
}

// NewAdGetV2DataAdvancedCreativeTypeFromValue returns a pointer to a valid AdGetV2DataAdvancedCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAdvancedCreativeTypeFromValue(v string) (*AdGetV2DataAdvancedCreativeType, error) {
	ev := AdGetV2DataAdvancedCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAdvancedCreativeType: valid values are %v", v, AllowedAdGetV2DataAdvancedCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAdvancedCreativeType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAdvancedCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_advanced_creative_type value
func (v AdGetV2DataAdvancedCreativeType) Ptr() *AdGetV2DataAdvancedCreativeType {
	return &v
}
