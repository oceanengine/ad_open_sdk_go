/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30Request struct for NativeAnchorCreateV30Request
type NativeAnchorCreateV30Request struct {
	//
	AdvertiserId *int64                                 `json:"advertiser_id,omitempty"`
	AnchorInfo   NativeAnchorCreateV30RequestAnchorInfo `json:"anchor_info"`
}
