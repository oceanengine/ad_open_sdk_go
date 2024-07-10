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

// DmpDataSourceCreateV2FileStorageType
type DmpDataSourceCreateV2FileStorageType int64

// List of dmp_data_source_create_v2_file_storage_type
const (
	Enum_0_DmpDataSourceCreateV2FileStorageType DmpDataSourceCreateV2FileStorageType = 0
)

// All allowed values of DmpDataSourceCreateV2FileStorageType enum
var AllowedDmpDataSourceCreateV2FileStorageTypeEnumValues = []DmpDataSourceCreateV2FileStorageType{
	0,
}

// NewDmpDataSourceCreateV2FileStorageTypeFromValue returns a pointer to a valid DmpDataSourceCreateV2FileStorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceCreateV2FileStorageTypeFromValue(v int64) (*DmpDataSourceCreateV2FileStorageType, error) {
	ev := DmpDataSourceCreateV2FileStorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceCreateV2FileStorageType: valid values are %v", v, AllowedDmpDataSourceCreateV2FileStorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceCreateV2FileStorageType) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceCreateV2FileStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_create_v2_file_storage_type value
func (v DmpDataSourceCreateV2FileStorageType) Ptr() *DmpDataSourceCreateV2FileStorageType {
	return &v
}
