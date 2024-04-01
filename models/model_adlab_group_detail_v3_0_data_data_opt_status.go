/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataOptStatus
type AdlabGroupDetailV30DataDataOptStatus string

// List of adlab_group_detail_v3.0_data_data_opt_status
const (
	DELETED_AdlabGroupDetailV30DataDataOptStatus           AdlabGroupDetailV30DataDataOptStatus = "DELETED"
	DISABLED_AdlabGroupDetailV30DataDataOptStatus          AdlabGroupDetailV30DataDataOptStatus = "DISABLED"
	ENABLED_AdlabGroupDetailV30DataDataOptStatus           AdlabGroupDetailV30DataDataOptStatus = "ENABLED"
	REACH_QUOTA_LIMIT_AdlabGroupDetailV30DataDataOptStatus AdlabGroupDetailV30DataDataOptStatus = "REACH_QUOTA_LIMIT"
)

// All allowed values of AdlabGroupDetailV30DataDataOptStatus enum
var AllowedAdlabGroupDetailV30DataDataOptStatusEnumValues = []AdlabGroupDetailV30DataDataOptStatus{
	"DELETED",
	"DISABLED",
	"ENABLED",
	"REACH_QUOTA_LIMIT",
}

// NewAdlabGroupDetailV30DataDataOptStatusFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataOptStatusFromValue(v string) (*AdlabGroupDetailV30DataDataOptStatus, error) {
	ev := AdlabGroupDetailV30DataDataOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataOptStatus: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataOptStatus) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_opt_status value
func (v AdlabGroupDetailV30DataDataOptStatus) Ptr() *AdlabGroupDetailV30DataDataOptStatus {
	return &v
}
