/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalCxtReportAudienceGetV30ResponseDataListInner struct for LocalCxtReportAudienceGetV30ResponseDataListInner
type LocalCxtReportAudienceGetV30ResponseDataListInner struct {
	//
	Age *string `json:"age,omitempty"`
	//
	City *string `json:"city,omitempty"`
	//
	Gender  *string                                                   `json:"gender,omitempty"`
	Metrics *LocalCxtReportAudienceGetV30ResponseDataListInnerMetrics `json:"metrics,omitempty"`
}
