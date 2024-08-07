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

// MaterialStatusUpdateV30DataOptStatus
type MaterialStatusUpdateV30DataOptStatus string

// List of material_status_update_v3.0_data_opt_status
const (
	DISABLE_MaterialStatusUpdateV30DataOptStatus MaterialStatusUpdateV30DataOptStatus = "DISABLE"
	ENABLE_MaterialStatusUpdateV30DataOptStatus  MaterialStatusUpdateV30DataOptStatus = "ENABLE"
)

// All allowed values of MaterialStatusUpdateV30DataOptStatus enum
var AllowedMaterialStatusUpdateV30DataOptStatusEnumValues = []MaterialStatusUpdateV30DataOptStatus{
	"DISABLE",
	"ENABLE",
}

// NewMaterialStatusUpdateV30DataOptStatusFromValue returns a pointer to a valid MaterialStatusUpdateV30DataOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMaterialStatusUpdateV30DataOptStatusFromValue(v string) (*MaterialStatusUpdateV30DataOptStatus, error) {
	ev := MaterialStatusUpdateV30DataOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MaterialStatusUpdateV30DataOptStatus: valid values are %v", v, AllowedMaterialStatusUpdateV30DataOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MaterialStatusUpdateV30DataOptStatus) IsValid() bool {
	for _, existing := range AllowedMaterialStatusUpdateV30DataOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to material_status_update_v3.0_data_opt_status value
func (v MaterialStatusUpdateV30DataOptStatus) Ptr() *MaterialStatusUpdateV30DataOptStatus {
	return &v
}
