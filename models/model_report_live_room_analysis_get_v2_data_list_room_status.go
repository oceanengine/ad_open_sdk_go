/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportLiveRoomAnalysisGetV2DataListRoomStatus
type ReportLiveRoomAnalysisGetV2DataListRoomStatus string

// List of report_live_room_analysis_get_v2_data_list_room_status
const (
	PAUSE_ReportLiveRoomAnalysisGetV2DataListRoomStatus     ReportLiveRoomAnalysisGetV2DataListRoomStatus = "PAUSE"
	PREPARING_ReportLiveRoomAnalysisGetV2DataListRoomStatus ReportLiveRoomAnalysisGetV2DataListRoomStatus = "PREPARING"
	LIVING_ReportLiveRoomAnalysisGetV2DataListRoomStatus    ReportLiveRoomAnalysisGetV2DataListRoomStatus = "LIVING"
	END_ReportLiveRoomAnalysisGetV2DataListRoomStatus       ReportLiveRoomAnalysisGetV2DataListRoomStatus = "END"
)

// Ptr returns reference to report_live_room_analysis_get_v2_data_list_room_status value
func (v ReportLiveRoomAnalysisGetV2DataListRoomStatus) Ptr() *ReportLiveRoomAnalysisGetV2DataListRoomStatus {
	return &v
}
