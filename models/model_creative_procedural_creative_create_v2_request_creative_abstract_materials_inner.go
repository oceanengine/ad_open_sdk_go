/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInner struct for CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInner
type CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInner struct {
	//
	MaterialId         *int64                                                                                     `json:"material_id,omitempty"`
	StructAbstractInfo *CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInnerStructAbstractInfo `json:"struct_abstract_info,omitempty"`
	TextAbstractInfo   *CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfo   `json:"text_abstract_info,omitempty"`
}
