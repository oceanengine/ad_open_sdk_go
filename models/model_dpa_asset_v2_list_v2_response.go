/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetV2ListV2Response struct for DpaAssetV2ListV2Response
type DpaAssetV2ListV2Response struct {
	//
	Code *int64                        `json:"code,omitempty"`
	Data *DpaAssetV2ListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
