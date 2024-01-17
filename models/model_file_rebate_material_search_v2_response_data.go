/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialSearchV2ResponseData
type FileRebateMaterialSearchV2ResponseData struct {
	CursorInfo *FileRebateMaterialSearchV2ResponseDataCursorInfo `json:"cursor_info,omitempty"`
	//
	Materials []*FileRebateMaterialSearchV2ResponseDataMaterialsInner `json:"materials,omitempty"`
	PageInfo  *FileRebateMaterialSearchV2ResponseDataPageInfo         `json:"page_info,omitempty"`
}
