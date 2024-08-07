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

// DpaProductStatusBatchUpdateV2OptStatus
type DpaProductStatusBatchUpdateV2OptStatus string

// List of dpa_product_status_batch_update_v2_opt_status
const (
	ENABLE_DpaProductStatusBatchUpdateV2OptStatus  DpaProductStatusBatchUpdateV2OptStatus = "ENABLE"
	DISABLE_DpaProductStatusBatchUpdateV2OptStatus DpaProductStatusBatchUpdateV2OptStatus = "DISABLE"
)

// All allowed values of DpaProductStatusBatchUpdateV2OptStatus enum
var AllowedDpaProductStatusBatchUpdateV2OptStatusEnumValues = []DpaProductStatusBatchUpdateV2OptStatus{
	"ENABLE",
	"DISABLE",
}

// NewDpaProductStatusBatchUpdateV2OptStatusFromValue returns a pointer to a valid DpaProductStatusBatchUpdateV2OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaProductStatusBatchUpdateV2OptStatusFromValue(v string) (*DpaProductStatusBatchUpdateV2OptStatus, error) {
	ev := DpaProductStatusBatchUpdateV2OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaProductStatusBatchUpdateV2OptStatus: valid values are %v", v, AllowedDpaProductStatusBatchUpdateV2OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaProductStatusBatchUpdateV2OptStatus) IsValid() bool {
	for _, existing := range AllowedDpaProductStatusBatchUpdateV2OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_product_status_batch_update_v2_opt_status value
func (v DpaProductStatusBatchUpdateV2OptStatus) Ptr() *DpaProductStatusBatchUpdateV2OptStatus {
	return &v
}
