/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductStatusBatchUpdateV2ResponseDataErrorListInner struct for DpaProductStatusBatchUpdateV2ResponseDataErrorListInner
type DpaProductStatusBatchUpdateV2ResponseDataErrorListInner struct {
	// 失败原因
	ErrorMsg *string `json:"error_msg,omitempty"`
	// 修改失败的商品id
	ProductId *int64 `json:"product_id,omitempty"`
}
