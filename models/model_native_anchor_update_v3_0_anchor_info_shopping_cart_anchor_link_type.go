/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType
type NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType string

// List of native_anchor_update_v3.0_anchor_info_shopping_cart_anchor_link_type
const (
	ONE_JUMP_NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType = "ONE_JUMP"
	TWO_JUMP_NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType = "TWO_JUMP"
)

// All allowed values of NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType enum
var AllowedNativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkTypeEnumValues = []NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType{
	"ONE_JUMP",
	"TWO_JUMP",
}

// NewNativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkTypeFromValue returns a pointer to a valid NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkTypeFromValue(v string) (*NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType, error) {
	ev := NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType: valid values are %v", v, AllowedNativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_update_v3.0_anchor_info_shopping_cart_anchor_link_type value
func (v NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType) Ptr() *NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType {
	return &v
}
