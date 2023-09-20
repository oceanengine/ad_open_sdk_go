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

// DpaAssetsListV2DataListStatus
type DpaAssetsListV2DataListStatus string

// List of dpa_assets_list_v2_data_list_status
const (
	ENABLE_DpaAssetsListV2DataListStatus  DpaAssetsListV2DataListStatus = "ENABLE"
	DISABLE_DpaAssetsListV2DataListStatus DpaAssetsListV2DataListStatus = "DISABLE"
)

// All allowed values of DpaAssetsListV2DataListStatus enum
var AllowedDpaAssetsListV2DataListStatusEnumValues = []DpaAssetsListV2DataListStatus{
	"ENABLE",
	"DISABLE",
}

// NewDpaAssetsListV2DataListStatusFromValue returns a pointer to a valid DpaAssetsListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsListV2DataListStatusFromValue(v string) (*DpaAssetsListV2DataListStatus, error) {
	ev := DpaAssetsListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsListV2DataListStatus: valid values are %v", v, AllowedDpaAssetsListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetsListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_list_v2_data_list_status value
func (v DpaAssetsListV2DataListStatus) Ptr() *DpaAssetsListV2DataListStatus {
	return &v
}
