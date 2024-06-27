/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceActionScene
type ProjectUpdateV30AudienceActionScene string

// List of project_update_v3.0_audience_action_scene
const (
	AD_ProjectUpdateV30AudienceActionScene         ProjectUpdateV30AudienceActionScene = "AD"
	APP_ProjectUpdateV30AudienceActionScene        ProjectUpdateV30AudienceActionScene = "APP"
	E_COMMERCE_ProjectUpdateV30AudienceActionScene ProjectUpdateV30AudienceActionScene = "E-COMMERCE"
	NEWS_ProjectUpdateV30AudienceActionScene       ProjectUpdateV30AudienceActionScene = "NEWS"
	SEARCH_ProjectUpdateV30AudienceActionScene     ProjectUpdateV30AudienceActionScene = "SEARCH"
)

// All allowed values of ProjectUpdateV30AudienceActionScene enum
var AllowedProjectUpdateV30AudienceActionSceneEnumValues = []ProjectUpdateV30AudienceActionScene{
	"AD",
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewProjectUpdateV30AudienceActionSceneFromValue returns a pointer to a valid ProjectUpdateV30AudienceActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceActionSceneFromValue(v string) (*ProjectUpdateV30AudienceActionScene, error) {
	ev := ProjectUpdateV30AudienceActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceActionScene: valid values are %v", v, AllowedProjectUpdateV30AudienceActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceActionScene) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_action_scene value
func (v ProjectUpdateV30AudienceActionScene) Ptr() *ProjectUpdateV30AudienceActionScene {
	return &v
}
