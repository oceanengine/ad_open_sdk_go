/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdLearingStatusGetV10ResponseDataListInner struct for QianchuanAdLearingStatusGetV10ResponseDataListInner
type QianchuanAdLearingStatusGetV10ResponseDataListInner struct {
	//
	AdId   *int64                                        `json:"ad_id,omitempty"`
	Status *QianchuanAdLearingStatusGetV10DataListStatus `json:"status,omitempty"`
}