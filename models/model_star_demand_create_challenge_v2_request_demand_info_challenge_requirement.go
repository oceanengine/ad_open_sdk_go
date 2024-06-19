/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestDemandInfoChallengeRequirement 产出物制作要求
type StarDemandCreateChallengeV2RequestDemandInfoChallengeRequirement struct {
	// 口播要求 60字内
	OralDemand *string `json:"oral_demand,omitempty"`
	// 其他要求 60字内
	OtherDemand []string `json:"other_demand,omitempty"`
	// 镜头要求 60字内
	SceneDemand string `json:"scene_demand"`
	//
	StrongRequirements []string `json:"strong_requirements,omitempty"`
	// 标题要求 80字内
	TitleDemand *string `json:"title_demand,omitempty"`
}
