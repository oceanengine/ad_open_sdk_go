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

// AdGetV2DataOptStatus
type AdGetV2DataOptStatus string

// List of ad_get_v2_data_opt_status
const (
	AD_STATUS_DISABLE_BY_QUOTA_AdGetV2DataOptStatus AdGetV2DataOptStatus = "AD_STATUS_DISABLE_BY_QUOTA"
	AD_STATUS_DISABLE_AdGetV2DataOptStatus          AdGetV2DataOptStatus = "AD_STATUS_DISABLE"
	AD_STATUS_FROZEN_AdGetV2DataOptStatus           AdGetV2DataOptStatus = "AD_STATUS_FROZEN"
	AD_STATUS_ENABLE_AdGetV2DataOptStatus           AdGetV2DataOptStatus = "AD_STATUS_ENABLE"
)

// All allowed values of AdGetV2DataOptStatus enum
var AllowedAdGetV2DataOptStatusEnumValues = []AdGetV2DataOptStatus{
	"AD_STATUS_DISABLE_BY_QUOTA",
	"AD_STATUS_DISABLE",
	"AD_STATUS_FROZEN",
	"AD_STATUS_ENABLE",
}

// NewAdGetV2DataOptStatusFromValue returns a pointer to a valid AdGetV2DataOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataOptStatusFromValue(v string) (*AdGetV2DataOptStatus, error) {
	ev := AdGetV2DataOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataOptStatus: valid values are %v", v, AllowedAdGetV2DataOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataOptStatus) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_opt_status value
func (v AdGetV2DataOptStatus) Ptr() *AdGetV2DataOptStatus {
	return &v
}
