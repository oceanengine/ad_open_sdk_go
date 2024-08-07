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

// DpaAssetV2ListV2DataAssetListSource
type DpaAssetV2ListV2DataAssetListSource string

// List of dpa_asset_v2_list_v2_data_asset_list_source
const (
	AUTO_DpaAssetV2ListV2DataAssetListSource   DpaAssetV2ListV2DataAssetListSource = "AUTO"
	MANUAL_DpaAssetV2ListV2DataAssetListSource DpaAssetV2ListV2DataAssetListSource = "MANUAL"
)

// All allowed values of DpaAssetV2ListV2DataAssetListSource enum
var AllowedDpaAssetV2ListV2DataAssetListSourceEnumValues = []DpaAssetV2ListV2DataAssetListSource{
	"AUTO",
	"MANUAL",
}

// NewDpaAssetV2ListV2DataAssetListSourceFromValue returns a pointer to a valid DpaAssetV2ListV2DataAssetListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetV2ListV2DataAssetListSourceFromValue(v string) (*DpaAssetV2ListV2DataAssetListSource, error) {
	ev := DpaAssetV2ListV2DataAssetListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetV2ListV2DataAssetListSource: valid values are %v", v, AllowedDpaAssetV2ListV2DataAssetListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetV2ListV2DataAssetListSource) IsValid() bool {
	for _, existing := range AllowedDpaAssetV2ListV2DataAssetListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_asset_v2_list_v2_data_asset_list_source value
func (v DpaAssetV2ListV2DataAssetListSource) Ptr() *DpaAssetV2ListV2DataAssetListSource {
	return &v
}
