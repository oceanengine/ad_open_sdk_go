/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataInventoryCatalog
type AdGetV2DataInventoryCatalog string

// List of ad_get_v2_data_inventory_catalog
const (
	UNIVERSAL_SMART_AdGetV2DataInventoryCatalog AdGetV2DataInventoryCatalog = "UNIVERSAL_SMART"
	MANUAL_AdGetV2DataInventoryCatalog          AdGetV2DataInventoryCatalog = "MANUAL"
	UNIVERSAL_AdGetV2DataInventoryCatalog       AdGetV2DataInventoryCatalog = "UNIVERSAL"
	SCENE_AdGetV2DataInventoryCatalog           AdGetV2DataInventoryCatalog = "SCENE"
	SMART_AdGetV2DataInventoryCatalog           AdGetV2DataInventoryCatalog = "SMART"
)

// Ptr returns reference to ad_get_v2_data_inventory_catalog value
func (v AdGetV2DataInventoryCatalog) Ptr() *AdGetV2DataInventoryCatalog {
	return &v
}
