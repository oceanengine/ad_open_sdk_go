/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialAttributesListV2ResponseData
type FileMaterialAttributesListV2ResponseData struct {
	// 素材及素材属性列表
	Materials []*FileMaterialAttributesListV2ResponseDataMaterialsInner `json:"materials"`
	Page      *FileMaterialAttributesListV2ResponseDataPage             `json:"page,omitempty"`
}
