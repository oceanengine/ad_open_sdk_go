/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserFundTransactionGetV2ResponseData
type AdvertiserFundTransactionGetV2ResponseData struct {
	//
	List     []*AdvertiserFundTransactionGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *AdvertiserFundDailyStatV2ResponseDataPageInfo         `json:"page_info,omitempty"`
}
