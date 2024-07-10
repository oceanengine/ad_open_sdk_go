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

// NativeAnchorDeleteV30AnchorType
type NativeAnchorDeleteV30AnchorType string

// List of native_anchor_delete_v3.0_anchor_type
const (
	APP_GAME_NativeAnchorDeleteV30AnchorType             NativeAnchorDeleteV30AnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorDeleteV30AnchorType NativeAnchorDeleteV30AnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorDeleteV30AnchorType             NativeAnchorDeleteV30AnchorType = "APP_SHOP"
	INSURANCE_NativeAnchorDeleteV30AnchorType            NativeAnchorDeleteV30AnchorType = "INSURANCE"
	ONLINE_SUBSCRIBE_NativeAnchorDeleteV30AnchorType     NativeAnchorDeleteV30AnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorDeleteV30AnchorType         NativeAnchorDeleteV30AnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorDeleteV30AnchorType        NativeAnchorDeleteV30AnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorDeleteV30AnchorType enum
var AllowedNativeAnchorDeleteV30AnchorTypeEnumValues = []NativeAnchorDeleteV30AnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorDeleteV30AnchorTypeFromValue returns a pointer to a valid NativeAnchorDeleteV30AnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorDeleteV30AnchorTypeFromValue(v string) (*NativeAnchorDeleteV30AnchorType, error) {
	ev := NativeAnchorDeleteV30AnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorDeleteV30AnchorType: valid values are %v", v, AllowedNativeAnchorDeleteV30AnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorDeleteV30AnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorDeleteV30AnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_delete_v3.0_anchor_type value
func (v NativeAnchorDeleteV30AnchorType) Ptr() *NativeAnchorDeleteV30AnchorType {
	return &v
}
