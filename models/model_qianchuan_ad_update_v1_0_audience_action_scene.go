/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceActionScene
type QianchuanAdUpdateV10AudienceActionScene string

// List of qianchuan_ad_update_v1.0_audience_action_scene
const (
	APP_QianchuanAdUpdateV10AudienceActionScene        QianchuanAdUpdateV10AudienceActionScene = "APP"
	E_COMMERCE_QianchuanAdUpdateV10AudienceActionScene QianchuanAdUpdateV10AudienceActionScene = "E-COMMERCE"
	NEWS_QianchuanAdUpdateV10AudienceActionScene       QianchuanAdUpdateV10AudienceActionScene = "NEWS"
	SEARCH_QianchuanAdUpdateV10AudienceActionScene     QianchuanAdUpdateV10AudienceActionScene = "SEARCH"
)

// All allowed values of QianchuanAdUpdateV10AudienceActionScene enum
var AllowedQianchuanAdUpdateV10AudienceActionSceneEnumValues = []QianchuanAdUpdateV10AudienceActionScene{
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewQianchuanAdUpdateV10AudienceActionSceneFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceActionSceneFromValue(v string) (*QianchuanAdUpdateV10AudienceActionScene, error) {
	ev := QianchuanAdUpdateV10AudienceActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceActionScene: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceActionScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_action_scene value
func (v QianchuanAdUpdateV10AudienceActionScene) Ptr() *QianchuanAdUpdateV10AudienceActionScene {
	return &v
}
