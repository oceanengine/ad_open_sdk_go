/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10MarketingScene
type QianchuanAdCreateV10MarketingScene string

// List of qianchuan_ad_create_v1.0_marketing_scene
const (
	FEED_QianchuanAdCreateV10MarketingScene          QianchuanAdCreateV10MarketingScene = "FEED"
	SEARCH_QianchuanAdCreateV10MarketingScene        QianchuanAdCreateV10MarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanAdCreateV10MarketingScene QianchuanAdCreateV10MarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanAdCreateV10MarketingScene enum
var AllowedQianchuanAdCreateV10MarketingSceneEnumValues = []QianchuanAdCreateV10MarketingScene{
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanAdCreateV10MarketingSceneFromValue returns a pointer to a valid QianchuanAdCreateV10MarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MarketingSceneFromValue(v string) (*QianchuanAdCreateV10MarketingScene, error) {
	ev := QianchuanAdCreateV10MarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MarketingScene: valid values are %v", v, AllowedQianchuanAdCreateV10MarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_marketing_scene value
func (v QianchuanAdCreateV10MarketingScene) Ptr() *QianchuanAdCreateV10MarketingScene {
	return &v
}
