/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAvailableEventsGetV2ResponseData
type EventManagerAvailableEventsGetV2ResponseData struct {
	// 可创建事件列表
	EventConfigs []*EventManagerAvailableEventsGetV2ResponseDataEventConfigsInner `json:"event_configs,omitempty"`
}
