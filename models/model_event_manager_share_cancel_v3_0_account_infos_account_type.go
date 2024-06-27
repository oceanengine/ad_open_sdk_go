/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerShareCancelV30AccountInfosAccountType
type EventManagerShareCancelV30AccountInfosAccountType string

// List of event_manager_share_cancel_v3.0_account_infos_account_type
const (
	AD_EventManagerShareCancelV30AccountInfosAccountType EventManagerShareCancelV30AccountInfosAccountType = "AD"
)

// All allowed values of EventManagerShareCancelV30AccountInfosAccountType enum
var AllowedEventManagerShareCancelV30AccountInfosAccountTypeEnumValues = []EventManagerShareCancelV30AccountInfosAccountType{
	"AD",
}

// NewEventManagerShareCancelV30AccountInfosAccountTypeFromValue returns a pointer to a valid EventManagerShareCancelV30AccountInfosAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareCancelV30AccountInfosAccountTypeFromValue(v string) (*EventManagerShareCancelV30AccountInfosAccountType, error) {
	ev := EventManagerShareCancelV30AccountInfosAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareCancelV30AccountInfosAccountType: valid values are %v", v, AllowedEventManagerShareCancelV30AccountInfosAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareCancelV30AccountInfosAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareCancelV30AccountInfosAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_cancel_v3.0_account_infos_account_type value
func (v EventManagerShareCancelV30AccountInfosAccountType) Ptr() *EventManagerShareCancelV30AccountInfosAccountType {
	return &v
}
