/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetAuthorListV2ResponseData
type StarMcnGetAuthorListV2ResponseData struct {
	//
	Data     []*StarMcnGetAuthorListV2ResponseDataDataInner `json:"data,omitempty"`
	PageInfo *StarMcnGetAuthorListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}