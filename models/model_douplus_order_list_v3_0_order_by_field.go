/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderListV30OrderByField
type DouplusOrderListV30OrderByField string

// List of douplus_order_list_v3.0_order_by_field
const (
	CREATE_TIME_DouplusOrderListV30OrderByField DouplusOrderListV30OrderByField = "CREATE_TIME"
	ORDER_ID_DouplusOrderListV30OrderByField    DouplusOrderListV30OrderByField = "ORDER_ID"
)

// All allowed values of DouplusOrderListV30OrderByField enum
var AllowedDouplusOrderListV30OrderByFieldEnumValues = []DouplusOrderListV30OrderByField{
	"CREATE_TIME",
	"ORDER_ID",
}

// NewDouplusOrderListV30OrderByFieldFromValue returns a pointer to a valid DouplusOrderListV30OrderByField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30OrderByFieldFromValue(v string) (*DouplusOrderListV30OrderByField, error) {
	ev := DouplusOrderListV30OrderByField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30OrderByField: valid values are %v", v, AllowedDouplusOrderListV30OrderByFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30OrderByField) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30OrderByFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_order_by_field value
func (v DouplusOrderListV30OrderByField) Ptr() *DouplusOrderListV30OrderByField {
	return &v
}
