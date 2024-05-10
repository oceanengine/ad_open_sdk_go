/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeUpdateV2RequestCreativeListInner struct for CreativeCustomCreativeUpdateV2RequestCreativeListInner
type CreativeCustomCreativeUpdateV2RequestCreativeListInner struct {
	//
	AbstractMaterials []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerAbstractMaterialsInner `json:"abstract_materials,omitempty"`
	//
	ComponentMaterials []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerComponentMaterialsInner `json:"component_materials,omitempty"`
	//
	CreativeId         *int64                                                                    `json:"creative_id,omitempty"`
	DecorationMaterial *CreativeCustomCreativeUpdateV2RequestCreativeListInnerDecorationMaterial `json:"decoration_material,omitempty"`
	DerivePosterCid    *CreativeCustomCreativeUpdateV2CreativeListDerivePosterCid                `json:"derive_poster_cid,omitempty"`
	//
	ImageMaterials      []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerImageMaterialsInner `json:"image_materials,omitempty"`
	ImageMode           CreativeCustomCreativeUpdateV2CreativeListImageMode                          `json:"image_mode"`
	InteractiveMaterial *CreativeCustomCreativeUpdateV2RequestCreativeListInnerInteractiveMaterial   `json:"interactive_material,omitempty"`
	PlayableMaterial    *CreativeCustomCreativeUpdateV2RequestCreativeListInnerPlayableMaterial      `json:"playable_material,omitempty"`
	SubTitleMaterial    *CreativeCustomCreativeUpdateV2RequestCreativeListInnerSubTitleMaterial      `json:"sub_title_material,omitempty"`
	//
	ThirdPartyId  *string                                                              `json:"third_party_id,omitempty"`
	TitleMaterial *CreativeCustomCreativeUpdateV2RequestCreativeListInnerTitleMaterial `json:"title_material,omitempty"`
	VideoMaterial *CreativeCustomCreativeUpdateV2RequestCreativeListInnerVideoMaterial `json:"video_material,omitempty"`
}
