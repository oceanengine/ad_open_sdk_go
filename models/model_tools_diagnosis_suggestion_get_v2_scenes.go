/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsDiagnosisSuggestionGetV2Scenes
type ToolsDiagnosisSuggestionGetV2Scenes string

// List of tools_diagnosis_suggestion_get_v2_scenes
const (
	CLEAN_ToolsDiagnosisSuggestionGetV2Scenes     ToolsDiagnosisSuggestionGetV2Scenes = "CLEAN"
	POTENTIAL_ToolsDiagnosisSuggestionGetV2Scenes ToolsDiagnosisSuggestionGetV2Scenes = "POTENTIAL"
)

// All allowed values of ToolsDiagnosisSuggestionGetV2Scenes enum
var AllowedToolsDiagnosisSuggestionGetV2ScenesEnumValues = []ToolsDiagnosisSuggestionGetV2Scenes{
	"CLEAN",
	"POTENTIAL",
}

// NewToolsDiagnosisSuggestionGetV2ScenesFromValue returns a pointer to a valid ToolsDiagnosisSuggestionGetV2Scenes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsDiagnosisSuggestionGetV2ScenesFromValue(v string) (*ToolsDiagnosisSuggestionGetV2Scenes, error) {
	ev := ToolsDiagnosisSuggestionGetV2Scenes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsDiagnosisSuggestionGetV2Scenes: valid values are %v", v, AllowedToolsDiagnosisSuggestionGetV2ScenesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsDiagnosisSuggestionGetV2Scenes) IsValid() bool {
	for _, existing := range AllowedToolsDiagnosisSuggestionGetV2ScenesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_diagnosis_suggestion_get_v2_scenes value
func (v ToolsDiagnosisSuggestionGetV2Scenes) Ptr() *ToolsDiagnosisSuggestionGetV2Scenes {
	return &v
}