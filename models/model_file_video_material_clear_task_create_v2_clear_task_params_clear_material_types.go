/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes
type FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes string

// List of file_video_material_clear_task_create_v2_clear_task_params_clear_material_types
const (
	INEFFICIENT_MATERIAL_FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes   FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes = "INEFFICIENT_MATERIAL"
	SIMILAR_MATERIAL_FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes       FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes = "SIMILAR_MATERIAL"
	SIMILAR_QUEUE_MATERIAL_FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes = "SIMILAR_QUEUE_MATERIAL"
)

// Ptr returns reference to file_video_material_clear_task_create_v2_clear_task_params_clear_material_types value
func (v FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes) Ptr() *FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes {
	return &v
}
