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

// StardeliveryTaskDetailV30DataStarTaskAnchorType
type StardeliveryTaskDetailV30DataStarTaskAnchorType string

// List of stardelivery_task_detail_v3.0_data_star_task_anchor_type
const (
	APP_GAME_StardeliveryTaskDetailV30DataStarTaskAnchorType             StardeliveryTaskDetailV30DataStarTaskAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_StardeliveryTaskDetailV30DataStarTaskAnchorType StardeliveryTaskDetailV30DataStarTaskAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_StardeliveryTaskDetailV30DataStarTaskAnchorType             StardeliveryTaskDetailV30DataStarTaskAnchorType = "APP_SHOP"
	INSURANCE_StardeliveryTaskDetailV30DataStarTaskAnchorType            StardeliveryTaskDetailV30DataStarTaskAnchorType = "INSURANCE"
	MICRO_APP_StardeliveryTaskDetailV30DataStarTaskAnchorType            StardeliveryTaskDetailV30DataStarTaskAnchorType = "MICRO_APP"
	MICRO_GAME_StardeliveryTaskDetailV30DataStarTaskAnchorType           StardeliveryTaskDetailV30DataStarTaskAnchorType = "MICRO_GAME"
	ONLINE_SUBSCRIBE_StardeliveryTaskDetailV30DataStarTaskAnchorType     StardeliveryTaskDetailV30DataStarTaskAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_StardeliveryTaskDetailV30DataStarTaskAnchorType         StardeliveryTaskDetailV30DataStarTaskAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_StardeliveryTaskDetailV30DataStarTaskAnchorType        StardeliveryTaskDetailV30DataStarTaskAnchorType = "SHOPPING_CART"
)

// All allowed values of StardeliveryTaskDetailV30DataStarTaskAnchorType enum
var AllowedStardeliveryTaskDetailV30DataStarTaskAnchorTypeEnumValues = []StardeliveryTaskDetailV30DataStarTaskAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"MICRO_APP",
	"MICRO_GAME",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewStardeliveryTaskDetailV30DataStarTaskAnchorTypeFromValue returns a pointer to a valid StardeliveryTaskDetailV30DataStarTaskAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskDetailV30DataStarTaskAnchorTypeFromValue(v string) (*StardeliveryTaskDetailV30DataStarTaskAnchorType, error) {
	ev := StardeliveryTaskDetailV30DataStarTaskAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskDetailV30DataStarTaskAnchorType: valid values are %v", v, AllowedStardeliveryTaskDetailV30DataStarTaskAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskDetailV30DataStarTaskAnchorType) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskDetailV30DataStarTaskAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_detail_v3.0_data_star_task_anchor_type value
func (v StardeliveryTaskDetailV30DataStarTaskAnchorType) Ptr() *StardeliveryTaskDetailV30DataStarTaskAnchorType {
	return &v
}
