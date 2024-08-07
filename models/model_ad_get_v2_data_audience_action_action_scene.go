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

// AdGetV2DataAudienceActionActionScene
type AdGetV2DataAudienceActionActionScene string

// List of ad_get_v2_data_audience_action_action_scene
const (
	SEARCH_AdGetV2DataAudienceActionActionScene     AdGetV2DataAudienceActionActionScene = "SEARCH"
	APP_AdGetV2DataAudienceActionActionScene        AdGetV2DataAudienceActionActionScene = "APP"
	NEWS_AdGetV2DataAudienceActionActionScene       AdGetV2DataAudienceActionActionScene = "NEWS"
	E_COMMERCE_AdGetV2DataAudienceActionActionScene AdGetV2DataAudienceActionActionScene = "E-COMMERCE"
	AD_AdGetV2DataAudienceActionActionScene         AdGetV2DataAudienceActionActionScene = "AD"
)

// All allowed values of AdGetV2DataAudienceActionActionScene enum
var AllowedAdGetV2DataAudienceActionActionSceneEnumValues = []AdGetV2DataAudienceActionActionScene{
	"SEARCH",
	"APP",
	"NEWS",
	"E-COMMERCE",
	"AD",
}

// NewAdGetV2DataAudienceActionActionSceneFromValue returns a pointer to a valid AdGetV2DataAudienceActionActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceActionActionSceneFromValue(v string) (*AdGetV2DataAudienceActionActionScene, error) {
	ev := AdGetV2DataAudienceActionActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceActionActionScene: valid values are %v", v, AllowedAdGetV2DataAudienceActionActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceActionActionScene) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceActionActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_action_action_scene value
func (v AdGetV2DataAudienceActionActionScene) Ptr() *AdGetV2DataAudienceActionActionScene {
	return &v
}
