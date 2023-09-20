/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaAssetsDetailReadV2DataListAssetType
type DpaAssetsDetailReadV2DataListAssetType string

// List of dpa_assets_detail_read_v2_data_list_asset_type
const (
	AUTO_DpaAssetsDetailReadV2DataListAssetType DpaAssetsDetailReadV2DataListAssetType = "AUTO"
)

// All allowed values of DpaAssetsDetailReadV2DataListAssetType enum
var AllowedDpaAssetsDetailReadV2DataListAssetTypeEnumValues = []DpaAssetsDetailReadV2DataListAssetType{
	"AUTO",
}

// NewDpaAssetsDetailReadV2DataListAssetTypeFromValue returns a pointer to a valid DpaAssetsDetailReadV2DataListAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsDetailReadV2DataListAssetTypeFromValue(v string) (*DpaAssetsDetailReadV2DataListAssetType, error) {
	ev := DpaAssetsDetailReadV2DataListAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsDetailReadV2DataListAssetType: valid values are %v", v, AllowedDpaAssetsDetailReadV2DataListAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsDetailReadV2DataListAssetType) IsValid() bool {
	for _, existing := range AllowedDpaAssetsDetailReadV2DataListAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_detail_read_v2_data_list_asset_type value
func (v DpaAssetsDetailReadV2DataListAssetType) Ptr() *DpaAssetsDetailReadV2DataListAssetType {
	return &v
}
