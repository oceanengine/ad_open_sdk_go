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

// QianchuanToolsEstimateAudienceV10AudienceActionScene
type QianchuanToolsEstimateAudienceV10AudienceActionScene string

// List of qianchuan_tools_estimate_audience_v1.0_audience_action_scene
const (
	APP_QianchuanToolsEstimateAudienceV10AudienceActionScene        QianchuanToolsEstimateAudienceV10AudienceActionScene = "APP"
	E_COMMERCE_QianchuanToolsEstimateAudienceV10AudienceActionScene QianchuanToolsEstimateAudienceV10AudienceActionScene = "E-COMMERCE"
	NEWS_QianchuanToolsEstimateAudienceV10AudienceActionScene       QianchuanToolsEstimateAudienceV10AudienceActionScene = "NEWS"
	SEARCH_QianchuanToolsEstimateAudienceV10AudienceActionScene     QianchuanToolsEstimateAudienceV10AudienceActionScene = "SEARCH"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceActionScene enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceActionSceneEnumValues = []QianchuanToolsEstimateAudienceV10AudienceActionScene{
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewQianchuanToolsEstimateAudienceV10AudienceActionSceneFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceActionSceneFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceActionScene, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceActionScene: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceActionScene) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_action_scene value
func (v QianchuanToolsEstimateAudienceV10AudienceActionScene) Ptr() *QianchuanToolsEstimateAudienceV10AudienceActionScene {
	return &v
}
