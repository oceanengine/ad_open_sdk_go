/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdvertiserGetV2DataListRealRecallMatchType
type ReportAdvertiserGetV2DataListRealRecallMatchType string

// List of report_advertiser_get_v2_data_list_real_recall_match_type
const (
	EXTENSIVE_ReportAdvertiserGetV2DataListRealRecallMatchType ReportAdvertiserGetV2DataListRealRecallMatchType = "EXTENSIVE"
	PHRASE_ReportAdvertiserGetV2DataListRealRecallMatchType    ReportAdvertiserGetV2DataListRealRecallMatchType = "PHRASE"
	PRECISION_ReportAdvertiserGetV2DataListRealRecallMatchType ReportAdvertiserGetV2DataListRealRecallMatchType = "PRECISION"
)

// Ptr returns reference to report_advertiser_get_v2_data_list_real_recall_match_type value
func (v ReportAdvertiserGetV2DataListRealRecallMatchType) Ptr() *ReportAdvertiserGetV2DataListRealRecallMatchType {
	return &v
}
