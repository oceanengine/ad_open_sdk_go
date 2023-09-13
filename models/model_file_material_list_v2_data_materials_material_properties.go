/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileMaterialListV2DataMaterialsMaterialProperties
type FileMaterialListV2DataMaterialsMaterialProperties string

// List of file_material_list_v2_data_materials_material_properties
const (
	AD_HIGH_QUALITY_MATERIAL_FileMaterialListV2DataMaterialsMaterialProperties  FileMaterialListV2DataMaterialsMaterialProperties = "AD_HIGH_QUALITY_MATERIAL"
	AIGC_FileMaterialListV2DataMaterialsMaterialProperties                      FileMaterialListV2DataMaterialsMaterialProperties = "AIGC"
	ECP_HIGH_QUALITY_MATERIAL_FileMaterialListV2DataMaterialsMaterialProperties FileMaterialListV2DataMaterialsMaterialProperties = "ECP_HIGH_QUALITY_MATERIAL"
	FIRST_PUBLISH_MATERIAL_FileMaterialListV2DataMaterialsMaterialProperties    FileMaterialListV2DataMaterialsMaterialProperties = "FIRST_PUBLISH_MATERIAL"
	INEFFICIENT_MATERIAL_FileMaterialListV2DataMaterialsMaterialProperties      FileMaterialListV2DataMaterialsMaterialProperties = "INEFFICIENT_MATERIAL"
	SIMILAR_MATERIAL_FileMaterialListV2DataMaterialsMaterialProperties          FileMaterialListV2DataMaterialsMaterialProperties = "SIMILAR_MATERIAL"
	SIMILAR_QUEUE_MATERIAL_FileMaterialListV2DataMaterialsMaterialProperties    FileMaterialListV2DataMaterialsMaterialProperties = "SIMILAR_QUEUE_MATERIAL"
)

// All allowed values of FileMaterialListV2DataMaterialsMaterialProperties enum
var AllowedFileMaterialListV2DataMaterialsMaterialPropertiesEnumValues = []FileMaterialListV2DataMaterialsMaterialProperties{
	"AD_HIGH_QUALITY_MATERIAL",
	"AIGC",
	"ECP_HIGH_QUALITY_MATERIAL",
	"FIRST_PUBLISH_MATERIAL",
	"INEFFICIENT_MATERIAL",
	"SIMILAR_MATERIAL",
	"SIMILAR_QUEUE_MATERIAL",
}

// NewFileMaterialListV2DataMaterialsMaterialPropertiesFromValue returns a pointer to a valid FileMaterialListV2DataMaterialsMaterialProperties
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialListV2DataMaterialsMaterialPropertiesFromValue(v string) (*FileMaterialListV2DataMaterialsMaterialProperties, error) {
	ev := FileMaterialListV2DataMaterialsMaterialProperties(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialListV2DataMaterialsMaterialProperties: valid values are %v", v, AllowedFileMaterialListV2DataMaterialsMaterialPropertiesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialListV2DataMaterialsMaterialProperties) IsValid() bool {
	for _, existing := range AllowedFileMaterialListV2DataMaterialsMaterialPropertiesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_list_v2_data_materials_material_properties value
func (v FileMaterialListV2DataMaterialsMaterialProperties) Ptr() *FileMaterialListV2DataMaterialsMaterialProperties {
	return &v
}
