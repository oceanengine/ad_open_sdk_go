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

// DpaAssetsListV2DataListSource
type DpaAssetsListV2DataListSource string

// List of dpa_assets_list_v2_data_list_source
const (
	MANUAL_DpaAssetsListV2DataListSource DpaAssetsListV2DataListSource = "MANUAL"
	AUTO_DpaAssetsListV2DataListSource   DpaAssetsListV2DataListSource = "AUTO"
)

// All allowed values of DpaAssetsListV2DataListSource enum
var AllowedDpaAssetsListV2DataListSourceEnumValues = []DpaAssetsListV2DataListSource{
	"MANUAL",
	"AUTO",
}

// NewDpaAssetsListV2DataListSourceFromValue returns a pointer to a valid DpaAssetsListV2DataListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsListV2DataListSourceFromValue(v string) (*DpaAssetsListV2DataListSource, error) {
	ev := DpaAssetsListV2DataListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsListV2DataListSource: valid values are %v", v, AllowedDpaAssetsListV2DataListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsListV2DataListSource) IsValid() bool {
	for _, existing := range AllowedDpaAssetsListV2DataListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_list_v2_data_list_source value
func (v DpaAssetsListV2DataListSource) Ptr() *DpaAssetsListV2DataListSource {
	return &v
}
