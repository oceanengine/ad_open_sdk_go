/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderCreateV30BidType
type DouplusOrderCreateV30BidType string

// List of douplus_order_create_v3.0_bid_type
const (
	AUTO_BID_DouplusOrderCreateV30BidType   DouplusOrderCreateV30BidType = "AUTO_BID"
	MANUAL_BID_DouplusOrderCreateV30BidType DouplusOrderCreateV30BidType = "MANUAL_BID"
)

// All allowed values of DouplusOrderCreateV30BidType enum
var AllowedDouplusOrderCreateV30BidTypeEnumValues = []DouplusOrderCreateV30BidType{
	"AUTO_BID",
	"MANUAL_BID",
}

// NewDouplusOrderCreateV30BidTypeFromValue returns a pointer to a valid DouplusOrderCreateV30BidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderCreateV30BidTypeFromValue(v string) (*DouplusOrderCreateV30BidType, error) {
	ev := DouplusOrderCreateV30BidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderCreateV30BidType: valid values are %v", v, AllowedDouplusOrderCreateV30BidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderCreateV30BidType) IsValid() bool {
	for _, existing := range AllowedDouplusOrderCreateV30BidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_create_v3.0_bid_type value
func (v DouplusOrderCreateV30BidType) Ptr() *DouplusOrderCreateV30BidType {
	return &v
}