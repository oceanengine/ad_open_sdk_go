/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserBudgetGetV2ResponseDataListInner struct for AdvertiserBudgetGetV2ResponseDataListInner
type AdvertiserBudgetGetV2ResponseDataListInner struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget     float64                                 `json:"budget"`
	BudgetMode AdvertiserBudgetGetV2DataListBudgetMode `json:"budget_mode"`
}
