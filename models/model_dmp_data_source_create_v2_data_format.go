/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DmpDataSourceCreateV2DataFormat
type DmpDataSourceCreateV2DataFormat int64

// List of dmp_data_source_create_v2_data_format
const (
	Enum_0_DmpDataSourceCreateV2DataFormat DmpDataSourceCreateV2DataFormat = 0
)

// All allowed values of DmpDataSourceCreateV2DataFormat enum
var AllowedDmpDataSourceCreateV2DataFormatEnumValues = []DmpDataSourceCreateV2DataFormat{
	0,
}

// NewDmpDataSourceCreateV2DataFormatFromValue returns a pointer to a valid DmpDataSourceCreateV2DataFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceCreateV2DataFormatFromValue(v int64) (*DmpDataSourceCreateV2DataFormat, error) {
	ev := DmpDataSourceCreateV2DataFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceCreateV2DataFormat: valid values are %v", v, AllowedDmpDataSourceCreateV2DataFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceCreateV2DataFormat) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceCreateV2DataFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_create_v2_data_format value
func (v DmpDataSourceCreateV2DataFormat) Ptr() *DmpDataSourceCreateV2DataFormat {
	return &v
}
