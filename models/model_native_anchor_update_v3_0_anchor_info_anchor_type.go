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

// NativeAnchorUpdateV30AnchorInfoAnchorType
type NativeAnchorUpdateV30AnchorInfoAnchorType string

// List of native_anchor_update_v3.0_anchor_info_anchor_type
const (
	APP_GAME_NativeAnchorUpdateV30AnchorInfoAnchorType             NativeAnchorUpdateV30AnchorInfoAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorUpdateV30AnchorInfoAnchorType NativeAnchorUpdateV30AnchorInfoAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorUpdateV30AnchorInfoAnchorType             NativeAnchorUpdateV30AnchorInfoAnchorType = "APP_SHOP"
	PRIVATE_CHAT_NativeAnchorUpdateV30AnchorInfoAnchorType         NativeAnchorUpdateV30AnchorInfoAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorUpdateV30AnchorInfoAnchorType        NativeAnchorUpdateV30AnchorInfoAnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorUpdateV30AnchorInfoAnchorType enum
var AllowedNativeAnchorUpdateV30AnchorInfoAnchorTypeEnumValues = []NativeAnchorUpdateV30AnchorInfoAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorUpdateV30AnchorInfoAnchorTypeFromValue returns a pointer to a valid NativeAnchorUpdateV30AnchorInfoAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorUpdateV30AnchorInfoAnchorTypeFromValue(v string) (*NativeAnchorUpdateV30AnchorInfoAnchorType, error) {
	ev := NativeAnchorUpdateV30AnchorInfoAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorUpdateV30AnchorInfoAnchorType: valid values are %v", v, AllowedNativeAnchorUpdateV30AnchorInfoAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorUpdateV30AnchorInfoAnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorUpdateV30AnchorInfoAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_update_v3.0_anchor_info_anchor_type value
func (v NativeAnchorUpdateV30AnchorInfoAnchorType) Ptr() *NativeAnchorUpdateV30AnchorInfoAnchorType {
	return &v
}
