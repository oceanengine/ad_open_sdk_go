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

// QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus
type QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus string

// List of qianchuan_aweme_authorized_get_v1.0_data_aweme_id_list_aweme_status
const (
	ANCHOR_FORBID_QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus                  QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus = "ANCHOR_FORBID"
	ANCHOR_REACH_UPPER_LIMIT_TODAY_QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus = "ANCHOR_REACH_UPPER_LIMIT_TODAY"
	NORMAL_QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus                         QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus = "NORMAL"
	UNKNOWN_QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus                        QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus = "UNKNOWN"
)

// All allowed values of QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus enum
var AllowedQianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatusEnumValues = []QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus{
	"ANCHOR_FORBID",
	"ANCHOR_REACH_UPPER_LIMIT_TODAY",
	"NORMAL",
	"UNKNOWN",
}

// NewQianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatusFromValue returns a pointer to a valid QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatusFromValue(v string) (*QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus, error) {
	ev := QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus: valid values are %v", v, AllowedQianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_authorized_get_v1.0_data_aweme_id_list_aweme_status value
func (v QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus) Ptr() *QianchuanAwemeAuthorizedGetV10DataAwemeIdListAwemeStatus {
	return &v
}
