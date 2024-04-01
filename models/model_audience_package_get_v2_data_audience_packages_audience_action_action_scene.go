/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene
type AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene string

// List of audience_package_get_v2_data_audience_packages_audience_action_action_scene
const (
	AD_AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene         AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene = "AD"
	APP_AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene        AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene = "APP"
	E_COMMERCE_AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene = "E-COMMERCE"
	NEWS_AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene       AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene = "NEWS"
	SEARCH_AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene     AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene = "SEARCH"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActionActionSceneEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene{
	"AD",
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceActionActionSceneFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceActionActionSceneFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActionActionSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActionActionSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_action_action_scene value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceActionActionScene {
	return &v
}
