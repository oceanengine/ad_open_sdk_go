/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderFinishV2Request struct for StarOrderFinishV2Request
type StarOrderFinishV2Request struct {
	// 任务ID
	OrderId int64 `json:"order_id"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
