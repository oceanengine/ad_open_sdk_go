/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareCancelV30Response struct for EventManagerShareCancelV30Response
type EventManagerShareCancelV30Response struct {
	//
	Code *int64                                  `json:"code,omitempty"`
	Data *EventManagerShareCancelV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
