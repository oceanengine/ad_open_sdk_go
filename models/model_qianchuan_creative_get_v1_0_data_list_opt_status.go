/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10DataListOptStatus
type QianchuanCreativeGetV10DataListOptStatus string

// List of qianchuan_creative_get_v1.0_data_list_opt_status
const (
	DELETE_QianchuanCreativeGetV10DataListOptStatus  QianchuanCreativeGetV10DataListOptStatus = "DELETE"
	DISABLE_QianchuanCreativeGetV10DataListOptStatus QianchuanCreativeGetV10DataListOptStatus = "DISABLE"
	ENABLE_QianchuanCreativeGetV10DataListOptStatus  QianchuanCreativeGetV10DataListOptStatus = "ENABLE"
)

// All allowed values of QianchuanCreativeGetV10DataListOptStatus enum
var AllowedQianchuanCreativeGetV10DataListOptStatusEnumValues = []QianchuanCreativeGetV10DataListOptStatus{
	"DELETE",
	"DISABLE",
	"ENABLE",
}

// NewQianchuanCreativeGetV10DataListOptStatusFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListOptStatusFromValue(v string) (*QianchuanCreativeGetV10DataListOptStatus, error) {
	ev := QianchuanCreativeGetV10DataListOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListOptStatus: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListOptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_opt_status value
func (v QianchuanCreativeGetV10DataListOptStatus) Ptr() *QianchuanCreativeGetV10DataListOptStatus {
	return &v
}
