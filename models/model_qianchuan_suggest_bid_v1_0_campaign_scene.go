/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanSuggestBidV10CampaignScene
type QianchuanSuggestBidV10CampaignScene string

// List of qianchuan_suggest_bid_v1.0_campaign_scene
const (
	DAILY_SALE_QianchuanSuggestBidV10CampaignScene                  QianchuanSuggestBidV10CampaignScene = "DAILY_SALE"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanSuggestBidV10CampaignScene QianchuanSuggestBidV10CampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
)

// All allowed values of QianchuanSuggestBidV10CampaignScene enum
var AllowedQianchuanSuggestBidV10CampaignSceneEnumValues = []QianchuanSuggestBidV10CampaignScene{
	"DAILY_SALE",
	"NEW_CUSTOMER_TRANSFORMATION",
}

// NewQianchuanSuggestBidV10CampaignSceneFromValue returns a pointer to a valid QianchuanSuggestBidV10CampaignScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestBidV10CampaignSceneFromValue(v string) (*QianchuanSuggestBidV10CampaignScene, error) {
	ev := QianchuanSuggestBidV10CampaignScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestBidV10CampaignScene: valid values are %v", v, AllowedQianchuanSuggestBidV10CampaignSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestBidV10CampaignScene) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestBidV10CampaignSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_bid_v1.0_campaign_scene value
func (v QianchuanSuggestBidV10CampaignScene) Ptr() *QianchuanSuggestBidV10CampaignScene {
	return &v
}
