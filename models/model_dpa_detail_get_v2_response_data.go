/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaDetailGetV2ResponseData
type DpaDetailGetV2ResponseData struct {
	//
	List     []*DpaDetailGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *DpaDetailGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
