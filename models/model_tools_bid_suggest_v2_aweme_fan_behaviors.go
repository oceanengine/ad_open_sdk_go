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

// ToolsBidSuggestV2AwemeFanBehaviors
type ToolsBidSuggestV2AwemeFanBehaviors string

// List of tools_bid_suggest_v2_aweme_fan_behaviors
const (
	SHARED_USER_ToolsBidSuggestV2AwemeFanBehaviors          ToolsBidSuggestV2AwemeFanBehaviors = "SHARED_USER"
	FOLLOWED_USER_ToolsBidSuggestV2AwemeFanBehaviors        ToolsBidSuggestV2AwemeFanBehaviors = "FOLLOWED_USER"
	LIVE_EXCEPTIONAL_ToolsBidSuggestV2AwemeFanBehaviors     ToolsBidSuggestV2AwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_ORDER_ToolsBidSuggestV2AwemeFanBehaviors     ToolsBidSuggestV2AwemeFanBehaviors = "LIVE_GOODS_ORDER"
	GOODS_CARTS_ORDER_ToolsBidSuggestV2AwemeFanBehaviors    ToolsBidSuggestV2AwemeFanBehaviors = "GOODS_CARTS_ORDER"
	GOODS_CARTS_CLICK_ToolsBidSuggestV2AwemeFanBehaviors    ToolsBidSuggestV2AwemeFanBehaviors = "GOODS_CARTS_CLICK"
	LIVE_EFFECTIVE_WATCH_ToolsBidSuggestV2AwemeFanBehaviors ToolsBidSuggestV2AwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIKED_USER_ToolsBidSuggestV2AwemeFanBehaviors           ToolsBidSuggestV2AwemeFanBehaviors = "LIKED_USER"
	LIVE_WATCH_ToolsBidSuggestV2AwemeFanBehaviors           ToolsBidSuggestV2AwemeFanBehaviors = "LIVE_WATCH"
	LIVE_COMMENT_ToolsBidSuggestV2AwemeFanBehaviors         ToolsBidSuggestV2AwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_GOODS_CLICK_ToolsBidSuggestV2AwemeFanBehaviors     ToolsBidSuggestV2AwemeFanBehaviors = "LIVE_GOODS_CLICK"
	COMMENTED_USER_ToolsBidSuggestV2AwemeFanBehaviors       ToolsBidSuggestV2AwemeFanBehaviors = "COMMENTED_USER"
)

// All allowed values of ToolsBidSuggestV2AwemeFanBehaviors enum
var AllowedToolsBidSuggestV2AwemeFanBehaviorsEnumValues = []ToolsBidSuggestV2AwemeFanBehaviors{
	"SHARED_USER",
	"FOLLOWED_USER",
	"LIVE_EXCEPTIONAL",
	"LIVE_GOODS_ORDER",
	"GOODS_CARTS_ORDER",
	"GOODS_CARTS_CLICK",
	"LIVE_EFFECTIVE_WATCH",
	"LIKED_USER",
	"LIVE_WATCH",
	"LIVE_COMMENT",
	"LIVE_GOODS_CLICK",
	"COMMENTED_USER",
}

// NewToolsBidSuggestV2AwemeFanBehaviorsFromValue returns a pointer to a valid ToolsBidSuggestV2AwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2AwemeFanBehaviorsFromValue(v string) (*ToolsBidSuggestV2AwemeFanBehaviors, error) {
	ev := ToolsBidSuggestV2AwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2AwemeFanBehaviors: valid values are %v", v, AllowedToolsBidSuggestV2AwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2AwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2AwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_aweme_fan_behaviors value
func (v ToolsBidSuggestV2AwemeFanBehaviors) Ptr() *ToolsBidSuggestV2AwemeFanBehaviors {
	return &v
}
