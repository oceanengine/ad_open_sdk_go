/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30Filter
type DouplusOrderListV30Filter struct {
	OrderCreateTime *DouplusOrderListV30FilterOrderCreateTime `json:"order_create_time,omitempty"`
	//
	OrderId   []int64                             `json:"order_id,omitempty"`
	SceneType *DouplusOrderListV30FilterSceneType `json:"scene_type,omitempty"`
	//
	Status []*DouplusOrderListV30FilterStatus `json:"status,omitempty"`
}
