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

// DpaAssetV2DetailReadV2DataAssetListAssetType
type DpaAssetV2DetailReadV2DataAssetListAssetType string

// List of dpa_asset_v2_detail_read_v2_data_asset_list_asset_type
const (
	AUTO_DpaAssetV2DetailReadV2DataAssetListAssetType DpaAssetV2DetailReadV2DataAssetListAssetType = "AUTO"
)

// All allowed values of DpaAssetV2DetailReadV2DataAssetListAssetType enum
var AllowedDpaAssetV2DetailReadV2DataAssetListAssetTypeEnumValues = []DpaAssetV2DetailReadV2DataAssetListAssetType{
	"AUTO",
}

// NewDpaAssetV2DetailReadV2DataAssetListAssetTypeFromValue returns a pointer to a valid DpaAssetV2DetailReadV2DataAssetListAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetV2DetailReadV2DataAssetListAssetTypeFromValue(v string) (*DpaAssetV2DetailReadV2DataAssetListAssetType, error) {
	ev := DpaAssetV2DetailReadV2DataAssetListAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetV2DetailReadV2DataAssetListAssetType: valid values are %v", v, AllowedDpaAssetV2DetailReadV2DataAssetListAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetV2DetailReadV2DataAssetListAssetType) IsValid() bool {
	for _, existing := range AllowedDpaAssetV2DetailReadV2DataAssetListAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_asset_v2_detail_read_v2_data_asset_list_asset_type value
func (v DpaAssetV2DetailReadV2DataAssetListAssetType) Ptr() *DpaAssetV2DetailReadV2DataAssetListAssetType {
	return &v
}
