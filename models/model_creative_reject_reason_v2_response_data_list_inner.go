/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeRejectReasonV2ResponseDataListInner struct for CreativeRejectReasonV2ResponseDataListInner
type CreativeRejectReasonV2ResponseDataListInner struct {
	//
	CreativeId *int64 `json:"creative_id,omitempty"`
	//
	MaterialReject []*CreativeRejectReasonV2ResponseDataListInnerMaterialRejectInner `json:"material_reject,omitempty"`
	//
	RejectData []*CreativeRejectReasonV2ResponseDataListInnerRejectDataInner `json:"reject_data,omitempty"`
}
