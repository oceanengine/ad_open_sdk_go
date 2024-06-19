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

// AdGetV2DataDeliveryPhase
type AdGetV2DataDeliveryPhase string

// List of ad_get_v2_data_delivery_phase
const (
	FIRST_PHASE_AdGetV2DataDeliveryPhase  AdGetV2DataDeliveryPhase = "FIRST_PHASE"
	SECOND_PHASE_AdGetV2DataDeliveryPhase AdGetV2DataDeliveryPhase = "SECOND_PHASE"
)

// All allowed values of AdGetV2DataDeliveryPhase enum
var AllowedAdGetV2DataDeliveryPhaseEnumValues = []AdGetV2DataDeliveryPhase{
	"FIRST_PHASE",
	"SECOND_PHASE",
}

// NewAdGetV2DataDeliveryPhaseFromValue returns a pointer to a valid AdGetV2DataDeliveryPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataDeliveryPhaseFromValue(v string) (*AdGetV2DataDeliveryPhase, error) {
	ev := AdGetV2DataDeliveryPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataDeliveryPhase: valid values are %v", v, AllowedAdGetV2DataDeliveryPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataDeliveryPhase) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataDeliveryPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_delivery_phase value
func (v AdGetV2DataDeliveryPhase) Ptr() *AdGetV2DataDeliveryPhase {
	return &v
}
