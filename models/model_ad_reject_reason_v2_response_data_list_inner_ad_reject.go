/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdRejectReasonV2ResponseDataListInnerAdReject
type AdRejectReasonV2ResponseDataListInnerAdReject struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	RejectData []*AdRejectReasonV2ResponseDataListInnerAdRejectRejectDataInner `json:"reject_data,omitempty"`
}
