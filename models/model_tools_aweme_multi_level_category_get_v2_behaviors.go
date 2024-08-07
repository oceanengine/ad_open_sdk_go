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

// ToolsAwemeMultiLevelCategoryGetV2Behaviors
type ToolsAwemeMultiLevelCategoryGetV2Behaviors string

// List of tools_aweme_multi_level_category_get_v2_behaviors
const (
	LIVE_WATCH_ToolsAwemeMultiLevelCategoryGetV2Behaviors           ToolsAwemeMultiLevelCategoryGetV2Behaviors = "LIVE_WATCH"
	FOLLOWED_USER_ToolsAwemeMultiLevelCategoryGetV2Behaviors        ToolsAwemeMultiLevelCategoryGetV2Behaviors = "FOLLOWED_USER"
	LIVE_COMMENT_ToolsAwemeMultiLevelCategoryGetV2Behaviors         ToolsAwemeMultiLevelCategoryGetV2Behaviors = "LIVE_COMMENT"
	LIVE_GOODS_ORDER_ToolsAwemeMultiLevelCategoryGetV2Behaviors     ToolsAwemeMultiLevelCategoryGetV2Behaviors = "LIVE_GOODS_ORDER"
	GOODS_CARTS_CLICK_ToolsAwemeMultiLevelCategoryGetV2Behaviors    ToolsAwemeMultiLevelCategoryGetV2Behaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_ToolsAwemeMultiLevelCategoryGetV2Behaviors    ToolsAwemeMultiLevelCategoryGetV2Behaviors = "GOODS_CARTS_ORDER"
	LIVE_GOODS_CLICK_ToolsAwemeMultiLevelCategoryGetV2Behaviors     ToolsAwemeMultiLevelCategoryGetV2Behaviors = "LIVE_GOODS_CLICK"
	LIVE_EFFECTIVE_WATCH_ToolsAwemeMultiLevelCategoryGetV2Behaviors ToolsAwemeMultiLevelCategoryGetV2Behaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_ToolsAwemeMultiLevelCategoryGetV2Behaviors     ToolsAwemeMultiLevelCategoryGetV2Behaviors = "LIVE_EXCEPTIONAL"
	LIKED_USER_ToolsAwemeMultiLevelCategoryGetV2Behaviors           ToolsAwemeMultiLevelCategoryGetV2Behaviors = "LIKED_USER"
	SHARED_USER_ToolsAwemeMultiLevelCategoryGetV2Behaviors          ToolsAwemeMultiLevelCategoryGetV2Behaviors = "SHARED_USER"
	COMMENTED_USER_ToolsAwemeMultiLevelCategoryGetV2Behaviors       ToolsAwemeMultiLevelCategoryGetV2Behaviors = "COMMENTED_USER"
)

// All allowed values of ToolsAwemeMultiLevelCategoryGetV2Behaviors enum
var AllowedToolsAwemeMultiLevelCategoryGetV2BehaviorsEnumValues = []ToolsAwemeMultiLevelCategoryGetV2Behaviors{
	"LIVE_WATCH",
	"FOLLOWED_USER",
	"LIVE_COMMENT",
	"LIVE_GOODS_ORDER",
	"GOODS_CARTS_CLICK",
	"GOODS_CARTS_ORDER",
	"LIVE_GOODS_CLICK",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"LIKED_USER",
	"SHARED_USER",
	"COMMENTED_USER",
}

// NewToolsAwemeMultiLevelCategoryGetV2BehaviorsFromValue returns a pointer to a valid ToolsAwemeMultiLevelCategoryGetV2Behaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeMultiLevelCategoryGetV2BehaviorsFromValue(v string) (*ToolsAwemeMultiLevelCategoryGetV2Behaviors, error) {
	ev := ToolsAwemeMultiLevelCategoryGetV2Behaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeMultiLevelCategoryGetV2Behaviors: valid values are %v", v, AllowedToolsAwemeMultiLevelCategoryGetV2BehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeMultiLevelCategoryGetV2Behaviors) IsValid() bool {
	for _, existing := range AllowedToolsAwemeMultiLevelCategoryGetV2BehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_multi_level_category_get_v2_behaviors value
func (v ToolsAwemeMultiLevelCategoryGetV2Behaviors) Ptr() *ToolsAwemeMultiLevelCategoryGetV2Behaviors {
	return &v
}
