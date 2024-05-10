/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2ResponseDataFormElementsInner struct for ClueCouponDetailV2ResponseDataFormElementsInner
type ClueCouponDetailV2ResponseDataFormElementsInner struct {
	AllowEmpty *ClueCouponDetailV2DataFormElementsAllowEmpty `json:"allow_empty,omitempty"`
	//
	DefaultValue *int64 `json:"default_value,omitempty"`
	//
	ElementId   *int64                                         `json:"element_id,omitempty"`
	ElementType *ClueCouponDetailV2DataFormElementsElementType `json:"element_type,omitempty"`
	IsUnique    *ClueCouponDetailV2DataFormElementsIsUnique    `json:"is_unique,omitempty"`
	//
	Label *string `json:"label,omitempty"`
	//
	Layer *int64 `json:"layer,omitempty"`
	//
	Sequence *int64 `json:"sequence,omitempty"`
}
