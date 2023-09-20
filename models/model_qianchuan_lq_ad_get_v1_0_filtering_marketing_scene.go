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

// QianchuanLqAdGetV10FilteringMarketingScene
type QianchuanLqAdGetV10FilteringMarketingScene string

// List of qianchuan_lq_ad_get_v1.0_filtering_marketing_scene
const (
	ALL_QianchuanLqAdGetV10FilteringMarketingScene           QianchuanLqAdGetV10FilteringMarketingScene = "ALL"
	FEED_QianchuanLqAdGetV10FilteringMarketingScene          QianchuanLqAdGetV10FilteringMarketingScene = "FEED"
	SEARCH_QianchuanLqAdGetV10FilteringMarketingScene        QianchuanLqAdGetV10FilteringMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanLqAdGetV10FilteringMarketingScene QianchuanLqAdGetV10FilteringMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanLqAdGetV10FilteringMarketingScene enum
var AllowedQianchuanLqAdGetV10FilteringMarketingSceneEnumValues = []QianchuanLqAdGetV10FilteringMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanLqAdGetV10FilteringMarketingSceneFromValue returns a pointer to a valid QianchuanLqAdGetV10FilteringMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanLqAdGetV10FilteringMarketingSceneFromValue(v string) (*QianchuanLqAdGetV10FilteringMarketingScene, error) {
	ev := QianchuanLqAdGetV10FilteringMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanLqAdGetV10FilteringMarketingScene: valid values are %v", v, AllowedQianchuanLqAdGetV10FilteringMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanLqAdGetV10FilteringMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanLqAdGetV10FilteringMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_lq_ad_get_v1.0_filtering_marketing_scene value
func (v QianchuanLqAdGetV10FilteringMarketingScene) Ptr() *QianchuanLqAdGetV10FilteringMarketingScene {
	return &v
}
