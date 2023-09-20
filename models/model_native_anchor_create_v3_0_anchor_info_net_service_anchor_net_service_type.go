/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType
type NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType string

// List of native_anchor_create_v3.0_anchor_info_net_service_anchor_net_service_type
const (
	GENERAL_NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType        NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType = "GENERAL"
	MICRO_APP_NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType      NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType = "MICRO_APP"
	QUICK_APP_NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType      NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType = "QUICK_APP"
	WECHAT_PACKAGE_NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType = "WECHAT_PACKAGE"
	WECOM_PACKAGE_NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType  NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType = "WECOM_PACKAGE"
)

// All allowed values of NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType enum
var AllowedNativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceTypeEnumValues = []NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType{
	"GENERAL",
	"MICRO_APP",
	"QUICK_APP",
	"WECHAT_PACKAGE",
	"WECOM_PACKAGE",
}

// NewNativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceTypeFromValue returns a pointer to a valid NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceTypeFromValue(v string) (*NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType, error) {
	ev := NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType: valid values are %v", v, AllowedNativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_create_v3.0_anchor_info_net_service_anchor_net_service_type value
func (v NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType) Ptr() *NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType {
	return &v
}
