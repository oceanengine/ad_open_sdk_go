/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaVideoGetV2DataListVideoInfoStatus
type DpaVideoGetV2DataListVideoInfoStatus string

// List of dpa_video_get_v2_data_list_video_info_status
const (
	AVAILABLE_DpaVideoGetV2DataListVideoInfoStatus   DpaVideoGetV2DataListVideoInfoStatus = "AVAILABLE"
	UNAVAILABLE_DpaVideoGetV2DataListVideoInfoStatus DpaVideoGetV2DataListVideoInfoStatus = "UNAVAILABLE"
)

// Ptr returns reference to dpa_video_get_v2_data_list_video_info_status value
func (v DpaVideoGetV2DataListVideoInfoStatus) Ptr() *DpaVideoGetV2DataListVideoInfoStatus {
	return &v
}
