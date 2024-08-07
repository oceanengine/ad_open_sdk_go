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

// ToolsEstimateAudienceV2Platform
type ToolsEstimateAudienceV2Platform string

// List of tools_estimate_audience_v2_platform
const (
	IPAD_ToolsEstimateAudienceV2Platform    ToolsEstimateAudienceV2Platform = "IPAD"
	NONE_ToolsEstimateAudienceV2Platform    ToolsEstimateAudienceV2Platform = "NONE"
	ANDROID_ToolsEstimateAudienceV2Platform ToolsEstimateAudienceV2Platform = "ANDROID"
	PC_ToolsEstimateAudienceV2Platform      ToolsEstimateAudienceV2Platform = "PC"
	IOS_ToolsEstimateAudienceV2Platform     ToolsEstimateAudienceV2Platform = "IOS"
	WAP_ToolsEstimateAudienceV2Platform     ToolsEstimateAudienceV2Platform = "WAP"
)

// All allowed values of ToolsEstimateAudienceV2Platform enum
var AllowedToolsEstimateAudienceV2PlatformEnumValues = []ToolsEstimateAudienceV2Platform{
	"IPAD",
	"NONE",
	"ANDROID",
	"PC",
	"IOS",
	"WAP",
}

// NewToolsEstimateAudienceV2PlatformFromValue returns a pointer to a valid ToolsEstimateAudienceV2Platform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2PlatformFromValue(v string) (*ToolsEstimateAudienceV2Platform, error) {
	ev := ToolsEstimateAudienceV2Platform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2Platform: valid values are %v", v, AllowedToolsEstimateAudienceV2PlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2Platform) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2PlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_platform value
func (v ToolsEstimateAudienceV2Platform) Ptr() *ToolsEstimateAudienceV2Platform {
	return &v
}
