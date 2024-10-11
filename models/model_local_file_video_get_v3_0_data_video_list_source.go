/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalFileVideoGetV30DataVideoListSource
type LocalFileVideoGetV30DataVideoListSource string

// List of local_file_video_get_v3.0_data_video_list_source
const (
	BP_PLATFORM_LocalFileVideoGetV30DataVideoListSource      LocalFileVideoGetV30DataVideoListSource = "BP_PLATFORM"
	CREATIVE_AIGC_LocalFileVideoGetV30DataVideoListSource    LocalFileVideoGetV30DataVideoListSource = "CREATIVE_AIGC"
	LOCAL_ADS_UPLOAD_LocalFileVideoGetV30DataVideoListSource LocalFileVideoGetV30DataVideoListSource = "LOCAL_ADS_UPLOAD"
	MAPI_LocalFileVideoGetV30DataVideoListSource             LocalFileVideoGetV30DataVideoListSource = "MAPI"
	STAR_LocalFileVideoGetV30DataVideoListSource             LocalFileVideoGetV30DataVideoListSource = "STAR"
)

// Ptr returns reference to local_file_video_get_v3.0_data_video_list_source value
func (v LocalFileVideoGetV30DataVideoListSource) Ptr() *LocalFileVideoGetV30DataVideoListSource {
	return &v
}