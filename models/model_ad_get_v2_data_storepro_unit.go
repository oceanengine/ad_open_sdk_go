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

// AdGetV2DataStoreproUnit
type AdGetV2DataStoreproUnit string

// List of ad_get_v2_data_storepro_unit
const (
	STORE_ACTIVITY_AdGetV2DataStoreproUnit AdGetV2DataStoreproUnit = "STORE_ACTIVITY"
	STORE_AdGetV2DataStoreproUnit          AdGetV2DataStoreproUnit = "STORE"
)

// All allowed values of AdGetV2DataStoreproUnit enum
var AllowedAdGetV2DataStoreproUnitEnumValues = []AdGetV2DataStoreproUnit{
	"STORE_ACTIVITY",
	"STORE",
}

// NewAdGetV2DataStoreproUnitFromValue returns a pointer to a valid AdGetV2DataStoreproUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataStoreproUnitFromValue(v string) (*AdGetV2DataStoreproUnit, error) {
	ev := AdGetV2DataStoreproUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataStoreproUnit: valid values are %v", v, AllowedAdGetV2DataStoreproUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataStoreproUnit) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataStoreproUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_storepro_unit value
func (v AdGetV2DataStoreproUnit) Ptr() *AdGetV2DataStoreproUnit {
	return &v
}
