/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDetailV2ResponseDataForm
type ClueFormDetailV2ResponseDataForm struct {
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Elements []*ClueCouponDetailV2ResponseDataFormElementsInner `json:"elements,omitempty"`
	//
	EnableLayer *interface{}                                  `json:"enable_layer,omitempty"`
	ExtendInfo  *ClueCouponDetailV2ResponseDataFormExtendInfo `json:"extend_info,omitempty"`
	//
	FormType *interface{} `json:"form_type,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	IsDel *interface{} `json:"is_del,omitempty"`
	//
	LayerSubmitText *string `json:"layer_submit_text,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SubmitText *string `json:"submit_text,omitempty"`
	//
	UpdateTime *string `json:"update_time,omitempty"`
	//
	ValidateType *interface{} `json:"validate_type,omitempty"`
}
