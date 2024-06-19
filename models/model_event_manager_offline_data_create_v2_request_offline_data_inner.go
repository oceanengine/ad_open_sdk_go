/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerOfflineDataCreateV2RequestOfflineDataInner struct for EventManagerOfflineDataCreateV2RequestOfflineDataInner
type EventManagerOfflineDataCreateV2RequestOfflineDataInner struct {
	//
	Callback *string `json:"callback,omitempty"`
	//
	Channel *string `json:"channel,omitempty"`
	//
	Customer *string `json:"customer,omitempty"`
	//
	EventTime string `json:"event_time"`
	//
	Idfa *string `json:"idfa,omitempty"`
	//
	IdfaMd5 *string `json:"idfa_md5,omitempty"`
	//
	ImeiMd5 *string `json:"imei_md5,omitempty"`
	//
	Label *int64 `json:"label,omitempty"`
	//
	MacMd5 *string `json:"mac_md5,omitempty"`
	//
	Oaid *string `json:"oaid,omitempty"`
	//
	OaidMd5 *string `json:"oaid_md5,omitempty"`
	//
	PayAmount *float64 `json:"pay_amount,omitempty"`
	//
	PhoneNumberMd5 *string `json:"phone_number_md5,omitempty"`
	//
	PhoneNumberSha256 *string `json:"phone_number_sha256,omitempty"`
	//
	ProductLine *string `json:"product_line,omitempty"`
	//
	Uid *string `json:"uid,omitempty"`
}
