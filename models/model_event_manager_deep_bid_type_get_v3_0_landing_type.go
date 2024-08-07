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

// EventManagerDeepBidTypeGetV30LandingType
type EventManagerDeepBidTypeGetV30LandingType string

// List of event_manager_deep_bid_type_get_v3.0_landing_type
const (
	APP_EventManagerDeepBidTypeGetV30LandingType            EventManagerDeepBidTypeGetV30LandingType = "APP"
	ARTICLE_EventManagerDeepBidTypeGetV30LandingType        EventManagerDeepBidTypeGetV30LandingType = "ARTICLE"
	BRAND_EXTERNAL_EventManagerDeepBidTypeGetV30LandingType EventManagerDeepBidTypeGetV30LandingType = "BRAND_EXTERNAL"
	DPA_EventManagerDeepBidTypeGetV30LandingType            EventManagerDeepBidTypeGetV30LandingType = "DPA"
	GOODS_EventManagerDeepBidTypeGetV30LandingType          EventManagerDeepBidTypeGetV30LandingType = "GOODS"
	LINK_EventManagerDeepBidTypeGetV30LandingType           EventManagerDeepBidTypeGetV30LandingType = "LINK"
	LIVE_EventManagerDeepBidTypeGetV30LandingType           EventManagerDeepBidTypeGetV30LandingType = "LIVE"
	MICRO_GAME_EventManagerDeepBidTypeGetV30LandingType     EventManagerDeepBidTypeGetV30LandingType = "MICRO_GAME"
	NATIVE_ACTION_EventManagerDeepBidTypeGetV30LandingType  EventManagerDeepBidTypeGetV30LandingType = "NATIVE_ACTION"
	QUICK_APP_EventManagerDeepBidTypeGetV30LandingType      EventManagerDeepBidTypeGetV30LandingType = "QUICK_APP"
	SHOP_EventManagerDeepBidTypeGetV30LandingType           EventManagerDeepBidTypeGetV30LandingType = "SHOP"
	STORE_EventManagerDeepBidTypeGetV30LandingType          EventManagerDeepBidTypeGetV30LandingType = "STORE"
)

// All allowed values of EventManagerDeepBidTypeGetV30LandingType enum
var AllowedEventManagerDeepBidTypeGetV30LandingTypeEnumValues = []EventManagerDeepBidTypeGetV30LandingType{
	"APP",
	"ARTICLE",
	"BRAND_EXTERNAL",
	"DPA",
	"GOODS",
	"LINK",
	"LIVE",
	"MICRO_GAME",
	"NATIVE_ACTION",
	"QUICK_APP",
	"SHOP",
	"STORE",
}

// NewEventManagerDeepBidTypeGetV30LandingTypeFromValue returns a pointer to a valid EventManagerDeepBidTypeGetV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerDeepBidTypeGetV30LandingTypeFromValue(v string) (*EventManagerDeepBidTypeGetV30LandingType, error) {
	ev := EventManagerDeepBidTypeGetV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerDeepBidTypeGetV30LandingType: valid values are %v", v, AllowedEventManagerDeepBidTypeGetV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerDeepBidTypeGetV30LandingType) IsValid() bool {
	for _, existing := range AllowedEventManagerDeepBidTypeGetV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_deep_bid_type_get_v3.0_landing_type value
func (v EventManagerDeepBidTypeGetV30LandingType) Ptr() *EventManagerDeepBidTypeGetV30LandingType {
	return &v
}
