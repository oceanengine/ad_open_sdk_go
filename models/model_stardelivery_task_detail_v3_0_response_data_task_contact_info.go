/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskDetailV30ResponseDataTaskContactInfo 联系人信息
type StardeliveryTaskDetailV30ResponseDataTaskContactInfo struct {
	// 联系人id
	ContactId *int64 `json:"contact_id,omitempty"`
	// 联系人姓名
	ContactName *string `json:"contact_name,omitempty"`
	// 联系人手机号
	ContactPhone *string `json:"contact_phone,omitempty"`
}
