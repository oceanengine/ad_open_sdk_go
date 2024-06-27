/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus
type QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus string

// List of qianchuan_tools_smart_boost_ad_boost_status_get_v1.0_data_raise_details_status
const (
	CANNOT_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "CANNOT_RAISE"
	CAN_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus    QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "CAN_RAISE"
	FINISH_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "FINISH_RAISE"
	HASPRE_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "HASPRE_RAISE"
	RAISE_FAILED_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "RAISE_FAILED"
	RAISING_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus      QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "RAISING"
)

// All allowed values of QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus enum
var AllowedQianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatusEnumValues = []QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus{
	"CANNOT_RAISE",
	"CAN_RAISE",
	"FINISH_RAISE",
	"HASPRE_RAISE",
	"RAISE_FAILED",
	"RAISING",
}

// NewQianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatusFromValue returns a pointer to a valid QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatusFromValue(v string) (*QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus, error) {
	ev := QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus: valid values are %v", v, AllowedQianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_smart_boost_ad_boost_status_get_v1.0_data_raise_details_status value
func (v QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus) Ptr() *QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus {
	return &v
}
