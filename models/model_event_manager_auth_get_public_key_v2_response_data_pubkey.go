/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAuthGetPublicKeyV2ResponseDataPubkey
type EventManagerAuthGetPublicKeyV2ResponseDataPubkey struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Credential *string `json:"credential,omitempty"`
	//
	KeyId *int64 `json:"key_id,omitempty"`
	//
	Pubkey *string `json:"pubkey,omitempty"`
}
