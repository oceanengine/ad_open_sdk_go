/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestCreative
type CreativeProceduralCreativeCreateV2RequestCreative struct {
	//
	AbstractMaterials []*CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInner `json:"abstract_materials,omitempty"`
	//
	ComponentMaterials []*CreativeProceduralCreativeCreateV2RequestCreativeComponentMaterialsInner `json:"component_materials,omitempty"`
	DecorationMaterial *CreativeProceduralCreativeCreateV2RequestCreativeDecorationMaterial        `json:"decoration_material,omitempty"`
	//
	ImageMaterials   []*CreativeProceduralCreativeCreateV2RequestCreativeImageMaterialsInner `json:"image_materials,omitempty"`
	SubTitleMaterial *CreativeProceduralCreativeCreateV2RequestCreativeSubTitleMaterial      `json:"sub_title_material,omitempty"`
	//
	TitleMaterials []*CreativeProceduralCreativeCreateV2RequestCreativeTitleMaterialsInner `json:"title_materials,omitempty"`
	//
	VideoMaterials []*CreativeProceduralCreativeCreateV2RequestCreativeVideoMaterialsInner `json:"video_materials,omitempty"`
}
