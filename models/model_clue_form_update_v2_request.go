/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormUpdateV2Request struct for ClueFormUpdateV2Request
type ClueFormUpdateV2Request struct {
	//
	AdvertiserId int64                              `json:"advertiser_id"`
	ExtendInfo   *ClueFormUpdateV2RequestExtendInfo `json:"extend_info,omitempty"`
	//
	InstanceId int64 `json:"instance_id"`
	//
	LayerSubmitText *string `json:"layer_submit_text,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SubmitText *string `json:"submit_text,omitempty"`
}
