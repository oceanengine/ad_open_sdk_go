/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInner struct for CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInner
type CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInner struct {
	//
	MaterialId         *int64                                                                                     `json:"material_id,omitempty"`
	StructAbstractInfo *CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInnerStructAbstractInfo `json:"struct_abstract_info,omitempty"`
	TextAbstractInfo   *CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfo   `json:"text_abstract_info,omitempty"`
}
