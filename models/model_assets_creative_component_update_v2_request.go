/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AssetsCreativeComponentUpdateV2Request struct for AssetsCreativeComponentUpdateV2Request
type AssetsCreativeComponentUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ComponentId   int64                                               `json:"component_id"`
	ComponentInfo AssetsCreativeComponentCreateV2RequestComponentInfo `json:"component_info"`
}