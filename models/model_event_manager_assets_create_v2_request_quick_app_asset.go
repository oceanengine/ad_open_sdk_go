/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAssetsCreateV2RequestQuickAppAsset 快应用资产信息
type EventManagerAssetsCreateV2RequestQuickAppAsset struct {
	// 快应用名称，长度限制为20，一个字符长度为1
	Name string `json:"name"`
	// 快应用包名，长度限制为140
	PackageName string `json:"package_name"`
}
