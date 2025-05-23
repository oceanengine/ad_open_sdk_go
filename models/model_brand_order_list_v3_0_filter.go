/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30Filter
type BrandOrderListV30Filter struct {
	// 投放产品
	AdForm    []*BrandOrderListV30FilterAdForm  `json:"ad_form,omitempty"`
	AppOrigin *BrandOrderListV30FilterAppOrigin `json:"app_origin,omitempty"`
	//
	AuditStatus []*BrandOrderListV30FilterAuditStatus `json:"audit_status,omitempty"`
	Classify    *BrandOrderListV30FilterClassify      `json:"classify,omitempty"`
	// 创建截止时间，格式：2024-01-01
	CreateEndTime *string `json:"create_end_time,omitempty"`
	// 创建起始时间，格式：2024-01-01
	CreateStartTime *string `json:"create_start_time,omitempty"`
	// 投放截止时间，格式：2024-01-01
	EndTime    *string                            `json:"end_time,omitempty"`
	GdSendType *BrandOrderListV30FilterGdSendType `json:"gd_send_type,omitempty"`
	// 预订单ID
	OrderIds []int64 `json:"order_ids,omitempty"`
	// 预订单名称
	OrderName *string `json:"order_name,omitempty"`
	// 预订单状态
	OrderStatus []*BrandOrderListV30FilterOrderStatus `json:"order_status,omitempty"`
	ProType     *BrandOrderListV30FilterProType       `json:"pro_type,omitempty"`
	// 投放起始时间，格式：2024-01-01
	StartTime *string `json:"start_time,omitempty"`
}
