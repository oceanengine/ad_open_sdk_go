/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceAction
type AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceAction struct {
	//
	ActionCategories []int64 `json:"action_categories,omitempty"`
	//
	ActionDays *int64 `json:"action_days,omitempty"`
	//
	ActionScene []*AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
}
