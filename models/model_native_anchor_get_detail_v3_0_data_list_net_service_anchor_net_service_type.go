/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType
type NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType string

// List of native_anchor_get_detail_v3.0_data_list_net_service_anchor_net_service_type
const (
	GENERAL_NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType             NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType = "GENERAL"
	MICRO_APP_NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType           NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType = "MICRO_APP"
	QUICK_APP_NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType           NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType = "QUICK_APP"
	WECHAT_PACKAGE_NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType      NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType = "WECHAT_PACKAGE"
	WECOM_PACKAGE_NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType       NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType = "WECOM_PACKAGE"
	WECHAT_EXTERNAL_URL_NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType = "WECHAT_EXTERNAL_URL"
)

// All allowed values of NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType enum
var AllowedNativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceTypeEnumValues = []NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType{
	"GENERAL",
	"MICRO_APP",
	"QUICK_APP",
	"WECHAT_PACKAGE",
	"WECOM_PACKAGE",
	"WECHAT_EXTERNAL_URL",
}

// NewNativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceTypeFromValue returns a pointer to a valid NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceTypeFromValue(v string) (*NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType, error) {
	ev := NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType: valid values are %v", v, AllowedNativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_net_service_anchor_net_service_type value
func (v NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType) Ptr() *NativeAnchorGetDetailV30DataListNetServiceAnchorNetServiceType {
	return &v
}
