/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageDeleteV2Response struct for AudiencePackageDeleteV2Response
type AudiencePackageDeleteV2Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *AudiencePackageDeleteV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
