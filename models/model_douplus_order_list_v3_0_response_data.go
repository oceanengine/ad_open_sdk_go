/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30ResponseData
type DouplusOrderListV30ResponseData struct {
	//
	OrderList []*DouplusOrderListV30ResponseDataOrderListInner `json:"order_list,omitempty"`
	PageInfo  *DouplusOrderListV30ResponseDataPageInfo         `json:"page_info,omitempty"`
}
