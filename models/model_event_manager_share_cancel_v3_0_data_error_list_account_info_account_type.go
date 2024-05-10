/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerShareCancelV30DataErrorListAccountInfoAccountType
type EventManagerShareCancelV30DataErrorListAccountInfoAccountType string

// List of event_manager_share_cancel_v3.0_data_error_list_account_info_account_type
const (
	AD_EventManagerShareCancelV30DataErrorListAccountInfoAccountType EventManagerShareCancelV30DataErrorListAccountInfoAccountType = "AD"
)

// All allowed values of EventManagerShareCancelV30DataErrorListAccountInfoAccountType enum
var AllowedEventManagerShareCancelV30DataErrorListAccountInfoAccountTypeEnumValues = []EventManagerShareCancelV30DataErrorListAccountInfoAccountType{
	"AD",
}

// NewEventManagerShareCancelV30DataErrorListAccountInfoAccountTypeFromValue returns a pointer to a valid EventManagerShareCancelV30DataErrorListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareCancelV30DataErrorListAccountInfoAccountTypeFromValue(v string) (*EventManagerShareCancelV30DataErrorListAccountInfoAccountType, error) {
	ev := EventManagerShareCancelV30DataErrorListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareCancelV30DataErrorListAccountInfoAccountType: valid values are %v", v, AllowedEventManagerShareCancelV30DataErrorListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareCancelV30DataErrorListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareCancelV30DataErrorListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_cancel_v3.0_data_error_list_account_info_account_type value
func (v EventManagerShareCancelV30DataErrorListAccountInfoAccountType) Ptr() *EventManagerShareCancelV30DataErrorListAccountInfoAccountType {
	return &v
}
