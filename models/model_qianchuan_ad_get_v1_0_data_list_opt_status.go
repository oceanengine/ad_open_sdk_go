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

// QianchuanAdGetV10DataListOptStatus
type QianchuanAdGetV10DataListOptStatus string

// List of qianchuan_ad_get_v1.0_data_list_opt_status
const (
	DELETE_QianchuanAdGetV10DataListOptStatus         QianchuanAdGetV10DataListOptStatus = "DELETE"
	DISABLE_QianchuanAdGetV10DataListOptStatus        QianchuanAdGetV10DataListOptStatus = "DISABLE"
	ENABLE_QianchuanAdGetV10DataListOptStatus         QianchuanAdGetV10DataListOptStatus = "ENABLE"
	QUOTA_DISABLE_QianchuanAdGetV10DataListOptStatus  QianchuanAdGetV10DataListOptStatus = "QUOTA_DISABLE"
	ROI2_DISABLE_QianchuanAdGetV10DataListOptStatus   QianchuanAdGetV10DataListOptStatus = "ROI2_DISABLE"
	SYSTEM_DISABLE_QianchuanAdGetV10DataListOptStatus QianchuanAdGetV10DataListOptStatus = "SYSTEM_DISABLE"
)

// All allowed values of QianchuanAdGetV10DataListOptStatus enum
var AllowedQianchuanAdGetV10DataListOptStatusEnumValues = []QianchuanAdGetV10DataListOptStatus{
	"DELETE",
	"DISABLE",
	"ENABLE",
	"QUOTA_DISABLE",
	"ROI2_DISABLE",
	"SYSTEM_DISABLE",
}

// NewQianchuanAdGetV10DataListOptStatusFromValue returns a pointer to a valid QianchuanAdGetV10DataListOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListOptStatusFromValue(v string) (*QianchuanAdGetV10DataListOptStatus, error) {
	ev := QianchuanAdGetV10DataListOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListOptStatus: valid values are %v", v, AllowedQianchuanAdGetV10DataListOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListOptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_opt_status value
func (v QianchuanAdGetV10DataListOptStatus) Ptr() *QianchuanAdGetV10DataListOptStatus {
	return &v
}
