/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAudioGetV2Response struct for FileAudioGetV2Response
type FileAudioGetV2Response struct {
	//
	Code *int64                      `json:"code,omitempty"`
	Data *FileAudioGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
