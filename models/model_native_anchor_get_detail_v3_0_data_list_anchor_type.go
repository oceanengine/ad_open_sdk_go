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

// NativeAnchorGetDetailV30DataListAnchorType
type NativeAnchorGetDetailV30DataListAnchorType string

// List of native_anchor_get_detail_v3.0_data_list_anchor_type
const (
	APP_GAME_NativeAnchorGetDetailV30DataListAnchorType             NativeAnchorGetDetailV30DataListAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorGetDetailV30DataListAnchorType NativeAnchorGetDetailV30DataListAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorGetDetailV30DataListAnchorType             NativeAnchorGetDetailV30DataListAnchorType = "APP_SHOP"
	INSURANCE_NativeAnchorGetDetailV30DataListAnchorType            NativeAnchorGetDetailV30DataListAnchorType = "INSURANCE"
	ONLINE_SUBSCRIBE_NativeAnchorGetDetailV30DataListAnchorType     NativeAnchorGetDetailV30DataListAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_NativeAnchorGetDetailV30DataListAnchorType         NativeAnchorGetDetailV30DataListAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorGetDetailV30DataListAnchorType        NativeAnchorGetDetailV30DataListAnchorType = "SHOPPING_CART"
)

// All allowed values of NativeAnchorGetDetailV30DataListAnchorType enum
var AllowedNativeAnchorGetDetailV30DataListAnchorTypeEnumValues = []NativeAnchorGetDetailV30DataListAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewNativeAnchorGetDetailV30DataListAnchorTypeFromValue returns a pointer to a valid NativeAnchorGetDetailV30DataListAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetDetailV30DataListAnchorTypeFromValue(v string) (*NativeAnchorGetDetailV30DataListAnchorType, error) {
	ev := NativeAnchorGetDetailV30DataListAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetDetailV30DataListAnchorType: valid values are %v", v, AllowedNativeAnchorGetDetailV30DataListAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetDetailV30DataListAnchorType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetDetailV30DataListAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_anchor_type value
func (v NativeAnchorGetDetailV30DataListAnchorType) Ptr() *NativeAnchorGetDetailV30DataListAnchorType {
	return &v
}
