/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceInterestActionListV2FilteringActionScene
type ReportAudienceInterestActionListV2FilteringActionScene string

// List of report_audience_interest_action_list_v2_filtering_action_scene
const (
	SEARCH_ReportAudienceInterestActionListV2FilteringActionScene     ReportAudienceInterestActionListV2FilteringActionScene = "SEARCH"
	E_COMMERCE_ReportAudienceInterestActionListV2FilteringActionScene ReportAudienceInterestActionListV2FilteringActionScene = "E-COMMERCE"
	NEWS_ReportAudienceInterestActionListV2FilteringActionScene       ReportAudienceInterestActionListV2FilteringActionScene = "NEWS"
	AD_ReportAudienceInterestActionListV2FilteringActionScene         ReportAudienceInterestActionListV2FilteringActionScene = "AD"
	APP_ReportAudienceInterestActionListV2FilteringActionScene        ReportAudienceInterestActionListV2FilteringActionScene = "APP"
)

// Ptr returns reference to report_audience_interest_action_list_v2_filtering_action_scene value
func (v ReportAudienceInterestActionListV2FilteringActionScene) Ptr() *ReportAudienceInterestActionListV2FilteringActionScene {
	return &v
}
