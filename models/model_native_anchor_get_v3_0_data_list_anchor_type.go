/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetV30DataListAnchorType
type NativeAnchorGetV30DataListAnchorType string

// List of native_anchor_get_v3.0_data_list_anchor_type
const (
	APP_GAME_NativeAnchorGetV30DataListAnchorType             NativeAnchorGetV30DataListAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorGetV30DataListAnchorType NativeAnchorGetV30DataListAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorGetV30DataListAnchorType             NativeAnchorGetV30DataListAnchorType = "APP_SHOP"
	INSURANCE_NativeAnchorGetV30DataListAnchorType            NativeAnchorGetV30DataListAnchorType = "INSURANCE"
	MICRO_APP_NativeAnchorGetV30DataListAnchorType            NativeAnchorGetV30DataListAnchorType = "MICRO_APP"
	ONLINE_SUBSCRIBE_NativeAnchorGetV30DataListAnchorType     NativeAnchorGetV30DataListAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorGetV30DataListAnchorType         NativeAnchorGetV30DataListAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorGetV30DataListAnchorType        NativeAnchorGetV30DataListAnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorGetV30DataListAnchorType enum
var AllowedNativeAnchorGetV30DataListAnchorTypeEnumValues = []NativeAnchorGetV30DataListAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"MICRO_APP",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorGetV30DataListAnchorTypeFromValue returns a pointer to a valid NativeAnchorGetV30DataListAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30DataListAnchorTypeFromValue(v string) (*NativeAnchorGetV30DataListAnchorType, error) {
	ev := NativeAnchorGetV30DataListAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30DataListAnchorType: valid values are %v", v, AllowedNativeAnchorGetV30DataListAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30DataListAnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30DataListAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_data_list_anchor_type value
func (v NativeAnchorGetV30DataListAnchorType) Ptr() *NativeAnchorGetV30DataListAnchorType {
	return &v
}