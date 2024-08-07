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

// ReportAdGetV2FilteringDeliveryMode
type ReportAdGetV2FilteringDeliveryMode string

// List of report_ad_get_v2_filtering_delivery_mode
const (
	STANDARD_ReportAdGetV2FilteringDeliveryMode   ReportAdGetV2FilteringDeliveryMode = "STANDARD"
	ADLAB_FREE_ReportAdGetV2FilteringDeliveryMode ReportAdGetV2FilteringDeliveryMode = "ADLAB_FREE"
)

// All allowed values of ReportAdGetV2FilteringDeliveryMode enum
var AllowedReportAdGetV2FilteringDeliveryModeEnumValues = []ReportAdGetV2FilteringDeliveryMode{
	"STANDARD",
	"ADLAB_FREE",
}

// NewReportAdGetV2FilteringDeliveryModeFromValue returns a pointer to a valid ReportAdGetV2FilteringDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringDeliveryModeFromValue(v string) (*ReportAdGetV2FilteringDeliveryMode, error) {
	ev := ReportAdGetV2FilteringDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringDeliveryMode: valid values are %v", v, AllowedReportAdGetV2FilteringDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringDeliveryMode) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_delivery_mode value
func (v ReportAdGetV2FilteringDeliveryMode) Ptr() *ReportAdGetV2FilteringDeliveryMode {
	return &v
}
