/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMediaGetV2ResponseData
type FileMediaGetV2ResponseData struct {
	//
	List     []*FileMediaGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *FileMediaGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
