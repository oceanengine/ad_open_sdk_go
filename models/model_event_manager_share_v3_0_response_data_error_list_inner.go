/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareV30ResponseDataErrorListInner struct for EventManagerShareV30ResponseDataErrorListInner
type EventManagerShareV30ResponseDataErrorListInner struct {
	AccountInfo    *EventManagerShareV30ResponseDataErrorListInnerAccountInfo `json:"account_info,omitempty"`
	AllAccountType *EventManagerShareV30DataErrorListAllAccountType           `json:"all_account_type,omitempty"`
	//
	ErrorMessage *string                                     `json:"error_message,omitempty"`
	ShareMode    *EventManagerShareV30DataErrorListShareMode `json:"share_mode,omitempty"`
}
