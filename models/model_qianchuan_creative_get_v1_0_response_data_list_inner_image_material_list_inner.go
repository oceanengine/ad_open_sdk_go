/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10ResponseDataListInnerImageMaterialListInner struct for QianchuanCreativeGetV10ResponseDataListInnerImageMaterialListInner
type QianchuanCreativeGetV10ResponseDataListInnerImageMaterialListInner struct {
	//
	ImageId   *string                                                    `json:"image_id,omitempty"`
	ImageMode *QianchuanCreativeGetV10DataListImageMaterialListImageMode `json:"image_mode,omitempty"`
	//
	IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"`
}
