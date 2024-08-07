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

// StardeliveryTaskCreateV30StarTaskAnchorType
type StardeliveryTaskCreateV30StarTaskAnchorType string

// List of stardelivery_task_create_v3.0_star_task_anchor_type
const (
	APP_GAME_StardeliveryTaskCreateV30StarTaskAnchorType             StardeliveryTaskCreateV30StarTaskAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_StardeliveryTaskCreateV30StarTaskAnchorType StardeliveryTaskCreateV30StarTaskAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_StardeliveryTaskCreateV30StarTaskAnchorType             StardeliveryTaskCreateV30StarTaskAnchorType = "APP_SHOP"
)

// All allowed values of StardeliveryTaskCreateV30StarTaskAnchorType enum
var AllowedStardeliveryTaskCreateV30StarTaskAnchorTypeEnumValues = []StardeliveryTaskCreateV30StarTaskAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
}

// NewStardeliveryTaskCreateV30StarTaskAnchorTypeFromValue returns a pointer to a valid StardeliveryTaskCreateV30StarTaskAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskCreateV30StarTaskAnchorTypeFromValue(v string) (*StardeliveryTaskCreateV30StarTaskAnchorType, error) {
	ev := StardeliveryTaskCreateV30StarTaskAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskCreateV30StarTaskAnchorType: valid values are %v", v, AllowedStardeliveryTaskCreateV30StarTaskAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskCreateV30StarTaskAnchorType) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskCreateV30StarTaskAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_create_v3.0_star_task_anchor_type value
func (v StardeliveryTaskCreateV30StarTaskAnchorType) Ptr() *StardeliveryTaskCreateV30StarTaskAnchorType {
	return &v
}
