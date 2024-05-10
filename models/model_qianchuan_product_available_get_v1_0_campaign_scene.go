/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanProductAvailableGetV10CampaignScene
type QianchuanProductAvailableGetV10CampaignScene string

// List of qianchuan_product_available_get_v1.0_campaign_scene
const (
	DAILY_SALE_QianchuanProductAvailableGetV10CampaignScene                  QianchuanProductAvailableGetV10CampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanProductAvailableGetV10CampaignScene                   QianchuanProductAvailableGetV10CampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanProductAvailableGetV10CampaignScene QianchuanProductAvailableGetV10CampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanProductAvailableGetV10CampaignScene           QianchuanProductAvailableGetV10CampaignScene = "NEW_PRODUCT_BOOST"
	PLANT_GRASS_QianchuanProductAvailableGetV10CampaignScene                 QianchuanProductAvailableGetV10CampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanProductAvailableGetV10CampaignScene                QianchuanProductAvailableGetV10CampaignScene = "PRODUCT_HEAT"
)

// All allowed values of QianchuanProductAvailableGetV10CampaignScene enum
var AllowedQianchuanProductAvailableGetV10CampaignSceneEnumValues = []QianchuanProductAvailableGetV10CampaignScene{
	"DAILY_SALE",
	"LIVE_HEAT",
	"NEW_CUSTOMER_TRANSFORMATION",
	"NEW_PRODUCT_BOOST",
	"PLANT_GRASS",
	"PRODUCT_HEAT",
}

// NewQianchuanProductAvailableGetV10CampaignSceneFromValue returns a pointer to a valid QianchuanProductAvailableGetV10CampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAvailableGetV10CampaignSceneFromValue(v string) (*QianchuanProductAvailableGetV10CampaignScene, error) {
	ev := QianchuanProductAvailableGetV10CampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAvailableGetV10CampaignScene: valid values are %v", v, AllowedQianchuanProductAvailableGetV10CampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAvailableGetV10CampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAvailableGetV10CampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_available_get_v1.0_campaign_scene value
func (v QianchuanProductAvailableGetV10CampaignScene) Ptr() *QianchuanProductAvailableGetV10CampaignScene {
	return &v
}
