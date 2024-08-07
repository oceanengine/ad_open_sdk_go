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

// DouplusOrderListV30OrderByType
type DouplusOrderListV30OrderByType string

// List of douplus_order_list_v3.0_order_by_type
const (
	ASC_DouplusOrderListV30OrderByType  DouplusOrderListV30OrderByType = "ASC"
	DESC_DouplusOrderListV30OrderByType DouplusOrderListV30OrderByType = "DESC"
)

// All allowed values of DouplusOrderListV30OrderByType enum
var AllowedDouplusOrderListV30OrderByTypeEnumValues = []DouplusOrderListV30OrderByType{
	"ASC",
	"DESC",
}

// NewDouplusOrderListV30OrderByTypeFromValue returns a pointer to a valid DouplusOrderListV30OrderByType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30OrderByTypeFromValue(v string) (*DouplusOrderListV30OrderByType, error) {
	ev := DouplusOrderListV30OrderByType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30OrderByType: valid values are %v", v, AllowedDouplusOrderListV30OrderByTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30OrderByType) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30OrderByTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_order_by_type value
func (v DouplusOrderListV30OrderByType) Ptr() *DouplusOrderListV30OrderByType {
	return &v
}
