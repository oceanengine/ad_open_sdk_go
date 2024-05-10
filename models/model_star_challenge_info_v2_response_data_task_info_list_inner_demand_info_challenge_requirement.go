/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoChallengeRequirement
type StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoChallengeRequirement struct {
	//
	OralDemand *string `json:"oral_demand,omitempty"`
	//
	OtherDemand []string `json:"other_demand,omitempty"`
	//
	SceneDemand *string `json:"scene_demand,omitempty"`
	//
	StrongRequirements []string `json:"strong_requirements,omitempty"`
	//
	TitleDemand *string `json:"title_demand,omitempty"`
}
