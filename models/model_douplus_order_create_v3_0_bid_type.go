/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderCreateV30BidType
type DouplusOrderCreateV30BidType string

// List of douplus_order_create_v3.0_bid_type
const (
	AUTO_BID_DouplusOrderCreateV30BidType   DouplusOrderCreateV30BidType = "AUTO_BID"
	MANUAL_BID_DouplusOrderCreateV30BidType DouplusOrderCreateV30BidType = "MANUAL_BID"
)

// Ptr returns reference to douplus_order_create_v3.0_bid_type value
func (v DouplusOrderCreateV30BidType) Ptr() *DouplusOrderCreateV30BidType {
	return &v
}
