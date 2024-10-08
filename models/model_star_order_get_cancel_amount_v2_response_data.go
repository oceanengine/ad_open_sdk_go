/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetCancelAmountV2ResponseData
type StarOrderGetCancelAmountV2ResponseData struct {
	// 赔付金额（元）
	Cost *int64 `json:"cost,omitempty"`
	// 赔付描述
	Desc *string `json:"desc,omitempty"`
	// 精确赔付金额（元）
	PreciseCost *float64 `json:"precise_cost,omitempty"`
}
