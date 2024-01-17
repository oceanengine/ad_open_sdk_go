/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderUpdateV2Request struct for StarOrderUpdateV2Request
type StarOrderUpdateV2Request struct {
	// 任务ID
	OrderId          int64                                     `json:"order_id"`
	OrderInfoChanges *StarOrderUpdateV2RequestOrderInfoChanges `json:"order_info_changes,omitempty"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
