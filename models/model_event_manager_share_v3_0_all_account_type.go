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

// EventManagerShareV30AllAccountType
type EventManagerShareV30AllAccountType string

// List of event_manager_share_v3.0_all_account_type
const (
	AD_EventManagerShareV30AllAccountType EventManagerShareV30AllAccountType = "AD"
)

// All allowed values of EventManagerShareV30AllAccountType enum
var AllowedEventManagerShareV30AllAccountTypeEnumValues = []EventManagerShareV30AllAccountType{
	"AD",
}

// NewEventManagerShareV30AllAccountTypeFromValue returns a pointer to a valid EventManagerShareV30AllAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareV30AllAccountTypeFromValue(v string) (*EventManagerShareV30AllAccountType, error) {
	ev := EventManagerShareV30AllAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareV30AllAccountType: valid values are %v", v, AllowedEventManagerShareV30AllAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareV30AllAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareV30AllAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_v3.0_all_account_type value
func (v EventManagerShareV30AllAccountType) Ptr() *EventManagerShareV30AllAccountType {
	return &v
}
