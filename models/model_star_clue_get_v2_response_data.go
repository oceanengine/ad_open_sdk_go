/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarClueGetV2ResponseData
type StarClueGetV2ResponseData struct {
	// 线索数据
	List     []*StarClueGetV2ResponseDataListInner `json:"list"`
	PageInfo *StarClueGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
