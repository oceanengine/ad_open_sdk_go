/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerEventConfigsGetV2ResponseData
type EventManagerEventConfigsGetV2ResponseData struct {
	// 已创建事件列表
	EventConfigs []*EventManagerEventConfigsGetV2ResponseDataEventConfigsInner `json:"event_configs,omitempty"`
}
