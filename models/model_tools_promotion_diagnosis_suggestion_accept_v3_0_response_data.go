/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionDiagnosisSuggestionAcceptV30ResponseData
type ToolsPromotionDiagnosisSuggestionAcceptV30ResponseData struct {
	//
	SuggestionAcceptFailed []*ToolsPromotionDiagnosisSuggestionAcceptV30ResponseDataSuggestionAcceptFailedInner `json:"suggestion_accept_failed,omitempty"`
	//
	SuggestionAccepted []string `json:"suggestion_accepted,omitempty"`
}
