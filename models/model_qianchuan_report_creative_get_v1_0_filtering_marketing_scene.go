/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportCreativeGetV10FilteringMarketingScene
type QianchuanReportCreativeGetV10FilteringMarketingScene string

// List of qianchuan_report_creative_get_v1.0_filtering_marketing_scene
const (
	ALL_QianchuanReportCreativeGetV10FilteringMarketingScene           QianchuanReportCreativeGetV10FilteringMarketingScene = "ALL"
	FEED_QianchuanReportCreativeGetV10FilteringMarketingScene          QianchuanReportCreativeGetV10FilteringMarketingScene = "FEED"
	SEARCH_QianchuanReportCreativeGetV10FilteringMarketingScene        QianchuanReportCreativeGetV10FilteringMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanReportCreativeGetV10FilteringMarketingScene QianchuanReportCreativeGetV10FilteringMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringMarketingScene enum
var AllowedQianchuanReportCreativeGetV10FilteringMarketingSceneEnumValues = []QianchuanReportCreativeGetV10FilteringMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanReportCreativeGetV10FilteringMarketingSceneFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringMarketingSceneFromValue(v string) (*QianchuanReportCreativeGetV10FilteringMarketingScene, error) {
	ev := QianchuanReportCreativeGetV10FilteringMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringMarketingScene: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_marketing_scene value
func (v QianchuanReportCreativeGetV10FilteringMarketingScene) Ptr() *QianchuanReportCreativeGetV10FilteringMarketingScene {
	return &v
}
