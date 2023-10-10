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

// EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType
type EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType string

// List of event_manager_available_events_get_v2_data_event_configs_properties_variable_type
const (
	INT_EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType    EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType = "Int"
	BIGINT_EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType = "Bigint"
	STRING_EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType = "String"
	FLOAT_EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType  EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType = "Float"
	BOOL_EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType   EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType = "Bool"
	ENUM_EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType   EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType = "Enum"
)

// All allowed values of EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType enum
var AllowedEventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableTypeEnumValues = []EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType{
	"Int",
	"Bigint",
	"String",
	"Float",
	"Bool",
	"Enum",
}

// NewEventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableTypeFromValue returns a pointer to a valid EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableTypeFromValue(v string) (*EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType, error) {
	ev := EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType: valid values are %v", v, AllowedEventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType) IsValid() bool {
	for _, existing := range AllowedEventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_available_events_get_v2_data_event_configs_properties_variable_type value
func (v EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType) Ptr() *EventManagerAvailableEventsGetV2DataEventConfigsPropertiesVariableType {
	return &v
}
