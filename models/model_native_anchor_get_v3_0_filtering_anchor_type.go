/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetV30FilteringAnchorType
type NativeAnchorGetV30FilteringAnchorType string

// List of native_anchor_get_v3.0_filtering_anchor_type
const (
	APP_GAME_NativeAnchorGetV30FilteringAnchorType             NativeAnchorGetV30FilteringAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorGetV30FilteringAnchorType NativeAnchorGetV30FilteringAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorGetV30FilteringAnchorType             NativeAnchorGetV30FilteringAnchorType = "APP_SHOP"
	INSURANCE_NativeAnchorGetV30FilteringAnchorType            NativeAnchorGetV30FilteringAnchorType = "INSURANCE"
	MICRO_APP_NativeAnchorGetV30FilteringAnchorType            NativeAnchorGetV30FilteringAnchorType = "MICRO_APP"
	ONLINE_SUBSCRIBE_NativeAnchorGetV30FilteringAnchorType     NativeAnchorGetV30FilteringAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorGetV30FilteringAnchorType         NativeAnchorGetV30FilteringAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorGetV30FilteringAnchorType        NativeAnchorGetV30FilteringAnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorGetV30FilteringAnchorType enum
var AllowedNativeAnchorGetV30FilteringAnchorTypeEnumValues = []NativeAnchorGetV30FilteringAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"MICRO_APP",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorGetV30FilteringAnchorTypeFromValue returns a pointer to a valid NativeAnchorGetV30FilteringAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30FilteringAnchorTypeFromValue(v string) (*NativeAnchorGetV30FilteringAnchorType, error) {
	ev := NativeAnchorGetV30FilteringAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30FilteringAnchorType: valid values are %v", v, AllowedNativeAnchorGetV30FilteringAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30FilteringAnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30FilteringAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_filtering_anchor_type value
func (v NativeAnchorGetV30FilteringAnchorType) Ptr() *NativeAnchorGetV30FilteringAnchorType {
	return &v
}
