/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAssetsCreateV2RequestThirdPartAsset 三方落地页资产信息
type EventManagerAssetsCreateV2RequestThirdPartAsset struct {
	// 落地页描述信息，长度限制为`150`，一个字符长度为1
	Description *string `json:"description,omitempty"`
	// 落地页名称，长度限制为`25`，一个字符长度为1
	Name *string `json:"name,omitempty"`
}
