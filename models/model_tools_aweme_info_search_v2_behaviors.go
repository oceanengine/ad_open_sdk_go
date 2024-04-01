/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeInfoSearchV2Behaviors
type ToolsAwemeInfoSearchV2Behaviors string

// List of tools_aweme_info_search_v2_behaviors
const (
	COMMENTED_USER_ToolsAwemeInfoSearchV2Behaviors       ToolsAwemeInfoSearchV2Behaviors = "COMMENTED_USER"
	FOLLOWED_USER_ToolsAwemeInfoSearchV2Behaviors        ToolsAwemeInfoSearchV2Behaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_ToolsAwemeInfoSearchV2Behaviors    ToolsAwemeInfoSearchV2Behaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_ToolsAwemeInfoSearchV2Behaviors    ToolsAwemeInfoSearchV2Behaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_ToolsAwemeInfoSearchV2Behaviors           ToolsAwemeInfoSearchV2Behaviors = "LIKED_USER"
	LIVE_COMMENT_ToolsAwemeInfoSearchV2Behaviors         ToolsAwemeInfoSearchV2Behaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_ToolsAwemeInfoSearchV2Behaviors ToolsAwemeInfoSearchV2Behaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_ToolsAwemeInfoSearchV2Behaviors     ToolsAwemeInfoSearchV2Behaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_ToolsAwemeInfoSearchV2Behaviors     ToolsAwemeInfoSearchV2Behaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_ToolsAwemeInfoSearchV2Behaviors     ToolsAwemeInfoSearchV2Behaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_ToolsAwemeInfoSearchV2Behaviors           ToolsAwemeInfoSearchV2Behaviors = "LIVE_WATCH"
	SHARED_USER_ToolsAwemeInfoSearchV2Behaviors          ToolsAwemeInfoSearchV2Behaviors = "SHARED_USER"
)

// All allowed values of ToolsAwemeInfoSearchV2Behaviors enum
var AllowedToolsAwemeInfoSearchV2BehaviorsEnumValues = []ToolsAwemeInfoSearchV2Behaviors{
	"COMMENTED_USER",
	"FOLLOWED_USER",
	"GOODS_CARTS_CLICK",
	"GOODS_CARTS_ORDER",
	"LIKED_USER",
	"LIVE_COMMENT",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"LIVE_GOODS_CLICK",
	"LIVE_GOODS_ORDER",
	"LIVE_WATCH",
	"SHARED_USER",
}

// NewToolsAwemeInfoSearchV2BehaviorsFromValue returns a pointer to a valid ToolsAwemeInfoSearchV2Behaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeInfoSearchV2BehaviorsFromValue(v string) (*ToolsAwemeInfoSearchV2Behaviors, error) {
	ev := ToolsAwemeInfoSearchV2Behaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeInfoSearchV2Behaviors: valid values are %v", v, AllowedToolsAwemeInfoSearchV2BehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeInfoSearchV2Behaviors) IsValid() bool {
	for _, existing := range AllowedToolsAwemeInfoSearchV2BehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_info_search_v2_behaviors value
func (v ToolsAwemeInfoSearchV2Behaviors) Ptr() *ToolsAwemeInfoSearchV2Behaviors {
	return &v
}
