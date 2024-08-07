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

// QianchuanSuggestRoiGoalV10MarketingScene
type QianchuanSuggestRoiGoalV10MarketingScene string

// List of qianchuan_suggest_roi_goal_v1.0_marketing_scene
const (
	FEED_QianchuanSuggestRoiGoalV10MarketingScene          QianchuanSuggestRoiGoalV10MarketingScene = "FEED"
	SEARCH_QianchuanSuggestRoiGoalV10MarketingScene        QianchuanSuggestRoiGoalV10MarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanSuggestRoiGoalV10MarketingScene QianchuanSuggestRoiGoalV10MarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanSuggestRoiGoalV10MarketingScene enum
var AllowedQianchuanSuggestRoiGoalV10MarketingSceneEnumValues = []QianchuanSuggestRoiGoalV10MarketingScene{
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanSuggestRoiGoalV10MarketingSceneFromValue returns a pointer to a valid QianchuanSuggestRoiGoalV10MarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestRoiGoalV10MarketingSceneFromValue(v string) (*QianchuanSuggestRoiGoalV10MarketingScene, error) {
	ev := QianchuanSuggestRoiGoalV10MarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestRoiGoalV10MarketingScene: valid values are %v", v, AllowedQianchuanSuggestRoiGoalV10MarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestRoiGoalV10MarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestRoiGoalV10MarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_roi_goal_v1.0_marketing_scene value
func (v QianchuanSuggestRoiGoalV10MarketingScene) Ptr() *QianchuanSuggestRoiGoalV10MarketingScene {
	return &v
}
