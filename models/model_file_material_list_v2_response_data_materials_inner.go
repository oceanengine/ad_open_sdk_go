/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialListV2ResponseDataMaterialsInner struct for FileMaterialListV2ResponseDataMaterialsInner
type FileMaterialListV2ResponseDataMaterialsInner struct {
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	MaterialProperties []*FileMaterialListV2DataMaterialsMaterialProperties `json:"material_properties,omitempty"`
	//
	MatetrialId *int64 `json:"matetrial_id,omitempty"`
}
