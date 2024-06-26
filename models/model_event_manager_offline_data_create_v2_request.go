/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerOfflineDataCreateV2Request struct for EventManagerOfflineDataCreateV2Request
type EventManagerOfflineDataCreateV2Request struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AssetsId *int64 `json:"assets_id,omitempty"`
	//
	EventType *string `json:"event_type,omitempty"`
	//
	OfflineData []*EventManagerOfflineDataCreateV2RequestOfflineDataInner `json:"offline_data,omitempty"`
}
