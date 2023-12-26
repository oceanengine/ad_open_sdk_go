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

// QianchuanCreativeGetV10FilteringMarketingScene
type QianchuanCreativeGetV10FilteringMarketingScene string

// List of qianchuan_creative_get_v1.0_filtering_marketing_scene
const (
	ALL_QianchuanCreativeGetV10FilteringMarketingScene           QianchuanCreativeGetV10FilteringMarketingScene = "ALL"
	FEED_QianchuanCreativeGetV10FilteringMarketingScene          QianchuanCreativeGetV10FilteringMarketingScene = "FEED"
	SEARCH_QianchuanCreativeGetV10FilteringMarketingScene        QianchuanCreativeGetV10FilteringMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanCreativeGetV10FilteringMarketingScene QianchuanCreativeGetV10FilteringMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanCreativeGetV10FilteringMarketingScene enum
var AllowedQianchuanCreativeGetV10FilteringMarketingSceneEnumValues = []QianchuanCreativeGetV10FilteringMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanCreativeGetV10FilteringMarketingSceneFromValue returns a pointer to a valid QianchuanCreativeGetV10FilteringMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10FilteringMarketingSceneFromValue(v string) (*QianchuanCreativeGetV10FilteringMarketingScene, error) {
	ev := QianchuanCreativeGetV10FilteringMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10FilteringMarketingScene: valid values are %v", v, AllowedQianchuanCreativeGetV10FilteringMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10FilteringMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10FilteringMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_filtering_marketing_scene value
func (v QianchuanCreativeGetV10FilteringMarketingScene) Ptr() *QianchuanCreativeGetV10FilteringMarketingScene {
	return &v
}
