/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdvertiserGetV10FilteringCampaignScene
type QianchuanReportAdvertiserGetV10FilteringCampaignScene string

// List of qianchuan_report_advertiser_get_v1.0_filtering_campaign_scene
const (
	DAILY_SALE_QianchuanReportAdvertiserGetV10FilteringCampaignScene                  QianchuanReportAdvertiserGetV10FilteringCampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanReportAdvertiserGetV10FilteringCampaignScene                   QianchuanReportAdvertiserGetV10FilteringCampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanReportAdvertiserGetV10FilteringCampaignScene QianchuanReportAdvertiserGetV10FilteringCampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	PLANT_GRASS_QianchuanReportAdvertiserGetV10FilteringCampaignScene                 QianchuanReportAdvertiserGetV10FilteringCampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanReportAdvertiserGetV10FilteringCampaignScene                QianchuanReportAdvertiserGetV10FilteringCampaignScene = "PRODUCT_HEAT"
)

// All allowed values of QianchuanReportAdvertiserGetV10FilteringCampaignScene enum
var AllowedQianchuanReportAdvertiserGetV10FilteringCampaignSceneEnumValues = []QianchuanReportAdvertiserGetV10FilteringCampaignScene{
	"DAILY_SALE",
	"LIVE_HEAT",
	"NEW_CUSTOMER_TRANSFORMATION",
	"PLANT_GRASS",
	"PRODUCT_HEAT",
}

// NewQianchuanReportAdvertiserGetV10FilteringCampaignSceneFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10FilteringCampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10FilteringCampaignSceneFromValue(v string) (*QianchuanReportAdvertiserGetV10FilteringCampaignScene, error) {
	ev := QianchuanReportAdvertiserGetV10FilteringCampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10FilteringCampaignScene: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10FilteringCampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10FilteringCampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10FilteringCampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_filtering_campaign_scene value
func (v QianchuanReportAdvertiserGetV10FilteringCampaignScene) Ptr() *QianchuanReportAdvertiserGetV10FilteringCampaignScene {
	return &v
}
