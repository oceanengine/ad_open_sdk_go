/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdUpdateBudgetV2Request struct for AdUpdateBudgetV2Request
type AdUpdateBudgetV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*AdUpdateBudgetV2RequestDataInner `json:"data"`
}
