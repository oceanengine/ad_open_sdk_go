/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaDetailGetV2FilteringStatus
type DpaDetailGetV2FilteringStatus int64

// List of dpa_detail_get_v2_filtering_status
const (
	Enum_0_DpaDetailGetV2FilteringStatus DpaDetailGetV2FilteringStatus = 0
	Enum_1_DpaDetailGetV2FilteringStatus DpaDetailGetV2FilteringStatus = 1
)

// All allowed values of DpaDetailGetV2FilteringStatus enum
var AllowedDpaDetailGetV2FilteringStatusEnumValues = []DpaDetailGetV2FilteringStatus{
	0,
	1,
}

// NewDpaDetailGetV2FilteringStatusFromValue returns a pointer to a valid DpaDetailGetV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaDetailGetV2FilteringStatusFromValue(v int64) (*DpaDetailGetV2FilteringStatus, error) {
	ev := DpaDetailGetV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaDetailGetV2FilteringStatus: valid values are %v", v, AllowedDpaDetailGetV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaDetailGetV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedDpaDetailGetV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_detail_get_v2_filtering_status value
func (v DpaDetailGetV2FilteringStatus) Ptr() *DpaDetailGetV2FilteringStatus {
	return &v
}
