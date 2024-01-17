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

// AdGetV2DataDpaOpenUrlType
type AdGetV2DataDpaOpenUrlType string

// List of ad_get_v2_data_dpa_open_url_type
const (
	CUSTOM_AdGetV2DataDpaOpenUrlType AdGetV2DataDpaOpenUrlType = "CUSTOM"
	DPA_AdGetV2DataDpaOpenUrlType    AdGetV2DataDpaOpenUrlType = "DPA"
	NONE_AdGetV2DataDpaOpenUrlType   AdGetV2DataDpaOpenUrlType = "NONE"
)

// All allowed values of AdGetV2DataDpaOpenUrlType enum
var AllowedAdGetV2DataDpaOpenUrlTypeEnumValues = []AdGetV2DataDpaOpenUrlType{
	"CUSTOM",
	"DPA",
	"NONE",
}

// NewAdGetV2DataDpaOpenUrlTypeFromValue returns a pointer to a valid AdGetV2DataDpaOpenUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataDpaOpenUrlTypeFromValue(v string) (*AdGetV2DataDpaOpenUrlType, error) {
	ev := AdGetV2DataDpaOpenUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataDpaOpenUrlType: valid values are %v", v, AllowedAdGetV2DataDpaOpenUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataDpaOpenUrlType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataDpaOpenUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_dpa_open_url_type value
func (v AdGetV2DataDpaOpenUrlType) Ptr() *AdGetV2DataDpaOpenUrlType {
	return &v
}
