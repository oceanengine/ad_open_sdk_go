/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DecorationCouponGetV30ResponseDataData 返回数据
type DecorationCouponGetV30ResponseDataData struct {
	// 卡券列表
	List     []*DecorationCouponGetV30ResponseDataDataListInner `json:"list,omitempty"`
	PageInfo *DecorationCouponGetV30ResponseDataDataPageInfo    `json:"page_info,omitempty"`
}
