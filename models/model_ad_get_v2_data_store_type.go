/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataStoreType
type AdGetV2DataStoreType string

// List of ad_get_v2_data_store_type
const (
	STORE_THIRT_PARTY_AdGetV2DataStoreType AdGetV2DataStoreType = "STORE_THIRT_PARTY"
	STORE_DOUYIN_AdGetV2DataStoreType      AdGetV2DataStoreType = "STORE_DOUYIN"
	STORE_LANDING_AdGetV2DataStoreType     AdGetV2DataStoreType = "STORE_LANDING"
	STORE_NORMAL_AdGetV2DataStoreType      AdGetV2DataStoreType = "STORE_NORMAL"
)

// Ptr returns reference to ad_get_v2_data_store_type value
func (v AdGetV2DataStoreType) Ptr() *AdGetV2DataStoreType {
	return &v
}
