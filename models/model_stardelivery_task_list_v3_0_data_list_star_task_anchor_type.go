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

// StardeliveryTaskListV30DataListStarTaskAnchorType
type StardeliveryTaskListV30DataListStarTaskAnchorType string

// List of stardelivery_task_list_v3.0_data_list_star_task_anchor_type
const (
	APP_GAME_StardeliveryTaskListV30DataListStarTaskAnchorType             StardeliveryTaskListV30DataListStarTaskAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_StardeliveryTaskListV30DataListStarTaskAnchorType StardeliveryTaskListV30DataListStarTaskAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_StardeliveryTaskListV30DataListStarTaskAnchorType             StardeliveryTaskListV30DataListStarTaskAnchorType = "APP_SHOP"
	INSURANCE_StardeliveryTaskListV30DataListStarTaskAnchorType            StardeliveryTaskListV30DataListStarTaskAnchorType = "INSURANCE"
	MICRO_APP_StardeliveryTaskListV30DataListStarTaskAnchorType            StardeliveryTaskListV30DataListStarTaskAnchorType = "MICRO_APP"
	MICRO_GAME_StardeliveryTaskListV30DataListStarTaskAnchorType           StardeliveryTaskListV30DataListStarTaskAnchorType = "MICRO_GAME"
	ONLINE_SUBSCRIBE_StardeliveryTaskListV30DataListStarTaskAnchorType     StardeliveryTaskListV30DataListStarTaskAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_StardeliveryTaskListV30DataListStarTaskAnchorType         StardeliveryTaskListV30DataListStarTaskAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_StardeliveryTaskListV30DataListStarTaskAnchorType        StardeliveryTaskListV30DataListStarTaskAnchorType = "SHOPPING_CART"
)

// All allowed values of StardeliveryTaskListV30DataListStarTaskAnchorType enum
var AllowedStardeliveryTaskListV30DataListStarTaskAnchorTypeEnumValues = []StardeliveryTaskListV30DataListStarTaskAnchorType{
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

// NewStardeliveryTaskListV30DataListStarTaskAnchorTypeFromValue returns a pointer to a valid StardeliveryTaskListV30DataListStarTaskAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskListV30DataListStarTaskAnchorTypeFromValue(v string) (*StardeliveryTaskListV30DataListStarTaskAnchorType, error) {
	ev := StardeliveryTaskListV30DataListStarTaskAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskListV30DataListStarTaskAnchorType: valid values are %v", v, AllowedStardeliveryTaskListV30DataListStarTaskAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskListV30DataListStarTaskAnchorType) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskListV30DataListStarTaskAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_list_v3.0_data_list_star_task_anchor_type value
func (v StardeliveryTaskListV30DataListStarTaskAnchorType) Ptr() *StardeliveryTaskListV30DataListStarTaskAnchorType {
	return &v
}
