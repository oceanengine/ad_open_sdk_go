/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoUpdateV2Request struct for FileVideoUpdateV2Request
type FileVideoUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Videos []*FileVideoUpdateV2RequestVideosInner `json:"videos,omitempty"`
}
