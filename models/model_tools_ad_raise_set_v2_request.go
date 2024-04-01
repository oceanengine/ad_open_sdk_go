/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdRaiseSetV2Request struct for ToolsAdRaiseSetV2Request
type ToolsAdRaiseSetV2Request struct {
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	BudgetValue *int64 `json:"budget_value,omitempty"`
	//
	ModifyValue *int64                   `json:"modify_value,omitempty"`
	OptType     ToolsAdRaiseSetV2OptType `json:"opt_type"`
}
