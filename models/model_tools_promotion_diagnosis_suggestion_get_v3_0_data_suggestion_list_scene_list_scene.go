/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene
type ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene string

// List of tools_promotion_diagnosis_suggestion_get_v3.0_data_suggestion_list_scene_list_scene
const (
	CLEAN_ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene     ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene = "CLEAN"
	POTENTIAL_ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene = "POTENTIAL"
	ZOMBIE_ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene    ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene = "ZOMBIE"
)

// All allowed values of ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene enum
var AllowedToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSceneEnumValues = []ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene{
	"CLEAN",
	"POTENTIAL",
	"ZOMBIE",
}

// NewToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSceneFromValue returns a pointer to a valid ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSceneFromValue(v string) (*ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene, error) {
	ev := ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene: valid values are %v", v, AllowedToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene) IsValid() bool {
	for _, existing := range AllowedToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_diagnosis_suggestion_get_v3.0_data_suggestion_list_scene_list_scene value
func (v ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene) Ptr() *ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene {
	return &v
}
