/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseDataListInnerDeliveryRange
type ProjectListV30ResponseDataListInnerDeliveryRange struct {
	InventoryCatalog *ProjectListV30DataListDeliveryRangeInventoryCatalog `json:"inventory_catalog,omitempty"`
	//
	InventoryType  []*ProjectListV30DataListDeliveryRangeInventoryType `json:"inventory_type,omitempty"`
	UnionVideoType *ProjectListV30DataListDeliveryRangeUnionVideoType  `json:"union_video_type,omitempty"`
}
