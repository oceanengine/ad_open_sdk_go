/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectStatusUpdateV30Request struct for ProjectStatusUpdateV30Request
type ProjectStatusUpdateV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*ProjectStatusUpdateV30RequestDataInner `json:"data"`
}
