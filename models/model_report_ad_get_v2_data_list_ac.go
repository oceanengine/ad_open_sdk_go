/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2DataListAc
type ReportAdGetV2DataListAc string

// List of report_ad_get_v2_data_list_ac
const (
	Enum_3_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "3G"
	Enum_5_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "5G"
	Enum_4_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "4G"
	WIFI_ReportAdGetV2DataListAc     ReportAdGetV2DataListAc = "WIFI"
	UNKNOWN_ReportAdGetV2DataListAc  ReportAdGetV2DataListAc = "unknown"
	Enum_2_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "2G"
)

// Ptr returns reference to report_ad_get_v2_data_list_ac value
func (v ReportAdGetV2DataListAc) Ptr() *ReportAdGetV2DataListAc {
	return &v
}
