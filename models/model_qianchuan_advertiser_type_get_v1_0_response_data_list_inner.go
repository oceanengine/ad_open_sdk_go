/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdvertiserTypeGetV10ResponseDataListInner struct for QianchuanAdvertiserTypeGetV10ResponseDataListInner
type QianchuanAdvertiserTypeGetV10ResponseDataListInner struct {
	//
	AdvertiserId *int64                                        `json:"advertiser_id,omitempty"`
	EcpType      *QianchuanAdvertiserTypeGetV10DataListEcpType `json:"ecp_type,omitempty"`
}
