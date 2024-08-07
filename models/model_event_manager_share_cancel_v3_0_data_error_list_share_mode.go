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

// EventManagerShareCancelV30DataErrorListShareMode
type EventManagerShareCancelV30DataErrorListShareMode string

// List of event_manager_share_cancel_v3.0_data_error_list_share_mode
const (
	ALL_EventManagerShareCancelV30DataErrorListShareMode  EventManagerShareCancelV30DataErrorListShareMode = "ALL"
	PART_EventManagerShareCancelV30DataErrorListShareMode EventManagerShareCancelV30DataErrorListShareMode = "PART"
)

// All allowed values of EventManagerShareCancelV30DataErrorListShareMode enum
var AllowedEventManagerShareCancelV30DataErrorListShareModeEnumValues = []EventManagerShareCancelV30DataErrorListShareMode{
	"ALL",
	"PART",
}

// NewEventManagerShareCancelV30DataErrorListShareModeFromValue returns a pointer to a valid EventManagerShareCancelV30DataErrorListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareCancelV30DataErrorListShareModeFromValue(v string) (*EventManagerShareCancelV30DataErrorListShareMode, error) {
	ev := EventManagerShareCancelV30DataErrorListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareCancelV30DataErrorListShareMode: valid values are %v", v, AllowedEventManagerShareCancelV30DataErrorListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareCancelV30DataErrorListShareMode) IsValid() bool {
	for _, existing := range AllowedEventManagerShareCancelV30DataErrorListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_cancel_v3.0_data_error_list_share_mode value
func (v EventManagerShareCancelV30DataErrorListShareMode) Ptr() *EventManagerShareCancelV30DataErrorListShareMode {
	return &v
}
