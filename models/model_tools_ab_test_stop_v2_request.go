/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestStopV2Request struct for ToolsAbTestStopV2Request
type ToolsAbTestStopV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	TestIds []int64 `json:"test_ids"`
}
