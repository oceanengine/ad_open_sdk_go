/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCostProtectStatusGetV30Response struct for PromotionCostProtectStatusGetV30Response
type PromotionCostProtectStatusGetV30Response struct {
	//
	Code *int64                                        `json:"code,omitempty"`
	Data *PromotionCostProtectStatusGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
