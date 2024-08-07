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

// EventManagerShareGetV30DataListShareMode
type EventManagerShareGetV30DataListShareMode string

// List of event_manager_share_get_v3.0_data_list_share_mode
const (
	ALL_EventManagerShareGetV30DataListShareMode  EventManagerShareGetV30DataListShareMode = "ALL"
	PART_EventManagerShareGetV30DataListShareMode EventManagerShareGetV30DataListShareMode = "PART"
)

// All allowed values of EventManagerShareGetV30DataListShareMode enum
var AllowedEventManagerShareGetV30DataListShareModeEnumValues = []EventManagerShareGetV30DataListShareMode{
	"ALL",
	"PART",
}

// NewEventManagerShareGetV30DataListShareModeFromValue returns a pointer to a valid EventManagerShareGetV30DataListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareGetV30DataListShareModeFromValue(v string) (*EventManagerShareGetV30DataListShareMode, error) {
	ev := EventManagerShareGetV30DataListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareGetV30DataListShareMode: valid values are %v", v, AllowedEventManagerShareGetV30DataListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareGetV30DataListShareMode) IsValid() bool {
	for _, existing := range AllowedEventManagerShareGetV30DataListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_get_v3.0_data_list_share_mode value
func (v EventManagerShareGetV30DataListShareMode) Ptr() *EventManagerShareGetV30DataListShareMode {
	return &v
}
