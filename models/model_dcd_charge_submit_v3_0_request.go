/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DcdChargeSubmitV30Request struct for DcdChargeSubmitV30Request
type DcdChargeSubmitV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 充值金额，单位为分（即实际金额*100）
	ChargeAmount int64 `json:"charge_amount"`
	// 充值请求唯一编号
	ChargeRequestId string                           `json:"charge_request_id"`
	PayInfo         DcdChargeSubmitV30RequestPayInfo `json:"pay_info"`
	Platform        DcdChargeSubmitV30Platform       `json:"platform"`
}
