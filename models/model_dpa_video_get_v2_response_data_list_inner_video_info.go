/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaVideoGetV2ResponseDataListInnerVideoInfo
type DpaVideoGetV2ResponseDataListInnerVideoInfo struct {
	//
	Height    *int64                                   `json:"height,omitempty"`
	ImageMode *DpaVideoGetV2DataListVideoInfoImageMode `json:"image_mode,omitempty"`
	Status    *DpaVideoGetV2DataListVideoInfoStatus    `json:"status,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}
