/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceActionActionScene
type AdGetV2DataAudienceActionActionScene string

// List of ad_get_v2_data_audience_action_action_scene
const (
	NEWS_AdGetV2DataAudienceActionActionScene       AdGetV2DataAudienceActionActionScene = "NEWS"
	AD_AdGetV2DataAudienceActionActionScene         AdGetV2DataAudienceActionActionScene = "AD"
	APP_AdGetV2DataAudienceActionActionScene        AdGetV2DataAudienceActionActionScene = "APP"
	E_COMMERCE_AdGetV2DataAudienceActionActionScene AdGetV2DataAudienceActionActionScene = "E-COMMERCE"
	SEARCH_AdGetV2DataAudienceActionActionScene     AdGetV2DataAudienceActionActionScene = "SEARCH"
)

// Ptr returns reference to ad_get_v2_data_audience_action_action_scene value
func (v AdGetV2DataAudienceActionActionScene) Ptr() *AdGetV2DataAudienceActionActionScene {
	return &v
}
