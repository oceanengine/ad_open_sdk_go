/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetV30DataListStatus
type NativeAnchorGetV30DataListStatus string

// List of native_anchor_get_v3.0_data_list_status
const (
	AUDIT_FAILD_NativeAnchorGetV30DataListStatus   NativeAnchorGetV30DataListStatus = "AUDIT_FAILD"
	AUDIT_SUCCESS_NativeAnchorGetV30DataListStatus NativeAnchorGetV30DataListStatus = "AUDIT_SUCCESS"
	CREATE_NativeAnchorGetV30DataListStatus        NativeAnchorGetV30DataListStatus = "CREATE"
	DELETE_NativeAnchorGetV30DataListStatus        NativeAnchorGetV30DataListStatus = "DELETE"
)

// All allowed values of NativeAnchorGetV30DataListStatus enum
var AllowedNativeAnchorGetV30DataListStatusEnumValues = []NativeAnchorGetV30DataListStatus{
	"AUDIT_FAILD",
	"AUDIT_SUCCESS",
	"CREATE",
	"DELETE",
}

// NewNativeAnchorGetV30DataListStatusFromValue returns a pointer to a valid NativeAnchorGetV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30DataListStatusFromValue(v string) (*NativeAnchorGetV30DataListStatus, error) {
	ev := NativeAnchorGetV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30DataListStatus: valid values are %v", v, AllowedNativeAnchorGetV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_data_list_status value
func (v NativeAnchorGetV30DataListStatus) Ptr() *NativeAnchorGetV30DataListStatus {
	return &v
}
