/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentQueryLinkV2ResponseData
type StarComponentQueryLinkV2ResponseData struct {
	//
	Components []*StarComponentQueryLinkV2ResponseDataComponentsInner `json:"components,omitempty"`
	PageInfo   *StarComponentQueryLinkV2ResponseDataPageInfo          `json:"page_info,omitempty"`
}
