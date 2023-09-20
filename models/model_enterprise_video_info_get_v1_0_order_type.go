/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseVideoInfoGetV10OrderType
type EnterpriseVideoInfoGetV10OrderType string

// List of enterprise_video_info_get_v1.0_order_type
const (
	ASC_EnterpriseVideoInfoGetV10OrderType  EnterpriseVideoInfoGetV10OrderType = "ASC"
	DESC_EnterpriseVideoInfoGetV10OrderType EnterpriseVideoInfoGetV10OrderType = "DESC"
)

// All allowed values of EnterpriseVideoInfoGetV10OrderType enum
var AllowedEnterpriseVideoInfoGetV10OrderTypeEnumValues = []EnterpriseVideoInfoGetV10OrderType{
	"ASC",
	"DESC",
}

// NewEnterpriseVideoInfoGetV10OrderTypeFromValue returns a pointer to a valid EnterpriseVideoInfoGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseVideoInfoGetV10OrderTypeFromValue(v string) (*EnterpriseVideoInfoGetV10OrderType, error) {
	ev := EnterpriseVideoInfoGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseVideoInfoGetV10OrderType: valid values are %v", v, AllowedEnterpriseVideoInfoGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseVideoInfoGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedEnterpriseVideoInfoGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_video_info_get_v1.0_order_type value
func (v EnterpriseVideoInfoGetV10OrderType) Ptr() *EnterpriseVideoInfoGetV10OrderType {
	return &v
}
