/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FilePlayableListV30DataMaterialsStatus
type FilePlayableListV30DataMaterialsStatus string

// List of file_playable_list_v3.0_data_materials_status
const (
	OFFLINE_FilePlayableListV30DataMaterialsStatus FilePlayableListV30DataMaterialsStatus = "OFFLINE"
	ONLINE_FilePlayableListV30DataMaterialsStatus  FilePlayableListV30DataMaterialsStatus = "ONLINE"
)

// Ptr returns reference to file_playable_list_v3.0_data_materials_status value
func (v FilePlayableListV30DataMaterialsStatus) Ptr() *FilePlayableListV30DataMaterialsStatus {
	return &v
}
