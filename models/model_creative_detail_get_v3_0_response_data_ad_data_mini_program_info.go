/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataAdDataMiniProgramInfo
type CreativeDetailGetV30ResponseDataAdDataMiniProgramInfo struct {
	//
	AppId *string `json:"app_id,omitempty"`
	//
	Params *string `json:"params,omitempty"`
	//
	StartPath *string                                            `json:"start_path,omitempty"`
	Type      *CreativeDetailGetV30DataAdDataMiniProgramInfoType `json:"type,omitempty"`
}
