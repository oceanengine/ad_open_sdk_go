/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AssetsCreativeComponentCreateV2RequestComponentInfo
type AssetsCreativeComponentCreateV2RequestComponentInfo struct {
	//
	ComponentData map[string]interface{} `json:"component_data"`
	//
	ComponentName string                                                    `json:"component_name"`
	ComponentType AssetsCreativeComponentCreateV2ComponentInfoComponentType `json:"component_type"`
}
