/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2DataListAc
type ReportCreativeGetV2DataListAc string

// List of report_creative_get_v2_data_list_ac
const (
	WIFI_ReportCreativeGetV2DataListAc     ReportCreativeGetV2DataListAc = "WIFI"
	Enum_4_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "4G"
	Enum_3_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "3G"
	Enum_2_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "2G"
	Enum_5_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "5G"
	UNKNOWN_ReportCreativeGetV2DataListAc  ReportCreativeGetV2DataListAc = "unknown"
)

// Ptr returns reference to report_creative_get_v2_data_list_ac value
func (v ReportCreativeGetV2DataListAc) Ptr() *ReportCreativeGetV2DataListAc {
	return &v
}
