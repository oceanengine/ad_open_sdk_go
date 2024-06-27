/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneCreateV2Request struct for ClueSmartphoneCreateV2Request
type ClueSmartphoneCreateV2Request struct {
	//
	AdvertiserId int64                              `json:"advertiser_id"`
	CallDisplay  *ClueSmartphoneCreateV2CallDisplay `json:"call_display,omitempty"`
	//
	Name string `json:"name"`
	//
	NoDisturbTimes []*ClueSmartphoneCreateV2RequestNoDisturbTimesInner `json:"no_disturb_times,omitempty"`
	//
	PhoneNumber  string                              `json:"phone_number"`
	ValidateType *ClueSmartphoneCreateV2ValidateType `json:"validate_type,omitempty"`
}
