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

// QianchuanAwemeAuthListGetV10FilteringAuthStatus
type QianchuanAwemeAuthListGetV10FilteringAuthStatus string

// List of qianchuan_aweme_auth_list_get_v1.0_filtering_auth_status
const (
	ALL_QianchuanAwemeAuthListGetV10FilteringAuthStatus       QianchuanAwemeAuthListGetV10FilteringAuthStatus = "ALL"
	EFFECTIVE_QianchuanAwemeAuthListGetV10FilteringAuthStatus QianchuanAwemeAuthListGetV10FilteringAuthStatus = "EFFECTIVE"
	EXPIRED_QianchuanAwemeAuthListGetV10FilteringAuthStatus   QianchuanAwemeAuthListGetV10FilteringAuthStatus = "EXPIRED"
	WAIT_BIND_QianchuanAwemeAuthListGetV10FilteringAuthStatus QianchuanAwemeAuthListGetV10FilteringAuthStatus = "WAIT_BIND"
)

// All allowed values of QianchuanAwemeAuthListGetV10FilteringAuthStatus enum
var AllowedQianchuanAwemeAuthListGetV10FilteringAuthStatusEnumValues = []QianchuanAwemeAuthListGetV10FilteringAuthStatus{
	"ALL",
	"EFFECTIVE",
	"EXPIRED",
	"WAIT_BIND",
}

// NewQianchuanAwemeAuthListGetV10FilteringAuthStatusFromValue returns a pointer to a valid QianchuanAwemeAuthListGetV10FilteringAuthStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeAuthListGetV10FilteringAuthStatusFromValue(v string) (*QianchuanAwemeAuthListGetV10FilteringAuthStatus, error) {
	ev := QianchuanAwemeAuthListGetV10FilteringAuthStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeAuthListGetV10FilteringAuthStatus: valid values are %v", v, AllowedQianchuanAwemeAuthListGetV10FilteringAuthStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeAuthListGetV10FilteringAuthStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeAuthListGetV10FilteringAuthStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_auth_list_get_v1.0_filtering_auth_status value
func (v QianchuanAwemeAuthListGetV10FilteringAuthStatus) Ptr() *QianchuanAwemeAuthListGetV10FilteringAuthStatus {
	return &v
}
