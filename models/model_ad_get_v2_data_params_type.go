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

// AdGetV2DataParamsType
type AdGetV2DataParamsType string

// List of ad_get_v2_data_params_type
const (
	CUSTOM_AdGetV2DataParamsType AdGetV2DataParamsType = "CUSTOM"
	DPA_AdGetV2DataParamsType    AdGetV2DataParamsType = "DPA"
)

// All allowed values of AdGetV2DataParamsType enum
var AllowedAdGetV2DataParamsTypeEnumValues = []AdGetV2DataParamsType{
	"CUSTOM",
	"DPA",
}

// NewAdGetV2DataParamsTypeFromValue returns a pointer to a valid AdGetV2DataParamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataParamsTypeFromValue(v string) (*AdGetV2DataParamsType, error) {
	ev := AdGetV2DataParamsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataParamsType: valid values are %v", v, AllowedAdGetV2DataParamsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataParamsType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataParamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_params_type value
func (v AdGetV2DataParamsType) Ptr() *AdGetV2DataParamsType {
	return &v
}
