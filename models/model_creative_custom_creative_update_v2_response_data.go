/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeUpdateV2ResponseData
type CreativeCustomCreativeUpdateV2ResponseData struct {
	//
	CreativeIds []int64 `json:"creative_ids,omitempty"`
	//
	Errors []*CreativeCustomCreativeUpdateV2ResponseDataErrorsInner `json:"errors,omitempty"`
}