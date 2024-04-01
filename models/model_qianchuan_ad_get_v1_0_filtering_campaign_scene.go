/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10FilteringCampaignScene
type QianchuanAdGetV10FilteringCampaignScene string

// List of qianchuan_ad_get_v1.0_filtering_campaign_scene
const (
	DAILY_SALE_QianchuanAdGetV10FilteringCampaignScene                  QianchuanAdGetV10FilteringCampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanAdGetV10FilteringCampaignScene                   QianchuanAdGetV10FilteringCampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanAdGetV10FilteringCampaignScene QianchuanAdGetV10FilteringCampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanAdGetV10FilteringCampaignScene           QianchuanAdGetV10FilteringCampaignScene = "NEW_PRODUCT_BOOST"
	PLANT_GRASS_QianchuanAdGetV10FilteringCampaignScene                 QianchuanAdGetV10FilteringCampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanAdGetV10FilteringCampaignScene                QianchuanAdGetV10FilteringCampaignScene = "PRODUCT_HEAT"
)

// All allowed values of QianchuanAdGetV10FilteringCampaignScene enum
var AllowedQianchuanAdGetV10FilteringCampaignSceneEnumValues = []QianchuanAdGetV10FilteringCampaignScene{
	"DAILY_SALE",
	"LIVE_HEAT",
	"NEW_CUSTOMER_TRANSFORMATION",
	"NEW_PRODUCT_BOOST",
	"PLANT_GRASS",
	"PRODUCT_HEAT",
}

// NewQianchuanAdGetV10FilteringCampaignSceneFromValue returns a pointer to a valid QianchuanAdGetV10FilteringCampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10FilteringCampaignSceneFromValue(v string) (*QianchuanAdGetV10FilteringCampaignScene, error) {
	ev := QianchuanAdGetV10FilteringCampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10FilteringCampaignScene: valid values are %v", v, AllowedQianchuanAdGetV10FilteringCampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10FilteringCampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10FilteringCampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_filtering_campaign_scene value
func (v QianchuanAdGetV10FilteringCampaignScene) Ptr() *QianchuanAdGetV10FilteringCampaignScene {
	return &v
}
