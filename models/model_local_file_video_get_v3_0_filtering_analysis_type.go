/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalFileVideoGetV30FilteringAnalysisType
type LocalFileVideoGetV30FilteringAnalysisType string

// List of local_file_video_get_v3.0_filtering_analysis_type
const (
	FIRST_PUBLISH_LocalFileVideoGetV30FilteringAnalysisType                  LocalFileVideoGetV30FilteringAnalysisType = "FIRST_PUBLISH"
	FIRST_PUBLISH_AND_HIGH_QUALITY_LocalFileVideoGetV30FilteringAnalysisType LocalFileVideoGetV30FilteringAnalysisType = "FIRST_PUBLISH_AND_HIGH_QUALITY"
	HIGH_QUALITY_LocalFileVideoGetV30FilteringAnalysisType                   LocalFileVideoGetV30FilteringAnalysisType = "HIGH_QUALITY"
)

// Ptr returns reference to local_file_video_get_v3.0_filtering_analysis_type value
func (v LocalFileVideoGetV30FilteringAnalysisType) Ptr() *LocalFileVideoGetV30FilteringAnalysisType {
	return &v
}