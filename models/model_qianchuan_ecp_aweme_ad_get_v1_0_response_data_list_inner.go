/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanEcpAwemeAdGetV10ResponseDataListInner struct for QianchuanEcpAwemeAdGetV10ResponseDataListInner
type QianchuanEcpAwemeAdGetV10ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdvertiserId  *int64                                          `json:"advertiser_id,omitempty"`
	MarketingGoal *QianchuanEcpAwemeAdGetV10DataListMarketingGoal `json:"marketing_goal,omitempty"`
	//
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	//
	OrderId    *int64                                       `json:"order_id,omitempty"`
	OrderStaus *QianchuanEcpAwemeAdGetV10DataListOrderStaus `json:"order_staus,omitempty"`
}
