/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeSimilarAuthorSearchV2Behaviors
type ToolsAwemeSimilarAuthorSearchV2Behaviors string

// List of tools_aweme_similar_author_search_v2_behaviors
const (
	COMMENTED_USER_ToolsAwemeSimilarAuthorSearchV2Behaviors       ToolsAwemeSimilarAuthorSearchV2Behaviors = "COMMENTED_USER"
	FOLLOWED_USER_ToolsAwemeSimilarAuthorSearchV2Behaviors        ToolsAwemeSimilarAuthorSearchV2Behaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_ToolsAwemeSimilarAuthorSearchV2Behaviors    ToolsAwemeSimilarAuthorSearchV2Behaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_ToolsAwemeSimilarAuthorSearchV2Behaviors    ToolsAwemeSimilarAuthorSearchV2Behaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_ToolsAwemeSimilarAuthorSearchV2Behaviors           ToolsAwemeSimilarAuthorSearchV2Behaviors = "LIKED_USER"
	LIVE_COMMENT_ToolsAwemeSimilarAuthorSearchV2Behaviors         ToolsAwemeSimilarAuthorSearchV2Behaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_ToolsAwemeSimilarAuthorSearchV2Behaviors ToolsAwemeSimilarAuthorSearchV2Behaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_ToolsAwemeSimilarAuthorSearchV2Behaviors     ToolsAwemeSimilarAuthorSearchV2Behaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_ToolsAwemeSimilarAuthorSearchV2Behaviors     ToolsAwemeSimilarAuthorSearchV2Behaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_ToolsAwemeSimilarAuthorSearchV2Behaviors     ToolsAwemeSimilarAuthorSearchV2Behaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_ToolsAwemeSimilarAuthorSearchV2Behaviors           ToolsAwemeSimilarAuthorSearchV2Behaviors = "LIVE_WATCH"
	SHARED_USER_ToolsAwemeSimilarAuthorSearchV2Behaviors          ToolsAwemeSimilarAuthorSearchV2Behaviors = "SHARED_USER"
)

// All allowed values of ToolsAwemeSimilarAuthorSearchV2Behaviors enum
var AllowedToolsAwemeSimilarAuthorSearchV2BehaviorsEnumValues = []ToolsAwemeSimilarAuthorSearchV2Behaviors{
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

// NewToolsAwemeSimilarAuthorSearchV2BehaviorsFromValue returns a pointer to a valid ToolsAwemeSimilarAuthorSearchV2Behaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeSimilarAuthorSearchV2BehaviorsFromValue(v string) (*ToolsAwemeSimilarAuthorSearchV2Behaviors, error) {
	ev := ToolsAwemeSimilarAuthorSearchV2Behaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeSimilarAuthorSearchV2Behaviors: valid values are %v", v, AllowedToolsAwemeSimilarAuthorSearchV2BehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeSimilarAuthorSearchV2Behaviors) IsValid() bool {
	for _, existing := range AllowedToolsAwemeSimilarAuthorSearchV2BehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_similar_author_search_v2_behaviors value
func (v ToolsAwemeSimilarAuthorSearchV2Behaviors) Ptr() *ToolsAwemeSimilarAuthorSearchV2Behaviors {
	return &v
}
