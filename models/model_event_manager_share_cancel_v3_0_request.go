/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareCancelV30Request struct for EventManagerShareCancelV30Request
type EventManagerShareCancelV30Request struct {
	//
	AccountInfos []*EventManagerShareCancelV30RequestAccountInfosInner `json:"account_infos,omitempty"`
	//
	AllAccountType []*EventManagerShareCancelV30AllAccountType `json:"all_account_type,omitempty"`
	//
	AssetId int64 `json:"asset_id"`
	//
	OrganizationId int64                               `json:"organization_id"`
	ShareMode      EventManagerShareCancelV30ShareMode `json:"share_mode"`
}
