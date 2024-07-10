/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SubscribeAccountsAddV30Request struct for SubscribeAccountsAddV30Request
type SubscribeAccountsAddV30Request struct {
	//
	AdvertiserIds []int64 `json:"advertiser_ids"`
	//
	AppId int64 `json:"app_id"`
	//
	CoreUserId int64 `json:"core_user_id"`
	//
	Events []string `json:"events,omitempty"`
	//
	SubscribeTaskId int64 `json:"subscribe_task_id"`
}
