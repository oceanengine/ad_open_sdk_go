/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30ResponseDataOrderListInnerAdListInner struct for DouplusOrderListV30ResponseDataOrderListInnerAdListInner
type DouplusOrderListV30ResponseDataOrderListInnerAdListInner struct {
	//
	AdId     *int64                                                            `json:"ad_id,omitempty"`
	AdStatus *DouplusOrderListV30DataOrderListAdListAdStatus                   `json:"ad_status,omitempty"`
	Audience *DouplusOrderListV30ResponseDataOrderListInnerAdListInnerAudience `json:"audience,omitempty"`
	BidMode  *DouplusOrderListV30DataOrderListAdListBidMode                    `json:"bid_mode,omitempty"`
	//
	Budget *int64 `json:"budget,omitempty"`
	//
	CpaBid *int64 `json:"cpa_bid,omitempty"`
	//
	DeliveryTime *float64 `json:"delivery_time,omitempty"`
	//
	ExternalAction *int64 `json:"external_action,omitempty"`
	//
	IsFans *bool `json:"is_fans,omitempty"`
}
