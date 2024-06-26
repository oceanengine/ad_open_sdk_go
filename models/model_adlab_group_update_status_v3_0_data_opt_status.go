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

// AdlabGroupUpdateStatusV30DataOptStatus
type AdlabGroupUpdateStatusV30DataOptStatus string

// List of adlab_group_update_status_v3.0_data_opt_status
const (
	DISABLED_AdlabGroupUpdateStatusV30DataOptStatus AdlabGroupUpdateStatusV30DataOptStatus = "DISABLED"
	ENABLED_AdlabGroupUpdateStatusV30DataOptStatus  AdlabGroupUpdateStatusV30DataOptStatus = "ENABLED"
)

// All allowed values of AdlabGroupUpdateStatusV30DataOptStatus enum
var AllowedAdlabGroupUpdateStatusV30DataOptStatusEnumValues = []AdlabGroupUpdateStatusV30DataOptStatus{
	"DISABLED",
	"ENABLED",
}

// NewAdlabGroupUpdateStatusV30DataOptStatusFromValue returns a pointer to a valid AdlabGroupUpdateStatusV30DataOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateStatusV30DataOptStatusFromValue(v string) (*AdlabGroupUpdateStatusV30DataOptStatus, error) {
	ev := AdlabGroupUpdateStatusV30DataOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateStatusV30DataOptStatus: valid values are %v", v, AllowedAdlabGroupUpdateStatusV30DataOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateStatusV30DataOptStatus) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateStatusV30DataOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_status_v3.0_data_opt_status value
func (v AdlabGroupUpdateStatusV30DataOptStatus) Ptr() *AdlabGroupUpdateStatusV30DataOptStatus {
	return &v
}
