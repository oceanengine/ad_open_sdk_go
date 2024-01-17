/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderDemanderCancelV2Request struct for StarOrderDemanderCancelV2Request
type StarOrderDemanderCancelV2Request struct {
	// 任务ID
	OrderId int64 `json:"order_id"`
	// 其他取消原因 100字内，reason_type=0时需提供
	Reason *string `json:"reason,omitempty"`
	// 取消原因类别 (0)：其他 (1)：信息填错，重新下单 (2)：推广计划改变 (3)：经协商一致 (4)：达人无档期，不接单 (5)：达人不配合 (6)：产出物不符合要求
	ReasonType int64 `json:"reason_type"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
