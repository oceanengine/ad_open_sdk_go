/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetExportBoostItemGroupResultV2ResponseData
type StarVasGetExportBoostItemGroupResultV2ResponseData struct {
	ExportStatus StarVasGetExportBoostItemGroupResultV2DataExportStatus `json:"export_status"`
	// 导出飞书文档地址，仅export_status为Success时有效
	Url *string `json:"url,omitempty"`
}
