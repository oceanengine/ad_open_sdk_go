/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataAddAmountInfo 续费信息
type QianchuanAwemeOrderDetailGetV10ResponseDataAddAmountInfo struct {
	// 续费订单金额之和
	AddAmount *float64 `json:"add_amount,omitempty"`
	// 续费次数
	AddAmountCnt *int64 `json:"add_amount_cnt,omitempty"`
	// 续费订单时长之和 - 短视频：xx天 - 直播：xx小时
	AddDeliveryTime *float64 `json:"add_delivery_time,omitempty"`
}