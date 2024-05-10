/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormCreateV2Request struct for ClueFormCreateV2Request
type ClueFormCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Elements    []*ClueFormCreateV2RequestElementsInner `json:"elements"`
	EnableLayer *ClueFormCreateV2EnableLayer            `json:"enable_layer,omitempty"`
	ExtendInfo  *ClueFormCreateV2RequestExtendInfo      `json:"extend_info,omitempty"`
	//
	LayerSubmitText *string `json:"layer_submit_text,omitempty"`
	//
	Name string `json:"name"`
	//
	SubmitText   string                       `json:"submit_text"`
	ValidateType ClueFormCreateV2ValidateType `json:"validate_type"`
}
