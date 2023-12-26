/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType
type EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType string

// List of event_manager_event_configs_get_v2_data_event_configs_properties_variable_type
const (
	INT_EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType    EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType = "Int"
	BIGINT_EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType = "Bigint"
	STRING_EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType = "String"
	FLOAT_EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType  EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType = "Float"
	BOOL_EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType   EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType = "Bool"
	ENUM_EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType   EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType = "Enum"
)

// All allowed values of EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType enum
var AllowedEventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableTypeEnumValues = []EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType{
	"Int",
	"Bigint",
	"String",
	"Float",
	"Bool",
	"Enum",
}

// NewEventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableTypeFromValue returns a pointer to a valid EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableTypeFromValue(v string) (*EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType, error) {
	ev := EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType: valid values are %v", v, AllowedEventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType) IsValid() bool {
	for _, existing := range AllowedEventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_event_configs_get_v2_data_event_configs_properties_variable_type value
func (v EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType) Ptr() *EventManagerEventConfigsGetV2DataEventConfigsPropertiesVariableType {
	return &v
}
