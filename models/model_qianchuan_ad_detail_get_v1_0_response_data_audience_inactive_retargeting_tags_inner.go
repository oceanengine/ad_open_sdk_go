/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataAudienceInactiveRetargetingTagsInner struct for QianchuanAdDetailGetV10ResponseDataAudienceInactiveRetargetingTagsInner
type QianchuanAdDetailGetV10ResponseDataAudienceInactiveRetargetingTagsInner struct {
	InactiveType *QianchuanAdDetailGetV10DataAudienceInactiveRetargetingTagsInactiveType `json:"inactive_type,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	RetargetingTag *int64 `json:"retargeting_tag,omitempty"`
}