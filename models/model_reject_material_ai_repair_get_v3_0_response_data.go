/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// RejectMaterialAiRepairGetV30ResponseData
type RejectMaterialAiRepairGetV30ResponseData struct {
	PageInfo *RejectMaterialAiRepairGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	RepairMaterialInfo []*RejectMaterialAiRepairGetV30ResponseDataRepairMaterialInfoInner `json:"repair_material_info,omitempty"`
}
