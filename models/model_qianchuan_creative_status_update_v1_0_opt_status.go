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

// QianchuanCreativeStatusUpdateV10OptStatus
type QianchuanCreativeStatusUpdateV10OptStatus string

// List of qianchuan_creative_status_update_v1.0_opt_status
const (
	REVIVE_QianchuanCreativeStatusUpdateV10OptStatus  QianchuanCreativeStatusUpdateV10OptStatus = "REVIVE"
	ENABLE_QianchuanCreativeStatusUpdateV10OptStatus  QianchuanCreativeStatusUpdateV10OptStatus = "ENABLE"
	DISABLE_QianchuanCreativeStatusUpdateV10OptStatus QianchuanCreativeStatusUpdateV10OptStatus = "DISABLE"
	DELETE_QianchuanCreativeStatusUpdateV10OptStatus  QianchuanCreativeStatusUpdateV10OptStatus = "DELETE"
)

// All allowed values of QianchuanCreativeStatusUpdateV10OptStatus enum
var AllowedQianchuanCreativeStatusUpdateV10OptStatusEnumValues = []QianchuanCreativeStatusUpdateV10OptStatus{
	"REVIVE",
	"ENABLE",
	"DISABLE",
	"DELETE",
}

// NewQianchuanCreativeStatusUpdateV10OptStatusFromValue returns a pointer to a valid QianchuanCreativeStatusUpdateV10OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeStatusUpdateV10OptStatusFromValue(v string) (*QianchuanCreativeStatusUpdateV10OptStatus, error) {
	ev := QianchuanCreativeStatusUpdateV10OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeStatusUpdateV10OptStatus: valid values are %v", v, AllowedQianchuanCreativeStatusUpdateV10OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeStatusUpdateV10OptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeStatusUpdateV10OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_status_update_v1.0_opt_status value
func (v QianchuanCreativeStatusUpdateV10OptStatus) Ptr() *QianchuanCreativeStatusUpdateV10OptStatus {
	return &v
}
