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

// CreativeCustomCreativeCreateV2AdDataAnchorType
type CreativeCustomCreativeCreateV2AdDataAnchorType string

// List of creative_custom_creative_create_v2_ad_data_anchor_type
const (
	APP_GAME_CreativeCustomCreativeCreateV2AdDataAnchorType             CreativeCustomCreativeCreateV2AdDataAnchorType = "APP_GAME"
	INSURANCE_CreativeCustomCreativeCreateV2AdDataAnchorType            CreativeCustomCreativeCreateV2AdDataAnchorType = "INSURANCE"
	ONLINE_SUBSCRIBE_CreativeCustomCreativeCreateV2AdDataAnchorType     CreativeCustomCreativeCreateV2AdDataAnchorType = "ONLINE_SUBSCRIBE"
	SHOPPING_CART_CreativeCustomCreativeCreateV2AdDataAnchorType        CreativeCustomCreativeCreateV2AdDataAnchorType = "SHOPPING_CART"
	PRIVATE_CHAT_CreativeCustomCreativeCreateV2AdDataAnchorType         CreativeCustomCreativeCreateV2AdDataAnchorType = "PRIVATE_CHAT"
	APP_INTERNET_SERVICE_CreativeCustomCreativeCreateV2AdDataAnchorType CreativeCustomCreativeCreateV2AdDataAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_CreativeCustomCreativeCreateV2AdDataAnchorType             CreativeCustomCreativeCreateV2AdDataAnchorType = "APP_SHOP"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataAnchorType enum
var AllowedCreativeCustomCreativeCreateV2AdDataAnchorTypeEnumValues = []CreativeCustomCreativeCreateV2AdDataAnchorType{
	"APP_GAME",
	"INSURANCE",
	"ONLINE_SUBSCRIBE",
	"SHOPPING_CART",
	"PRIVATE_CHAT",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
}

// NewCreativeCustomCreativeCreateV2AdDataAnchorTypeFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataAnchorTypeFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataAnchorType, error) {
	ev := CreativeCustomCreativeCreateV2AdDataAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataAnchorType: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataAnchorType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_anchor_type value
func (v CreativeCustomCreativeCreateV2AdDataAnchorType) Ptr() *CreativeCustomCreativeCreateV2AdDataAnchorType {
	return &v
}
