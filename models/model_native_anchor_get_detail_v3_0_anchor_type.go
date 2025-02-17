/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30AnchorType
type NativeAnchorGetDetailV30AnchorType string

// List of native_anchor_get_detail_v3.0_anchor_type
const (
	APP_GAME_NativeAnchorGetDetailV30AnchorType             NativeAnchorGetDetailV30AnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_NativeAnchorGetDetailV30AnchorType NativeAnchorGetDetailV30AnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_NativeAnchorGetDetailV30AnchorType             NativeAnchorGetDetailV30AnchorType = "APP_SHOP"
	INSURANCE_NativeAnchorGetDetailV30AnchorType            NativeAnchorGetDetailV30AnchorType = "INSURANCE"
	MICRO_APP_NativeAnchorGetDetailV30AnchorType            NativeAnchorGetDetailV30AnchorType = "MICRO_APP"
	MICRO_GAME_NativeAnchorGetDetailV30AnchorType           NativeAnchorGetDetailV30AnchorType = "MICRO_GAME"
	PRIVATE_CHAT_NativeAnchorGetDetailV30AnchorType         NativeAnchorGetDetailV30AnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_NativeAnchorGetDetailV30AnchorType        NativeAnchorGetDetailV30AnchorType = "SHOPPING_CART"
)

// Ptr returns reference to native_anchor_get_detail_v3.0_anchor_type value
func (v NativeAnchorGetDetailV30AnchorType) Ptr() *NativeAnchorGetDetailV30AnchorType {
	return &v
}
