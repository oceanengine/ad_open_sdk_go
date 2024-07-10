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

// CreativeCustomCreativeUpdateV2AdDataAnchorType
type CreativeCustomCreativeUpdateV2AdDataAnchorType string

// List of creative_custom_creative_update_v2_ad_data_anchor_type
const (
	PRIVATE_CHAT_CreativeCustomCreativeUpdateV2AdDataAnchorType         CreativeCustomCreativeUpdateV2AdDataAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_CreativeCustomCreativeUpdateV2AdDataAnchorType        CreativeCustomCreativeUpdateV2AdDataAnchorType = "SHOPPING_CART"
	APP_INTERNET_SERVICE_CreativeCustomCreativeUpdateV2AdDataAnchorType CreativeCustomCreativeUpdateV2AdDataAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_CreativeCustomCreativeUpdateV2AdDataAnchorType             CreativeCustomCreativeUpdateV2AdDataAnchorType = "APP_SHOP"
	ONLINE_SUBSCRIBE_CreativeCustomCreativeUpdateV2AdDataAnchorType     CreativeCustomCreativeUpdateV2AdDataAnchorType = "ONLINE_SUBSCRIBE"
	APP_GAME_CreativeCustomCreativeUpdateV2AdDataAnchorType             CreativeCustomCreativeUpdateV2AdDataAnchorType = "APP_GAME"
	INSURANCE_CreativeCustomCreativeUpdateV2AdDataAnchorType            CreativeCustomCreativeUpdateV2AdDataAnchorType = "INSURANCE"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataAnchorType enum
var AllowedCreativeCustomCreativeUpdateV2AdDataAnchorTypeEnumValues = []CreativeCustomCreativeUpdateV2AdDataAnchorType{
	"PRIVATE_CHAT",
	"SHOPPING_CART",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"ONLINE_SUBSCRIBE",
	"APP_GAME",
	"INSURANCE",
}

// NewCreativeCustomCreativeUpdateV2AdDataAnchorTypeFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataAnchorTypeFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataAnchorType, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataAnchorType: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataAnchorType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_anchor_type value
func (v CreativeCustomCreativeUpdateV2AdDataAnchorType) Ptr() *CreativeCustomCreativeUpdateV2AdDataAnchorType {
	return &v
}
