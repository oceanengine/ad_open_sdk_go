/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalFileVideoGetV30DataVideoListMaterialProperties
type LocalFileVideoGetV30DataVideoListMaterialProperties string

// List of local_file_video_get_v3.0_data_video_list_material_properties
const (
	COPY_LocalFileVideoGetV30DataVideoListMaterialProperties          LocalFileVideoGetV30DataVideoListMaterialProperties = "COPY"
	FIRST_PUBLISH_LocalFileVideoGetV30DataVideoListMaterialProperties LocalFileVideoGetV30DataVideoListMaterialProperties = "FIRST_PUBLISH"
	HIGH_QUALITY_LocalFileVideoGetV30DataVideoListMaterialProperties  LocalFileVideoGetV30DataVideoListMaterialProperties = "HIGH_QUALITY"
	LOW_EFFICIENY_LocalFileVideoGetV30DataVideoListMaterialProperties LocalFileVideoGetV30DataVideoListMaterialProperties = "LOW_EFFICIENY"
	LOW_QUALITY_LocalFileVideoGetV30DataVideoListMaterialProperties   LocalFileVideoGetV30DataVideoListMaterialProperties = "LOW_QUALITY"
	SIMILAR_LocalFileVideoGetV30DataVideoListMaterialProperties       LocalFileVideoGetV30DataVideoListMaterialProperties = "SIMILAR"
)

// Ptr returns reference to local_file_video_get_v3.0_data_video_list_material_properties value
func (v LocalFileVideoGetV30DataVideoListMaterialProperties) Ptr() *LocalFileVideoGetV30DataVideoListMaterialProperties {
	return &v
}
