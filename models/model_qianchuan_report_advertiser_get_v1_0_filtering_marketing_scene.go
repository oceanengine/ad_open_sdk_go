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

// QianchuanReportAdvertiserGetV10FilteringMarketingScene
type QianchuanReportAdvertiserGetV10FilteringMarketingScene string

// List of qianchuan_report_advertiser_get_v1.0_filtering_marketing_scene
const (
	ALL_QianchuanReportAdvertiserGetV10FilteringMarketingScene           QianchuanReportAdvertiserGetV10FilteringMarketingScene = "ALL"
	FEED_QianchuanReportAdvertiserGetV10FilteringMarketingScene          QianchuanReportAdvertiserGetV10FilteringMarketingScene = "FEED"
	SEARCH_QianchuanReportAdvertiserGetV10FilteringMarketingScene        QianchuanReportAdvertiserGetV10FilteringMarketingScene = "SEARCH"
	SHOPPING_MALL_QianchuanReportAdvertiserGetV10FilteringMarketingScene QianchuanReportAdvertiserGetV10FilteringMarketingScene = "SHOPPING_MALL"
)

// All allowed values of QianchuanReportAdvertiserGetV10FilteringMarketingScene enum
var AllowedQianchuanReportAdvertiserGetV10FilteringMarketingSceneEnumValues = []QianchuanReportAdvertiserGetV10FilteringMarketingScene{
	"ALL",
	"FEED",
	"SEARCH",
	"SHOPPING_MALL",
}

// NewQianchuanReportAdvertiserGetV10FilteringMarketingSceneFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10FilteringMarketingScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10FilteringMarketingSceneFromValue(v string) (*QianchuanReportAdvertiserGetV10FilteringMarketingScene, error) {
	ev := QianchuanReportAdvertiserGetV10FilteringMarketingScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10FilteringMarketingScene: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10FilteringMarketingSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10FilteringMarketingScene) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10FilteringMarketingSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_filtering_marketing_scene value
func (v QianchuanReportAdvertiserGetV10FilteringMarketingScene) Ptr() *QianchuanReportAdvertiserGetV10FilteringMarketingScene {
	return &v
}
