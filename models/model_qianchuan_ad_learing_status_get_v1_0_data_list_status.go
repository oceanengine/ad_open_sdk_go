/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdLearingStatusGetV10DataListStatus
type QianchuanAdLearingStatusGetV10DataListStatus string

// List of qianchuan_ad_learing_status_get_v1.0_data_list_status
const (
	LEARNING_QianchuanAdLearingStatusGetV10DataListStatus     QianchuanAdLearingStatusGetV10DataListStatus = "LEARNING"
	LEARNED_QianchuanAdLearingStatusGetV10DataListStatus      QianchuanAdLearingStatusGetV10DataListStatus = "LEARNED"
	LEARN_FAILED_QianchuanAdLearingStatusGetV10DataListStatus QianchuanAdLearingStatusGetV10DataListStatus = "LEARN_FAILED"
	DEFAULT_QianchuanAdLearingStatusGetV10DataListStatus      QianchuanAdLearingStatusGetV10DataListStatus = "DEFAULT"
)

// All allowed values of QianchuanAdLearingStatusGetV10DataListStatus enum
var AllowedQianchuanAdLearingStatusGetV10DataListStatusEnumValues = []QianchuanAdLearingStatusGetV10DataListStatus{
	"LEARNING",
	"LEARNED",
	"LEARN_FAILED",
	"DEFAULT",
}

// NewQianchuanAdLearingStatusGetV10DataListStatusFromValue returns a pointer to a valid QianchuanAdLearingStatusGetV10DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdLearingStatusGetV10DataListStatusFromValue(v string) (*QianchuanAdLearingStatusGetV10DataListStatus, error) {
	ev := QianchuanAdLearingStatusGetV10DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdLearingStatusGetV10DataListStatus: valid values are %v", v, AllowedQianchuanAdLearingStatusGetV10DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdLearingStatusGetV10DataListStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdLearingStatusGetV10DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_learing_status_get_v1.0_data_list_status value
func (v QianchuanAdLearingStatusGetV10DataListStatus) Ptr() *QianchuanAdLearingStatusGetV10DataListStatus {
	return &v
}
