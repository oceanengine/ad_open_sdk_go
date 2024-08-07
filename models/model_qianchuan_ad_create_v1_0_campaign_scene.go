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

// QianchuanAdCreateV10CampaignScene
type QianchuanAdCreateV10CampaignScene string

// List of qianchuan_ad_create_v1.0_campaign_scene
const (
	DAILY_SALE_QianchuanAdCreateV10CampaignScene                  QianchuanAdCreateV10CampaignScene = "DAILY_SALE"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanAdCreateV10CampaignScene QianchuanAdCreateV10CampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanAdCreateV10CampaignScene           QianchuanAdCreateV10CampaignScene = "NEW_PRODUCT_BOOST"
)

// All allowed values of QianchuanAdCreateV10CampaignScene enum
var AllowedQianchuanAdCreateV10CampaignSceneEnumValues = []QianchuanAdCreateV10CampaignScene{
	"DAILY_SALE",
	"NEW_CUSTOMER_TRANSFORMATION",
	"NEW_PRODUCT_BOOST",
}

// NewQianchuanAdCreateV10CampaignSceneFromValue returns a pointer to a valid QianchuanAdCreateV10CampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10CampaignSceneFromValue(v string) (*QianchuanAdCreateV10CampaignScene, error) {
	ev := QianchuanAdCreateV10CampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10CampaignScene: valid values are %v", v, AllowedQianchuanAdCreateV10CampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10CampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10CampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_campaign_scene value
func (v QianchuanAdCreateV10CampaignScene) Ptr() *QianchuanAdCreateV10CampaignScene {
	return &v
}
