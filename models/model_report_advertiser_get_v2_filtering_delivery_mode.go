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

// ReportAdvertiserGetV2FilteringDeliveryMode
type ReportAdvertiserGetV2FilteringDeliveryMode string

// List of report_advertiser_get_v2_filtering_delivery_mode
const (
	ADLAB_FREE_ReportAdvertiserGetV2FilteringDeliveryMode ReportAdvertiserGetV2FilteringDeliveryMode = "ADLAB_FREE"
	STANDARD_ReportAdvertiserGetV2FilteringDeliveryMode   ReportAdvertiserGetV2FilteringDeliveryMode = "STANDARD"
)

// All allowed values of ReportAdvertiserGetV2FilteringDeliveryMode enum
var AllowedReportAdvertiserGetV2FilteringDeliveryModeEnumValues = []ReportAdvertiserGetV2FilteringDeliveryMode{
	"ADLAB_FREE",
	"STANDARD",
}

// NewReportAdvertiserGetV2FilteringDeliveryModeFromValue returns a pointer to a valid ReportAdvertiserGetV2FilteringDeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2FilteringDeliveryModeFromValue(v string) (*ReportAdvertiserGetV2FilteringDeliveryMode, error) {
	ev := ReportAdvertiserGetV2FilteringDeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2FilteringDeliveryMode: valid values are %v", v, AllowedReportAdvertiserGetV2FilteringDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2FilteringDeliveryMode) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2FilteringDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_filtering_delivery_mode value
func (v ReportAdvertiserGetV2FilteringDeliveryMode) Ptr() *ReportAdvertiserGetV2FilteringDeliveryMode {
	return &v
}
