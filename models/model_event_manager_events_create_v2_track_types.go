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

// EventManagerEventsCreateV2TrackTypes
type EventManagerEventsCreateV2TrackTypes string

// List of event_manager_events_create_v2_track_types
const (
	JSSDK_EventManagerEventsCreateV2TrackTypes             EventManagerEventsCreateV2TrackTypes = "JSSDK"
	XPATH_EventManagerEventsCreateV2TrackTypes             EventManagerEventsCreateV2TrackTypes = "XPATH"
	APPLICATION_API_EventManagerEventsCreateV2TrackTypes   EventManagerEventsCreateV2TrackTypes = "APPLICATION_API"
	EXTERNAL_API_EventManagerEventsCreateV2TrackTypes      EventManagerEventsCreateV2TrackTypes = "EXTERNAL_API"
	APPLICATION_SDK_EventManagerEventsCreateV2TrackTypes   EventManagerEventsCreateV2TrackTypes = "APPLICATION_SDK"
	MINI_PROGRAME_SDK_EventManagerEventsCreateV2TrackTypes EventManagerEventsCreateV2TrackTypes = "MINI_PROGRAME_SDK"
	MINI_PROGRAME_API_EventManagerEventsCreateV2TrackTypes EventManagerEventsCreateV2TrackTypes = "MINI_PROGRAME_API"
	OFFLINE_CALLBACK_EventManagerEventsCreateV2TrackTypes  EventManagerEventsCreateV2TrackTypes = "OFFLINE_CALLBACK"
	OFFLINE_API_EventManagerEventsCreateV2TrackTypes       EventManagerEventsCreateV2TrackTypes = "OFFLINE_API"
	APPLOG_EventManagerEventsCreateV2TrackTypes            EventManagerEventsCreateV2TrackTypes = "APPLOG"
	TETRIS_EventManagerEventsCreateV2TrackTypes            EventManagerEventsCreateV2TrackTypes = "TETRIS"
	TETRIS_FE_EventManagerEventsCreateV2TrackTypes         EventManagerEventsCreateV2TrackTypes = "TETRIS_FE"
	FLY_FISH_EventManagerEventsCreateV2TrackTypes          EventManagerEventsCreateV2TrackTypes = "FLY_FISH"
	E_CHAT_EventManagerEventsCreateV2TrackTypes            EventManagerEventsCreateV2TrackTypes = "E_CHAT"
	QUICK_APP_API_EventManagerEventsCreateV2TrackTypes     EventManagerEventsCreateV2TrackTypes = "QUICK_APP_API"
)

// All allowed values of EventManagerEventsCreateV2TrackTypes enum
var AllowedEventManagerEventsCreateV2TrackTypesEnumValues = []EventManagerEventsCreateV2TrackTypes{
	"JSSDK",
	"XPATH",
	"APPLICATION_API",
	"EXTERNAL_API",
	"APPLICATION_SDK",
	"MINI_PROGRAME_SDK",
	"MINI_PROGRAME_API",
	"OFFLINE_CALLBACK",
	"OFFLINE_API",
	"APPLOG",
	"TETRIS",
	"TETRIS_FE",
	"FLY_FISH",
	"E_CHAT",
	"QUICK_APP_API",
}

// NewEventManagerEventsCreateV2TrackTypesFromValue returns a pointer to a valid EventManagerEventsCreateV2TrackTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerEventsCreateV2TrackTypesFromValue(v string) (*EventManagerEventsCreateV2TrackTypes, error) {
	ev := EventManagerEventsCreateV2TrackTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerEventsCreateV2TrackTypes: valid values are %v", v, AllowedEventManagerEventsCreateV2TrackTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerEventsCreateV2TrackTypes) IsValid() bool {
	for _, existing := range AllowedEventManagerEventsCreateV2TrackTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_events_create_v2_track_types value
func (v EventManagerEventsCreateV2TrackTypes) Ptr() *EventManagerEventsCreateV2TrackTypes {
	return &v
}
