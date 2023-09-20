/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareV30Request struct for EventManagerShareV30Request
type EventManagerShareV30Request struct {
	//
	AccountInfos []*EventManagerShareV30RequestAccountInfosInner `json:"account_infos,omitempty"`
	//
	AllAccountType []*EventManagerShareV30AllAccountType `json:"all_account_type,omitempty"`
	//
	AssetId int64 `json:"asset_id"`
	//
	OrganizationId int64                         `json:"organization_id"`
	ShareMode      EventManagerShareV30ShareMode `json:"share_mode"`
}
