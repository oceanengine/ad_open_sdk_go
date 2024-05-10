/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceActivateType
type AdGetV2DataAudienceActivateType string

// List of ad_get_v2_data_audience_activate_type
const (
	ONE_MONTH_2_THREE_MONTH_AdGetV2DataAudienceActivateType AdGetV2DataAudienceActivateType = "ONE_MONTH_2_THREE_MONTH"
	THREE_MONTH_EAILIER_AdGetV2DataAudienceActivateType     AdGetV2DataAudienceActivateType = "THREE_MONTH_EAILIER"
	UNLIMITED_AdGetV2DataAudienceActivateType               AdGetV2DataAudienceActivateType = "UNLIMITED"
	WITH_IN_A_MONTH_AdGetV2DataAudienceActivateType         AdGetV2DataAudienceActivateType = "WITH_IN_A_MONTH"
)

// All allowed values of AdGetV2DataAudienceActivateType enum
var AllowedAdGetV2DataAudienceActivateTypeEnumValues = []AdGetV2DataAudienceActivateType{
	"ONE_MONTH_2_THREE_MONTH",
	"THREE_MONTH_EAILIER",
	"UNLIMITED",
	"WITH_IN_A_MONTH",
}

// NewAdGetV2DataAudienceActivateTypeFromValue returns a pointer to a valid AdGetV2DataAudienceActivateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceActivateTypeFromValue(v string) (*AdGetV2DataAudienceActivateType, error) {
	ev := AdGetV2DataAudienceActivateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceActivateType: valid values are %v", v, AllowedAdGetV2DataAudienceActivateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceActivateType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceActivateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_activate_type value
func (v AdGetV2DataAudienceActivateType) Ptr() *AdGetV2DataAudienceActivateType {
	return &v
}
