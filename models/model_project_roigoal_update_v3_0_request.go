/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectRoigoalUpdateV30Request struct for ProjectRoigoalUpdateV30Request
type ProjectRoigoalUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*ProjectRoigoalUpdateV30RequestDataInner `json:"data"`
}