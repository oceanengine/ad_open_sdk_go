/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorCreateV30AnchorInfoAnchorType
type NativeAnchorCreateV30AnchorInfoAnchorType string

// List of native_anchor_create_v3.0_anchor_info_anchor_type
const (
	APP_GAME_NativeAnchorCreateV30AnchorInfoAnchorType             NativeAnchorCreateV30AnchorInfoAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorCreateV30AnchorInfoAnchorType NativeAnchorCreateV30AnchorInfoAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorCreateV30AnchorInfoAnchorType             NativeAnchorCreateV30AnchorInfoAnchorType = "APP_SHOP"
	MICRO_APP_NativeAnchorCreateV30AnchorInfoAnchorType            NativeAnchorCreateV30AnchorInfoAnchorType = "MICRO_APP"
	ONLINE_SUBSCRIBE_NativeAnchorCreateV30AnchorInfoAnchorType     NativeAnchorCreateV30AnchorInfoAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorCreateV30AnchorInfoAnchorType         NativeAnchorCreateV30AnchorInfoAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorCreateV30AnchorInfoAnchorType        NativeAnchorCreateV30AnchorInfoAnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorCreateV30AnchorInfoAnchorType enum
var AllowedNativeAnchorCreateV30AnchorInfoAnchorTypeEnumValues = []NativeAnchorCreateV30AnchorInfoAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"MICRO_APP",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorCreateV30AnchorInfoAnchorTypeFromValue returns a pointer to a valid NativeAnchorCreateV30AnchorInfoAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorCreateV30AnchorInfoAnchorTypeFromValue(v string) (*NativeAnchorCreateV30AnchorInfoAnchorType, error) {
	ev := NativeAnchorCreateV30AnchorInfoAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorCreateV30AnchorInfoAnchorType: valid values are %v", v, AllowedNativeAnchorCreateV30AnchorInfoAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorCreateV30AnchorInfoAnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorCreateV30AnchorInfoAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_create_v3.0_anchor_info_anchor_type value
func (v NativeAnchorCreateV30AnchorInfoAnchorType) Ptr() *NativeAnchorCreateV30AnchorInfoAnchorType {
	return &v
}
