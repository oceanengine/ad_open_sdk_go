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

// EventManagerShareCancelV30ShareMode
type EventManagerShareCancelV30ShareMode string

// List of event_manager_share_cancel_v3.0_share_mode
const (
	ALL_EventManagerShareCancelV30ShareMode  EventManagerShareCancelV30ShareMode = "ALL"
	PART_EventManagerShareCancelV30ShareMode EventManagerShareCancelV30ShareMode = "PART"
)

// All allowed values of EventManagerShareCancelV30ShareMode enum
var AllowedEventManagerShareCancelV30ShareModeEnumValues = []EventManagerShareCancelV30ShareMode{
	"ALL",
	"PART",
}

// NewEventManagerShareCancelV30ShareModeFromValue returns a pointer to a valid EventManagerShareCancelV30ShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareCancelV30ShareModeFromValue(v string) (*EventManagerShareCancelV30ShareMode, error) {
	ev := EventManagerShareCancelV30ShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareCancelV30ShareMode: valid values are %v", v, AllowedEventManagerShareCancelV30ShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareCancelV30ShareMode) IsValid() bool {
	for _, existing := range AllowedEventManagerShareCancelV30ShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_cancel_v3.0_share_mode value
func (v EventManagerShareCancelV30ShareMode) Ptr() *EventManagerShareCancelV30ShareMode {
	return &v
}
