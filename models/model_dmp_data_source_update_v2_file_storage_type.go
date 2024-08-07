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

// DmpDataSourceUpdateV2FileStorageType
type DmpDataSourceUpdateV2FileStorageType int64

// List of dmp_data_source_update_v2_file_storage_type
const (
	Enum_0_DmpDataSourceUpdateV2FileStorageType DmpDataSourceUpdateV2FileStorageType = 0
)

// All allowed values of DmpDataSourceUpdateV2FileStorageType enum
var AllowedDmpDataSourceUpdateV2FileStorageTypeEnumValues = []DmpDataSourceUpdateV2FileStorageType{
	0,
}

// NewDmpDataSourceUpdateV2FileStorageTypeFromValue returns a pointer to a valid DmpDataSourceUpdateV2FileStorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceUpdateV2FileStorageTypeFromValue(v int64) (*DmpDataSourceUpdateV2FileStorageType, error) {
	ev := DmpDataSourceUpdateV2FileStorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceUpdateV2FileStorageType: valid values are %v", v, AllowedDmpDataSourceUpdateV2FileStorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceUpdateV2FileStorageType) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceUpdateV2FileStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_update_v2_file_storage_type value
func (v DmpDataSourceUpdateV2FileStorageType) Ptr() *DmpDataSourceUpdateV2FileStorageType {
	return &v
}
