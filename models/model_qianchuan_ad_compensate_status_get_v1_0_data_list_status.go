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

// QianchuanAdCompensateStatusGetV10DataListStatus
type QianchuanAdCompensateStatusGetV10DataListStatus string

// List of qianchuan_ad_compensate_status_get_v1.0_data_list_status
const (
	SUCCESS_QianchuanAdCompensateStatusGetV10DataListStatus QianchuanAdCompensateStatusGetV10DataListStatus = "SUCCESS"
	FAILED_QianchuanAdCompensateStatusGetV10DataListStatus  QianchuanAdCompensateStatusGetV10DataListStatus = "FAILED"
)

// All allowed values of QianchuanAdCompensateStatusGetV10DataListStatus enum
var AllowedQianchuanAdCompensateStatusGetV10DataListStatusEnumValues = []QianchuanAdCompensateStatusGetV10DataListStatus{
	"SUCCESS",
	"FAILED",
}

// NewQianchuanAdCompensateStatusGetV10DataListStatusFromValue returns a pointer to a valid QianchuanAdCompensateStatusGetV10DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCompensateStatusGetV10DataListStatusFromValue(v string) (*QianchuanAdCompensateStatusGetV10DataListStatus, error) {
	ev := QianchuanAdCompensateStatusGetV10DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCompensateStatusGetV10DataListStatus: valid values are %v", v, AllowedQianchuanAdCompensateStatusGetV10DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCompensateStatusGetV10DataListStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCompensateStatusGetV10DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_compensate_status_get_v1.0_data_list_status value
func (v QianchuanAdCompensateStatusGetV10DataListStatus) Ptr() *QianchuanAdCompensateStatusGetV10DataListStatus {
	return &v
}
