/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdBudgetUpdateV10Request struct for QianchuanAdBudgetUpdateV10Request
type QianchuanAdBudgetUpdateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*QianchuanAdBudgetUpdateV10RequestDataInner `json:"data"`
}
