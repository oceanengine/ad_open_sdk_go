/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2DeviceType
type ToolsBidSuggestV2DeviceType string

// List of tools_bid_suggest_v2_device_type
const (
	MOBILE_ToolsBidSuggestV2DeviceType ToolsBidSuggestV2DeviceType = "MOBILE"
	PAD_ToolsBidSuggestV2DeviceType    ToolsBidSuggestV2DeviceType = "PAD"
)

// Ptr returns reference to tools_bid_suggest_v2_device_type value
func (v ToolsBidSuggestV2DeviceType) Ptr() *ToolsBidSuggestV2DeviceType {
	return &v
}
