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

// DmpDataSourceUpdateV2DataFormat
type DmpDataSourceUpdateV2DataFormat int64

// List of dmp_data_source_update_v2_data_format
const (
	Enum_0_DmpDataSourceUpdateV2DataFormat DmpDataSourceUpdateV2DataFormat = 0
)

// All allowed values of DmpDataSourceUpdateV2DataFormat enum
var AllowedDmpDataSourceUpdateV2DataFormatEnumValues = []DmpDataSourceUpdateV2DataFormat{
	0,
}

// NewDmpDataSourceUpdateV2DataFormatFromValue returns a pointer to a valid DmpDataSourceUpdateV2DataFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpDataSourceUpdateV2DataFormatFromValue(v int64) (*DmpDataSourceUpdateV2DataFormat, error) {
	ev := DmpDataSourceUpdateV2DataFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpDataSourceUpdateV2DataFormat: valid values are %v", v, AllowedDmpDataSourceUpdateV2DataFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpDataSourceUpdateV2DataFormat) IsValid() bool {
	for _, existing := range AllowedDmpDataSourceUpdateV2DataFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_data_source_update_v2_data_format value
func (v DmpDataSourceUpdateV2DataFormat) Ptr() *DmpDataSourceUpdateV2DataFormat {
	return &v
}
