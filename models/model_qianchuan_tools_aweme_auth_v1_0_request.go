/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsAwemeAuthV10Request struct for QianchuanToolsAwemeAuthV10Request
type QianchuanToolsAwemeAuthV10Request struct {
	//
	AdvertiserId int64                              `json:"advertiser_id"`
	AuthType     QianchuanToolsAwemeAuthV10AuthType `json:"auth_type"`
	//
	AwemeId string `json:"aweme_id"`
	//
	Code string `json:"code"`
	//
	EndTime string `json:"end_time"`
}
