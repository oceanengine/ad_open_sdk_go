/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoMaterialClearTaskResultGetV2ResponseData
type FileVideoMaterialClearTaskResultGetV2ResponseData struct {
	//
	ClearResult []*FileVideoMaterialClearTaskResultGetV2ResponseDataClearResultInner `json:"clear_result,omitempty"`
	PageInfo    *FileVideoMaterialClearTaskResultGetV2ResponseDataPageInfo           `json:"page_info,omitempty"`
}
