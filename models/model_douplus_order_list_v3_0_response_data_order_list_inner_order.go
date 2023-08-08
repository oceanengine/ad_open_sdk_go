/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30ResponseDataOrderListInnerOrder
type DouplusOrderListV30ResponseDataOrderListInnerOrder struct {
	//
	Budget    *int64                                          `json:"budget,omitempty"`
	LiveScene *DouplusOrderListV30DataOrderListOrderLiveScene `json:"live_scene,omitempty"`
	//
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	//
	OrderId   *int64                                          `json:"order_id,omitempty"`
	SceneType *DouplusOrderListV30DataOrderListOrderSceneType `json:"scene_type,omitempty"`
	//
	TaskId     *int64                                           `json:"task_id,omitempty"`
	TaskStatus *DouplusOrderListV30DataOrderListOrderTaskStatus `json:"task_status,omitempty"`
}
