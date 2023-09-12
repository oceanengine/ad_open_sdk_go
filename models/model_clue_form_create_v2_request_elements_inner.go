/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormCreateV2RequestElementsInner struct for ClueFormCreateV2RequestElementsInner
type ClueFormCreateV2RequestElementsInner struct {
	AllowEmpty *ClueFormCreateV2ElementsAllowEmpty `json:"allow_empty,omitempty"`
	//
	DefaultValue *int64                              `json:"default_value,omitempty"`
	ElementType  ClueFormCreateV2ElementsElementType `json:"element_type"`
	IsUnique     *ClueFormCreateV2ElementsIsUnique   `json:"is_unique,omitempty"`
	//
	Label string                         `json:"label"`
	Layer *ClueFormCreateV2ElementsLayer `json:"layer,omitempty"`
	//
	Sequence *int64 `json:"sequence,omitempty"`
	//
	Value *string `json:"value,omitempty"`
}
