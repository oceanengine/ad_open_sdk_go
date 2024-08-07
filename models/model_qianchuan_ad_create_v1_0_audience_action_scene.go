/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceActionScene
type QianchuanAdCreateV10AudienceActionScene string

// List of qianchuan_ad_create_v1.0_audience_action_scene
const (
	APP_QianchuanAdCreateV10AudienceActionScene        QianchuanAdCreateV10AudienceActionScene = "APP"
	E_COMMERCE_QianchuanAdCreateV10AudienceActionScene QianchuanAdCreateV10AudienceActionScene = "E-COMMERCE"
	NEWS_QianchuanAdCreateV10AudienceActionScene       QianchuanAdCreateV10AudienceActionScene = "NEWS"
	SEARCH_QianchuanAdCreateV10AudienceActionScene     QianchuanAdCreateV10AudienceActionScene = "SEARCH"
)

// All allowed values of QianchuanAdCreateV10AudienceActionScene enum
var AllowedQianchuanAdCreateV10AudienceActionSceneEnumValues = []QianchuanAdCreateV10AudienceActionScene{
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewQianchuanAdCreateV10AudienceActionSceneFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceActionSceneFromValue(v string) (*QianchuanAdCreateV10AudienceActionScene, error) {
	ev := QianchuanAdCreateV10AudienceActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceActionScene: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceActionScene) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_action_scene value
func (v QianchuanAdCreateV10AudienceActionScene) Ptr() *QianchuanAdCreateV10AudienceActionScene {
	return &v
}
