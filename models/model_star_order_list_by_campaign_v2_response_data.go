/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderListByCampaignV2ResponseData
type StarOrderListByCampaignV2ResponseData struct {
	// 任务列表
	OrderList []*StarOrderListByCampaignV2ResponseDataOrderListInner `json:"order_list,omitempty"`
	PageInfo  *StarOrderListByCampaignV2ResponseDataPageInfo         `json:"page_info,omitempty"`
}
