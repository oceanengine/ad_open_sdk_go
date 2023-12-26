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

// SubscribeAccountsListV30DataAdvertisersStatus
type SubscribeAccountsListV30DataAdvertisersStatus string

// List of subscribe_accounts_list_v3.0_data_advertisers_status
const (
	OK_SubscribeAccountsListV30DataAdvertisersStatus           SubscribeAccountsListV30DataAdvertisersStatus = "OK"
	PENDING_SubscribeAccountsListV30DataAdvertisersStatus      SubscribeAccountsListV30DataAdvertisersStatus = "PENDING"
	UNAUTHORIZED_SubscribeAccountsListV30DataAdvertisersStatus SubscribeAccountsListV30DataAdvertisersStatus = "UNAUTHORIZED"
	UNKNOWN_SubscribeAccountsListV30DataAdvertisersStatus      SubscribeAccountsListV30DataAdvertisersStatus = "UNKNOWN"
)

// All allowed values of SubscribeAccountsListV30DataAdvertisersStatus enum
var AllowedSubscribeAccountsListV30DataAdvertisersStatusEnumValues = []SubscribeAccountsListV30DataAdvertisersStatus{
	"OK",
	"PENDING",
	"UNAUTHORIZED",
	"UNKNOWN",
}

// NewSubscribeAccountsListV30DataAdvertisersStatusFromValue returns a pointer to a valid SubscribeAccountsListV30DataAdvertisersStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscribeAccountsListV30DataAdvertisersStatusFromValue(v string) (*SubscribeAccountsListV30DataAdvertisersStatus, error) {
	ev := SubscribeAccountsListV30DataAdvertisersStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscribeAccountsListV30DataAdvertisersStatus: valid values are %v", v, AllowedSubscribeAccountsListV30DataAdvertisersStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscribeAccountsListV30DataAdvertisersStatus) IsValid() bool {
	for _, existing := range AllowedSubscribeAccountsListV30DataAdvertisersStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to subscribe_accounts_list_v3.0_data_advertisers_status value
func (v SubscribeAccountsListV30DataAdvertisersStatus) Ptr() *SubscribeAccountsListV30DataAdvertisersStatus {
	return &v
}
