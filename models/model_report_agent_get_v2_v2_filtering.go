/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAgentGetV2V2Filtering
type ReportAgentGetV2V2Filtering struct {
	AccountSource *ReportAgentGetV2V2FilteringAccountSource `json:"account_source,omitempty"`
	AccountStatus *ReportAgentGetV2V2FilteringAccountStatus `json:"account_status,omitempty"`
	Active        *ReportAgentGetV2V2FilteringActive        `json:"active,omitempty"`
	//
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
	//
	CompanyName *string `json:"company_name,omitempty"`
	//
	EndAuditPassTime **string `json:"end_audit_pass_time,omitempty"`
	//
	FirstIndustry *string `json:"first_industry,omitempty"`
	//
	SecondIndustry *string `json:"second_industry,omitempty"`
	//
	StartAuditPassTime **string `json:"start_audit_pass_time,omitempty"`
}
