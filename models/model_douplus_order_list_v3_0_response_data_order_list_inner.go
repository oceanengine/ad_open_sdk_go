/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30ResponseDataOrderListInner struct for DouplusOrderListV30ResponseDataOrderListInner
type DouplusOrderListV30ResponseDataOrderListInner struct {
	//
	AdList []*DouplusOrderListV30ResponseDataOrderListInnerAdListInner `json:"ad_list,omitempty"`
	//
	ItemInfoList []*DouplusOrderListV30ResponseDataOrderListInnerItemInfoListInner `json:"item_info_list,omitempty"`
	LiveRoomInfo *DouplusOrderListV30ResponseDataOrderListInnerLiveRoomInfo        `json:"live_room_info,omitempty"`
	Order        *DouplusOrderListV30ResponseDataOrderListInnerOrder               `json:"order,omitempty"`
}
