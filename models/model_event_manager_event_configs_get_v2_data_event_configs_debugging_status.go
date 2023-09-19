/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus
type EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus string

// List of event_manager_event_configs_get_v2_data_event_configs_debugging_status
const (
	INACTIVE_EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus = "Inactive"
	ACTIVE_EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus   EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus = "Active"
)

// All allowed values of EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus enum
var AllowedEventManagerEventConfigsGetV2DataEventConfigsDebuggingStatusEnumValues = []EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus{
	"Inactive",
	"Active",
}

// NewEventManagerEventConfigsGetV2DataEventConfigsDebuggingStatusFromValue returns a pointer to a valid EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerEventConfigsGetV2DataEventConfigsDebuggingStatusFromValue(v string) (*EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus, error) {
	ev := EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus: valid values are %v", v, AllowedEventManagerEventConfigsGetV2DataEventConfigsDebuggingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus) IsValid() bool {
	for _, existing := range AllowedEventManagerEventConfigsGetV2DataEventConfigsDebuggingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_event_configs_get_v2_data_event_configs_debugging_status value
func (v EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus) Ptr() *EventManagerEventConfigsGetV2DataEventConfigsDebuggingStatus {
	return &v
}
