/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceCityV2ResponseDataMetricsDict
type ReportAudienceCityV2ResponseDataMetricsDict struct {
	//
	Click *int64 `json:"click,omitempty"`
	//
	Convert *int64 `json:"convert,omitempty"`
	//
	Cost *float64 `json:"cost,omitempty"`
	//
	DownloadFinish *int64 `json:"download_finish,omitempty"`
	//
	Show *int64 `json:"show,omitempty"`
}