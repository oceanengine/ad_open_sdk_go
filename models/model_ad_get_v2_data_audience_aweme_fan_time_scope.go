/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceAwemeFanTimeScope
type AdGetV2DataAudienceAwemeFanTimeScope string

// List of ad_get_v2_data_audience_aweme_fan_time_scope
const (
	FIFTEEN_DAYS_AdGetV2DataAudienceAwemeFanTimeScope AdGetV2DataAudienceAwemeFanTimeScope = "FIFTEEN_DAYS"
	SIXTY_DAYS_AdGetV2DataAudienceAwemeFanTimeScope   AdGetV2DataAudienceAwemeFanTimeScope = "SIXTY_DAYS"
	THIRTY_DAYS_AdGetV2DataAudienceAwemeFanTimeScope  AdGetV2DataAudienceAwemeFanTimeScope = "THIRTY_DAYS"
)

// Ptr returns reference to ad_get_v2_data_audience_aweme_fan_time_scope value
func (v AdGetV2DataAudienceAwemeFanTimeScope) Ptr() *AdGetV2DataAudienceAwemeFanTimeScope {
	return &v
}
