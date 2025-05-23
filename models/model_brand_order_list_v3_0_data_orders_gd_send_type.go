/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30DataOrdersGdSendType
type BrandOrderListV30DataOrdersGdSendType string

// List of brand_order_list_v3.0_data_orders_gd_send_type
const (
	CPV_BrandOrderListV30DataOrdersGdSendType                BrandOrderListV30DataOrdersGdSendType = "CPV"
	CPV_EFFECTIVE_BrandOrderListV30DataOrdersGdSendType      BrandOrderListV30DataOrdersGdSendType = "CPV_EFFECTIVE"
	CTR_BrandOrderListV30DataOrdersGdSendType                BrandOrderListV30DataOrdersGdSendType = "CTR"
	CVR_BrandOrderListV30DataOrdersGdSendType                BrandOrderListV30DataOrdersGdSendType = "CVR"
	FANS_INCR_BrandOrderListV30DataOrdersGdSendType          BrandOrderListV30DataOrdersGdSendType = "FANS_INCR"
	FORM_BrandOrderListV30DataOrdersGdSendType               BrandOrderListV30DataOrdersGdSendType = "FORM"
	HOISTING_BrandOrderListV30DataOrdersGdSendType           BrandOrderListV30DataOrdersGdSendType = "HOISTING"
	INTERACTIVE_BrandOrderListV30DataOrdersGdSendType        BrandOrderListV30DataOrdersGdSendType = "INTERACTIVE"
	LIVE_INTERACTIVE_BrandOrderListV30DataOrdersGdSendType   BrandOrderListV30DataOrdersGdSendType = "LIVE_INTERACTIVE"
	LIVE_ROOM_BrandOrderListV30DataOrdersGdSendType          BrandOrderListV30DataOrdersGdSendType = "LIVE_ROOM"
	LIVE_WATICH_BrandOrderListV30DataOrdersGdSendType        BrandOrderListV30DataOrdersGdSendType = "LIVE_WATICH"
	PLANT_GRASS_BrandOrderListV30DataOrdersGdSendType        BrandOrderListV30DataOrdersGdSendType = "PLANT_GRASS"
	REACH_BrandOrderListV30DataOrdersGdSendType              BrandOrderListV30DataOrdersGdSendType = "REACH"
	SEQUENCE_BrandOrderListV30DataOrdersGdSendType           BrandOrderListV30DataOrdersGdSendType = "SEQUENCE"
	STRIDE_PLANT_GRASS_BrandOrderListV30DataOrdersGdSendType BrandOrderListV30DataOrdersGdSendType = "STRIDE_PLANT_GRASS"
)

// Ptr returns reference to brand_order_list_v3.0_data_orders_gd_send_type value
func (v BrandOrderListV30DataOrdersGdSendType) Ptr() *BrandOrderListV30DataOrdersGdSendType {
	return &v
}
