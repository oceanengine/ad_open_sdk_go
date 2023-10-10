/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetV30ResponseDataListInner struct for NativeAnchorGetV30ResponseDataListInner
type NativeAnchorGetV30ResponseDataListInner struct {
	//
	AnchorId *string `json:"anchor_id,omitempty"`
	//
	AnchorName *string                               `json:"anchor_name,omitempty"`
	AnchorType *NativeAnchorGetV30DataListAnchorType `json:"anchor_type,omitempty"`
	//
	AndroidPackageName *string `json:"android_package_name,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	IosPackageName *string                           `json:"ios_package_name,omitempty"`
	Source         *NativeAnchorGetV30DataListSource `json:"source,omitempty"`
	Status         *NativeAnchorGetV30DataListStatus `json:"status,omitempty"`
}
