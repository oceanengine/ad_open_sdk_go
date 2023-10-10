/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormListV2ResponseDataListInner struct for ClueFormListV2ResponseDataListInner
type ClueFormListV2ResponseDataListInner struct {
	//
	CreateTime  *string                            `json:"create_time,omitempty"`
	EnableLayer *ClueFormListV2DataListEnableLayer `json:"enable_layer,omitempty"`
	FormType    *ClueFormListV2DataListFormType    `json:"form_type,omitempty"`
	//
	InstanceId *int64                       `json:"instance_id,omitempty"`
	IsDel      *ClueFormListV2DataListIsDel `json:"is_del,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	UpdateTime   *string                             `json:"update_time,omitempty"`
	ValidateType *ClueFormListV2DataListValidateType `json:"validate_type,omitempty"`
}
