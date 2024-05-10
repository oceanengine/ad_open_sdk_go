/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderListV30DataOrderListAdListBidMode
type DouplusOrderListV30DataOrderListAdListBidMode string

// List of douplus_order_list_v3.0_data_order_list_ad_list_bid_mode
const (
	Enum_保播放_DouplusOrderListV30DataOrderListAdListBidMode  DouplusOrderListV30DataOrderListAdListBidMode = "保播放"
	Enum_保转化_DouplusOrderListV30DataOrderListAdListBidMode  DouplusOrderListV30DataOrderListAdListBidMode = "保转化"
	Enum_自动出价_DouplusOrderListV30DataOrderListAdListBidMode DouplusOrderListV30DataOrderListAdListBidMode = "自动出价"
)

// All allowed values of DouplusOrderListV30DataOrderListAdListBidMode enum
var AllowedDouplusOrderListV30DataOrderListAdListBidModeEnumValues = []DouplusOrderListV30DataOrderListAdListBidMode{
	"保播放",
	"保转化",
	"自动出价",
}

// NewDouplusOrderListV30DataOrderListAdListBidModeFromValue returns a pointer to a valid DouplusOrderListV30DataOrderListAdListBidMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30DataOrderListAdListBidModeFromValue(v string) (*DouplusOrderListV30DataOrderListAdListBidMode, error) {
	ev := DouplusOrderListV30DataOrderListAdListBidMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30DataOrderListAdListBidMode: valid values are %v", v, AllowedDouplusOrderListV30DataOrderListAdListBidModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30DataOrderListAdListBidMode) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30DataOrderListAdListBidModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_data_order_list_ad_list_bid_mode value
func (v DouplusOrderListV30DataOrderListAdListBidMode) Ptr() *DouplusOrderListV30DataOrderListAdListBidMode {
	return &v
}
