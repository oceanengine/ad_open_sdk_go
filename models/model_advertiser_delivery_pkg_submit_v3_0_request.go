/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgSubmitV30Request struct for AdvertiserDeliveryPkgSubmitV30Request
type AdvertiserDeliveryPkgSubmitV30Request struct {
	// 广告主账户ID
	AdvertiserId int64                                            `json:"advertiser_id"`
	DeliveryPkg  AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkg `json:"delivery_pkg"`
}
