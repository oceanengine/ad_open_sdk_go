/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AssetsCreativeComponentGetV2ResponseDataListInner struct for AssetsCreativeComponentGetV2ResponseDataListInner
type AssetsCreativeComponentGetV2ResponseDataListInner struct {
	//
	ComponentData map[string]interface{} `json:"component_data,omitempty"`
	//
	ComponentId *int64 `json:"component_id,omitempty"`
	//
	ComponentName *string `json:"component_name,omitempty"`
	//
	ComponentType *string `json:"component_type,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Status *string `json:"status,omitempty"`
}
